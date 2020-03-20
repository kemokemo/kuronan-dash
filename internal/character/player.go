package character

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/assets/messages"
	"github.com/kemokemo/kuronan-dash/assets/se"
	"github.com/kemokemo/kuronan-dash/internal/view"
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

// Player is a player character.
type Player struct {
	position      view.Vector
	StandingImage *ebiten.Image
	Description   string
	animation     *StepAnimation
	previous      State
	current       State
	velocity      view.Vector
	lanes         Lanes
	jumpSe        *se.Player
}

// SetLanes sets the lanes information.
func (p *Player) SetLanes(heights []float64) error {
	p.lanes = Lanes{}
	charaHeights := []float64{}
	_, h := p.StandingImage.Size()

	for index := 0; index < len(heights); index++ {
		charaHeights = append(charaHeights, heights[index]-float64(h))
	}

	err := p.lanes.SetHeights(charaHeights)
	if err != nil {
		return err
	}

	// set the player at the top lane.
	p.position = view.Vector{X: 10.0, Y: float64(charaHeights[0])}

	return nil
}

// Start starts dash!
func (p *Player) Start() {
	p.current = Dash
}

// Pause pauses this character.
func (p *Player) Pause() {
	if p.current == Pause {
		return
	}
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
		if p.lanes.IsReachedTarget(p.position.Y) {
			p.current = p.previous
		}
	case Descending:
		if p.lanes.IsReachedTarget(p.position.Y) {
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
		p.velocity.X = 1.0
		p.velocity.Y = 0.0
		p.animation.AddStep(1)
	case Dash:
		p.velocity.X = 2.0
		p.velocity.Y = 0.0
		p.animation.AddStep(2)
	case Ascending:
		p.velocity.X = 1.0
		p.velocity.Y = -2.0
		p.animation.AddStep(1)
	case Descending:
		p.velocity.X = 1.0
		p.velocity.Y = 2.0
		p.animation.AddStep(1)
	default:
		// Don't move
	}
	p.position = p.position.Add(p.velocity)
}

// Draw draws the character image.
func (p *Player) Draw(screen *ebiten.Image) error {
	// TODO: 状態を出してデバッグするよ
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Position: %v, State: pre(%v) cur(%v)",
		p.position, p.previous, p.current))

	// TODO: ダッシュ中とか奥義中とか状態に応じて多少前後しつつ、ほぼ画面中央に描画したい
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(view.ScreenWidth/4, p.position.Y)
	return screen.DrawImage(p.animation.GetCurrentFrame(), op)
}

func (p *Player) playSe() error {
	if p.previous != Ascending && p.current == Ascending {
		return p.jumpSe.Play()
	}
	return nil
}

func (p *Player) GetPosition() view.Vector {
	return p.position
}

// GetVelocity returns the velocity of this playable character.
func (p *Player) GetVelocity() view.Vector {
	return p.velocity
}

// Close closes the inner resources.
func (p *Player) Close() error {
	return p.jumpSe.Close()
}
