package character

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	gauge "github.com/kemokemo/ebiten-gauge"
	vpad "github.com/kemokemo/ebiten-virtualpad"
	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/assets/se"
	"github.com/kemokemo/kuronan-dash/internal/anime"
	"github.com/kemokemo/kuronan-dash/internal/field"
	"github.com/kemokemo/kuronan-dash/internal/move"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// Player is a player character.
type Player struct {
	// Specified at creation and not changed
	StandingImage  *ebiten.Image
	MapIcon        *ebiten.Image
	Description    string
	attackImage    *ebiten.Image
	skillImage     *ebiten.Image
	skillEffect    *ebiten.Image
	animation      *anime.StepAnimation
	spReadyIcon    *ebiten.Image
	spReadyIconOp  *gauge.BlinkingOp
	walkIcon       *ebiten.Image
	walkIconOp     *gauge.BlinkingOp
	jumpSe         *se.Player
	dropSe         *se.Player
	collisionSe    *se.Player
	attackSe       *se.Player
	spVoice        *se.Player
	atkMaxDuration int
	spMaxDuration  int

	// Update each time based on the internal status and other information
	op         *ebiten.DrawImageOptions
	atkOp      *ebiten.DrawImageOptions
	spEffectOp *ebiten.DrawImageOptions
	spOp       *ebiten.DrawImageOptions
	vc         move.VelocityController
	scrollV    *view.Vector
	tempPosV   *view.Vector
	charaPosV  *view.Vector
	tempDrawV  *view.Vector
	charaDrawV *view.Vector
	rect       *view.HitRectangle
	atkRect    *view.HitRectangle

	// Initialization is required before starting the stage.
	stateMachine *move.StateMachine
	previous     move.State
	current      move.State
	stamina      *Stamina
	sumTicks     float64
	power        float64
	tension      *Tension

	soundTypeCh chan se.SoundType
}

// InitializeWithLanesInfo sets the lanes information.
// The player can run on the lane or move between lanes based on the lane drawing height information received in the argument.
func (p *Player) InitializeWithLanes(lanes *field.Lanes) error {
	p.previous = move.Wait
	p.current = move.Wait
	p.stamina.Initialize()
	p.tension.Initialize()

	var err error
	p.soundTypeCh = make(chan se.SoundType)
	p.stateMachine, err = move.NewStateMachine(lanes, p.atkMaxDuration, p.spMaxDuration)
	if err != nil {
		return err
	}
	p.stateMachine.SetSeChan(p.soundTypeCh)

	// set the player at the top lane.
	w, h := p.StandingImage.Size()
	sw, sh := p.skillEffect.Size()
	aw, ah := p.attackImage.Size()

	initialY := lanes.GetTargetLaneHeight() - float64(h) + field.FieldOffset
	p.charaPosV = &view.Vector{X: 0.0, Y: 0.0}
	p.charaDrawV = &view.Vector{X: 0.0, Y: 0.0}
	p.scrollV = &view.Vector{X: 0.0, Y: 0.0}
	p.op = &ebiten.DrawImageOptions{}
	p.op.GeoM.Translate(view.DrawPosition, initialY)
	p.spEffectOp = &ebiten.DrawImageOptions{}
	p.spEffectOp.GeoM.Translate(view.DrawPosition-float64((sw-w)/2), initialY-float64(sh-h))
	p.atkOp = &ebiten.DrawImageOptions{}
	p.atkOp.GeoM.Translate(view.DrawPosition+float64(w)+5, initialY+20)
	p.spOp = &ebiten.DrawImageOptions{}

	p.spReadyIcon = images.SpecialReadyIcon
	p.spReadyIconOp = gauge.NewBlinkingOp()
	p.spReadyIconOp.SetInterval(20)
	p.spReadyIconOp.Op.GeoM.Translate(view.DrawPosition+float64(w-2), initialY-7.0)
	p.walkIcon = images.WalkStateIcon
	p.walkIconOp = gauge.NewBlinkingOp()
	p.walkIconOp.SetInterval(20)
	p.walkIconOp.Op.GeoM.Translate(view.DrawPosition-float64(5), initialY-7.0)

	rectOffset := 3.0
	p.rect = view.NewHitRectangle(
		view.Vector{X: view.DrawPosition + rectOffset, Y: initialY + rectOffset},
		view.Vector{X: view.DrawPosition + float64(w) - rectOffset, Y: initialY + float64(h) - rectOffset})
	p.atkRect = view.NewHitRectangle(
		view.Vector{X: view.DrawPosition + rectOffset, Y: initialY + 20 + rectOffset},
		view.Vector{X: view.DrawPosition + float64(w) + 5 + float64(aw) - rectOffset, Y: initialY + 20 + float64(ah) - rectOffset})

	return nil
}

