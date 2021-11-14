package character

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	vpad "github.com/kemokemo/ebiten-virtualpad"
	"github.com/kemokemo/kuronan-dash/assets/se"
	"github.com/kemokemo/kuronan-dash/internal/anime"
	"github.com/kemokemo/kuronan-dash/internal/field"
	"github.com/kemokemo/kuronan-dash/internal/move"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// Player is a player character.
type Player struct {
	// Specified at creation and not changed
	StandingImage *ebiten.Image
	Description   string
	animation     *anime.StepAnimation
	jumpSe        *se.Player
	dropSe        *se.Player

	// Update each time based on the internal status and other information
	op         *ebiten.DrawImageOptions
	vc         move.VelocityController
	scrollV    *view.Vector
	tempPosV   *view.Vector
	charaPosV  *view.Vector
	tempDrawV  *view.Vector
	charaDrawV *view.Vector
	rect       *view.HitRectangle

	// Initialization is required before starting the stage.
	stateMachine *move.StateMachine
	previous     move.State
	current      move.State
	stamina      *Stamina
	sumTicks     float64
}

// InitializeWithLanesInfo sets the lanes information.
// The player can run on the lane or move between lanes based on the lane drawing height information received in the argument.
func (p *Player) InitializeWithLanes(lanes *field.Lanes) error {
	p.previous = move.Pause
	p.current = move.Dash
	p.stamina.Initialize()

	var err error
	p.stateMachine, err = move.NewStateMachine(lanes)
	if err != nil {
		return err
	}

	// set the player at the top lane.
	w, h := p.StandingImage.Size()

	initialY := lanes.GetTargetLaneHeight() - float64(h) + field.FieldOffset
	p.charaPosV = &view.Vector{X: 0.0, Y: 0.0}
	p.charaDrawV = &view.Vector{X: 0.0, Y: 0.0}
	p.scrollV = &view.Vector{X: 0.0, Y: 0.0}
	p.op = &ebiten.DrawImageOptions{}
	p.op.GeoM.Translate(view.DrawPosition, initialY)

	rectOffset := 3.0
	p.rect = view.NewHitRectangle(
		view.Vector{X: view.DrawPosition + rectOffset, Y: initialY + rectOffset},
		view.Vector{X: view.DrawPosition + float64(w) - rectOffset, Y: initialY + float64(h) - rectOffset})

	return nil
}

func (p *Player) SetInputChecker(laneRectArray []image.Rectangle, upBtn, downBtn vpad.TriggerButton) {
	p.stateMachine.SetInputChecker(laneRectArray, upBtn, downBtn)
}

// Start starts playing.
func (p *Player) Start() {
	p.current = move.Dash
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
	p.current = p.stateMachine.Update(p.stamina.GetStamina(), p.charaPosV)

	p.sumTicks += 1.0 / ebiten.CurrentTPS()
	if p.sumTicks >= 0.05 {
		p.sumTicks = 0.0
		p.stamina.Consumes(p.current)
	}

	// 次に動くべき速度にオフセットを適用
	p.updateVelWithOffset(p.stateMachine.GetOffsetV())

	p.animation.AddStep(p.charaPosV.X)
	p.op.GeoM.Translate(p.charaDrawV.X, p.charaDrawV.Y)
	p.rect.Add(p.charaDrawV)
}

func (p *Player) updateVelWithOffset(offsetV *view.Vector) {
	p.charaPosV.X = p.tempPosV.X
	p.charaPosV.Y = p.tempPosV.Y + offsetV.Y

	p.charaDrawV.X = p.tempDrawV.X
	p.charaDrawV.Y = p.tempDrawV.Y + offsetV.Y
}

// Draw draws the character image.
func (p *Player) Draw(screen *ebiten.Image) {
	// TODO: ダッシュ中とか奥義中とか状態に応じて多少前後しつつ、ほぼ画面中央に描画したい
	screen.DrawImage(p.animation.GetCurrentFrame(), p.op)
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

// GetRectangle returns the edge rentangle of this player.
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
	var err, e error
	e = p.jumpSe.Close()
	if e != nil {
		err = fmt.Errorf("%v:%v", err, e)
	}
	e = p.dropSe.Close()
	if e != nil {
		err = fmt.Errorf("%v:%v", err, e)
	}
	return err
}

func (p *Player) GetHeight() float64 {
	_, h := p.StandingImage.Size()
	return float64(h)
}
