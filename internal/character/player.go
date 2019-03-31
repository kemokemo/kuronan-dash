package character

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/assets/messages"
	"github.com/kemokemo/kuronan-dash/assets/se"
)

// player characters
var (
	Kurona     *Player
	Koma       *Player
	Shishimaru *Player

	// Selected is the selected player.
	Selected *Player
)

// NewPlayers load all player characters.
func NewPlayers() error {
	Kurona = &Player{
		StandingImage: images.KuronaStanding,
		Description:   messages.DescKurona,
		animation:     NewStepAnimation(images.KuronaAnimation, 5),
		jumpSe:        se.Jump,
	}

	Koma = &Player{
		StandingImage: images.KomaStanding,
		Description:   messages.DescKoma,
		animation:     NewStepAnimation(images.KomaAnimation, 5),
		jumpSe:        se.Jump,
	}

	Shishimaru = &Player{
		StandingImage: images.ShishimaruStanding,
		Description:   messages.DescShishimaru,
		animation:     NewStepAnimation(images.ShishimaruAnimation, 5),
		jumpSe:        se.Jump,
	}

	Selected = Kurona

	return nil
}

// Position describes the position by x and y.
type Position struct {
	X int
	Y int
}

// Player is a player character.
type Player struct {
	Position      Position
	StandingImage *ebiten.Image
	Description   string
	animation     *StepAnimation
	moved         bool
	state         State
	jumpSe        *se.Player
}

// SetInitialPosition sets the initial position for this character.
func (p *Player) SetInitialPosition(pos Position) {
	p.Position = pos
}

// Move moves the character regarding the user input.
func (p *Player) Move() {
	p.moved = false
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.GamepadAxis(0, 0) <= -0.5 {
		p.Position.X--
		p.moved = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.GamepadAxis(0, 0) >= 0.5 {
		p.Position.X++
		p.moved = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.GamepadAxis(0, 1) <= -0.5 {
		p.Position.Y--
		p.moved = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.GamepadAxis(0, 1) >= 0.5 {
		p.Position.Y++
		p.moved = true
	}

	if p.moved {
		p.animation.AddStep(1)
	}
}

// Draw draws the character image.
func (p *Player) Draw(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.Position.X), float64(p.Position.Y))
	return screen.DrawImage(p.animation.GetCurrentFrame(), op)
}

// SetState sets the status for this character.
func (p *Player) SetState(state State) {
	p.state = state
}

// PlaySe plays a sound effect according to the status of this character.
func (p *Player) PlaySe() error {
	// TODO: ステータスに応じたSEを再生
	if p.state == Ascending {
		// TODO: このままだとジャンプ中SE再生し続けるので対策が必要
		return p.jumpSe.Play()
	}
	return nil
}

// Close closes the inner resources.
func (p *Player) Close() error {
	return p.jumpSe.Close()
}
