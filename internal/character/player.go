package character

import (
	"log"

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
	previous      state
	current       state
	jumpSe        *se.Player
}

// SetPosition sets the position of this character.
func (p *Player) SetPosition(pos Position) {
	p.Position = pos
}

// Update updates the character regarding the user input.
func (p *Player) Update() {
	p.updateState()
	p.updatePosition()
	err := p.playSe()
	if err != nil {
		log.Println("failed to play SE:", err)
		return
	}
}

func (p *Player) updateState() {
	// todo:
	// ひとつ上のレーンに到達したら上昇を終了

	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.GamepadAxis(0, 1) <= -0.5 {
		if p.current == stop || p.current == walk || p.current == dash {
			p.previous = p.current
			p.current = ascending
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.GamepadAxis(0, 1) >= 0.5 {
		if p.current == stop || p.current == walk || p.current == dash {
			p.previous = p.current
			p.current = descending
		}
	}
}

func (p *Player) updatePosition() {
	// todo: 固定値での移動ではなくキャラごと、stateごとの初速度と加速度から算出される速度で移動させる
	switch p.current {
	case walk:
		p.Position.X++
		p.animation.AddStep(1)
	case dash:
		p.Position.X += 2
		p.animation.AddStep(2)
	default:
		// Don't move
	}
}

// Draw draws the character image.
func (p *Player) Draw(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.Position.X), float64(p.Position.Y))
	return screen.DrawImage(p.animation.GetCurrentFrame(), op)
}

func (p *Player) playSe() error {
	if p.previous != ascending && p.current == ascending {
		return p.jumpSe.Play()
	}
	return nil
}

// Close closes the inner resources.
func (p *Player) Close() error {
	return p.jumpSe.Close()
}