func (p *Player) SetInputChecker(laneRectArray []image.Rectangle, upBtn, downBtn, atkBtn, spBtn vpad.TriggerButton) {
	p.stateMachine.SetInputChecker(laneRectArray, upBtn, downBtn, atkBtn, spBtn)
}

// Start starts playing.
func (p *Player) Start() {
	p.current = move.Dash

	go p.playSounds()
}

func (p *Player) playSounds() {
	for s := range p.soundTypeCh {
		switch s {
		case se.Jump:
			p.jumpSe.Play()
		case se.Drop:
			p.dropSe.Play()
		case se.Attack:
			p.attackSe.Play()
		case se.Blocked:
			p.collisionSe.Play()
		default:
			log.Println("unknown sound type, ", s)
		}
	}
}

// Pause pauses this character.
func (p *Player) Pause() {
	if p.current == move.Pause {
		return
	}
	p.previous = p.current
	p.current = move.Pause
}

// ReStart starts again this character.
func (p *Player) ReStart() {
	p.current = p.previous
}

// Update updates the character regarding the user input.
func (p *Player) Update() {
	// ひとつ前に更新したStateをもとに、次に動くべき速度を入手
	p.vc.SetState(p.current)
	p.scrollV, p.tempPosV, p.tempDrawV = p.vc.GetVelocity()

	// 次に動くべき速度から次のStateを決定
	// State更新処理で判明した、レーンにめり込まないようにするためのオフセットを入手
	p.current = p.stateMachine.Update(
		p.stamina.GetStamina(),
		p.tension.Get(),
		p.tension.IsMax(),
		p.charaPosV)
	if p.current == move.SkillEffect || p.current == move.Pause {
		return
	}

	p.sumTicks += 1.0 / ebiten.ActualTPS()
	if p.sumTicks >= 0.05 {
		p.sumTicks = 0.0
		p.stamina.ConsumesByState(p.current)
		p.tension.AddByState(p.current)
	}

	// 次に動くべき速度にオフセットを適用
	p.updateVelWithOffset(p.stateMachine.GetOffsetV())

	p.animation.AddStep(p.charaPosV.X)
	p.op.GeoM.Translate(p.charaDrawV.X, p.charaDrawV.Y)
	p.spEffectOp.GeoM.Translate(p.charaDrawV.X, p.charaDrawV.Y)
	p.spReadyIconOp.Op.GeoM.Translate(p.charaDrawV.X, p.charaDrawV.Y)
	if p.tension.IsMax() {
		p.spReadyIconOp.Update()
	}
	p.walkIconOp.Op.GeoM.Translate(p.charaDrawV.X, p.charaDrawV.Y)
	if p.current == move.Walk || p.current == move.SkillWalk {
		p.walkIconOp.Update()
	}
	p.atkOp.GeoM.Translate(p.charaDrawV.X, p.charaDrawV.Y)
	p.rect.Add(p.charaDrawV)
	p.atkRect.Add(p.charaDrawV)
}

