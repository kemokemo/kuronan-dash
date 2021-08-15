package character

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
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
	charaPosV  *view.Vector
	charaDrawV *view.Vector
	rect       *view.HitRectangle

	// Initialization is required before starting the stage.
	stateMachine move.StateMachine
	previous     move.State
	current      move.State
	stamina      *Stamina
}

// InitializeWithLanesInfo sets the lanes information.
// The player can run on the lane or move between lanes based on the lane drawing height information received in the argument.
func (p *Player) InitializeWithLanesInfo(heights []float64) error {
	p.previous = move.Pause
	p.current = move.Dash
	p.stamina.Initialize()

	cH := []float64{}
	_, h := p.StandingImage.Size()
	for i := 0; i < len(heights); i++ {
		cH = append(cH, heights[i]-float64(h))
	}

	p.stateMachine = move.NewStateMachine()
	err := p.stateMachine.SetHeights(cH)
	if err != nil {
		return err
	}

	// set the player at the top lane.
	initialY := float64(cH[0]) + field.FieldOffset
	rectOffset := 3.0

	p.charaPosV = &view.Vector{X: 0.0, Y: 0.0}
	p.scrollV = &view.Vector{X: 0.0, Y: 0.0}
	p.op = &ebiten.DrawImageOptions{}
	p.op.GeoM.Translate(view.DrawPosition, initialY)

	w, h := p.StandingImage.Size()
	p.rect = view.NewHitRectangle(
		view.Vector{X: view.DrawPosition + rectOffset, Y: initialY + rectOffset},
		view.Vector{X: view.DrawPosition + float64(w) - rectOffset, Y: initialY + float64(h) - rectOffset})

	return nil
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
	p.current = p.stateMachine.Update(p.stamina.GetStamina(), p.charaPosV)

	p.stamina.Consumes(p.current)
	p.scrollV, p.charaPosV, p.charaDrawV = p.vc.GetVelocity(p.current)

	p.animation.AddStep(p.charaPosV.X)
	p.op.GeoM.Translate(p.charaDrawV.X, p.charaDrawV.Y)
	p.rect.Add(p.charaDrawV)

	// todo
	log.Printf("state:%v, posV-Y:%v, drawV-Y:%v", p.current, p.charaPosV.Y, p.charaDrawV.Y)
	log.Printf("pos:%v, drawPos:%v", p.stateMachine.GetPosition(), p.op.GeoM.String())
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
