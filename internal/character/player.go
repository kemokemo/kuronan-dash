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
	previous      State
	current       State
	lanes         Lanes
	jumpSe        *se.Player
}

// SetLanes sets the lanes information.
func (p *Player) SetLanes(heights []int) error {
	p.lanes = Lanes{}
	charaHeights := []int{}
	_, h := p.StandingImage.Size()

	for index := 0; index < len(heights); index++ {
		charaHeights = append(charaHeights, heights[index]-h)
	}

	err := p.lanes.SetHeights(charaHeights)
	if err != nil {
		return err
	}

	// set the player at the top lane.
	p.Position = Position{X: 10, Y: charaHeights[0]}

	return nil
}

// Start starts dash!
func (p *Player) Start() {
	p.current = Dash
}

// Stop stops this character.
func (p *Player) Stop() {
	p.previous = p.current
	p.current = Stop
}

// Pause pauses this character.
func (p *Player) Pause() {
	p.previous = p.current
	p.current = Pause
}

// ReStart starts again this character.
func (p *Player) ReStart() {
	p.current = p.previous
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
	switch p.current {
	case Pause:
		return
	case Ascending:
		if p.lanes.IsReachedTarget(p.Position.Y) {
			p.current = p.previous
		}
	case Descending:
		if p.lanes.IsReachedTarget(p.Position.Y) {
			p.current = p.previous
		}
	default:
		if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.GamepadAxis(0, 1) <= -0.5 {
			if !p.lanes.IsTop() {
				if p.lanes.Ascend() {
					p.previous = p.current
					p.current = Ascending
				}
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.GamepadAxis(0, 1) >= 0.5 {
			if !p.lanes.IsBottom() {
				if p.lanes.Descend() {
					p.previous = p.current
					p.current = Descending
				}
			}
		}
	}
}

func (p *Player) updatePosition() {
	// todo: 固定値での移動ではなくキャラごと、stateごとの初速度と加速度から算出される速度で移動させる
	switch p.current {
	case Walk:
		p.Position.X++
		p.animation.AddStep(1)
	case Dash:
		p.Position.X++
		p.animation.AddStep(2)
	case Ascending:
		p.Position.X++
		p.Position.Y -= 2
		p.animation.AddStep(1)
	case Descending:
		p.Position.X++
		p.Position.Y += 2
		p.animation.AddStep(1)
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
	if p.previous != Ascending && p.current == Ascending {
		return p.jumpSe.Play()
	}
	return nil
}

// GetState returns the current state of this character.
func (p *Player) GetState() State {
	return p.current
}

// Close closes the inner resources.
func (p *Player) Close() error {
	return p.jumpSe.Close()
}