func (p *Player) updateVelWithOffset(offsetV *view.Vector) {
	p.charaPosV.X = p.tempPosV.X + offsetV.X
	p.charaPosV.Y = p.tempPosV.Y + offsetV.Y

	p.charaDrawV.X = p.tempDrawV.X + offsetV.X
	p.charaDrawV.Y = p.tempDrawV.Y + offsetV.Y
}

// Draw draws the character image.
func (p *Player) Draw(screen *ebiten.Image) {
	if p.current == move.Wait {
		screen.DrawImage(p.StandingImage, p.op)
		return
	}

	if p.current == move.SkillDash || p.current == move.SkillWalk || p.current == move.SkillAscending || p.current == move.SkillDescending {
		screen.DrawImage(p.skillEffect, p.spEffectOp)
	}
	screen.DrawImage(p.animation.GetCurrentFrame(), p.op)
	if p.stateMachine.DrawAttack() {
		screen.DrawImage(p.attackImage, p.atkOp)
	}

	if p.current == move.Walk || p.current == move.SkillWalk {
		screen.DrawImage(p.walkIcon, p.walkIconOp.Op)
	}

	if p.tension.IsMax() {
		screen.DrawImage(p.spReadyIcon, p.spReadyIconOp.Op)
	}
}

func (p *Player) DrawSkillEffect(screen *ebiten.Image) {
	screen.DrawImage(p.skillImage, p.spOp)
}

// GetPosition return the current position of this player.
func (p *Player) GetPosition() *view.Vector {
	return p.stateMachine.GetPosition()
}

// GetScrollVelocity returns the velocity to scroll field parts.
func (p *Player) GetScrollVelocity() *view.Vector {
	return p.scrollV
}

// GetStamina returns the stamina value fo this character.
func (p *Player) GetStamina() int {
	return p.stamina.GetStamina()
}

func (p *Player) GetMaxStamina() float64 {
	return p.stamina.GetMaxStamina()
}

// GetRectangle returns the edge rectangle of this player.
func (p *Player) GetRectangle() *view.HitRectangle {
	return p.rect
}

// BeBlocked puts the player in a position where the path is blocked by an obstacle.
func (p *Player) BeBlocked(isBlocked bool) {
	p.stateMachine.SetBlockState(isBlocked)
}

// Eat eats foods and restore stamina value by argument value.
func (p *Player) Eat(foodVol int) {
	p.stamina.Add(foodVol)
}

// Close closes the inner resources.
func (p *Player) Close() error {
	// assets側でcloseするので、ここではcloseしない
	close(p.soundTypeCh)
	return nil
}

func (p *Player) GetHeight() float64 {
	_, h := p.StandingImage.Size()
	return float64(h)
}

func (p *Player) IsAttacked() (bool, *view.HitRectangle, float64) {
	return p.stateMachine.Attacked(), p.atkRect, p.power
}

func (p *Player) ConsumeStaminaByAttack(num int) {
	p.stamina.ConsumeByAttack(num)
}

func (p *Player) AddTension(num int) {
	p.tension.AddByAttack(num)
}

func (p *Player) GetTension() int {
	return p.tension.Get()
}

func (p *Player) GetMaxTension() float64 {
	return p.tension.GetMax()
}

func (p *Player) StartSpEffect() bool {
	if p.stateMachine.StartSpEffect() {
		p.spVoice.Play()
		return true
	}
	return false
}

func (p *Player) UpdateSkillEffect() {
	p.stateMachine.UpdateSkillEffect(p.spVoice.IsPlaying())
}

func (p *Player) FinishSpEffect() bool {
	return p.stateMachine.FinishSpEffect()
}

func (p *Player) SetVolumeFlag(isVolumeOn bool) {
	p.spVoice.SetVolumeFlag(isVolumeOn)

	p.jumpSe.SetVolumeFlag(isVolumeOn)
	p.dropSe.SetVolumeFlag(isVolumeOn)
	p.collisionSe.SetVolumeFlag(isVolumeOn)
	p.attackSe.SetVolumeFlag(isVolumeOn)
}
