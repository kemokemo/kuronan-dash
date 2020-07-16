package character

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
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
		stamina:       NewStamina(130, 6),
	}

	Koma = &Player{
		StandingImage: images.KomaStanding,
		Description:   messages.DescKoma,
		animation:     NewStepAnimation(images.KomaAnimation, 5),
		jumpSe:        se.Jump,
		stamina:       NewStamina(160, 11),
	}

	Shishimaru = &Player{
		StandingImage: images.ShishimaruStanding,
		Description:   messages.DescShishimaru,
		animation:     NewStepAnimation(images.ShishimaruAnimation, 5),
		jumpSe:        se.Jump,
		stamina:       NewStamina(200, 17),
	}

	Selected = Kurona

	return nil
}

// Player is a player character.
type Player struct {
	position      view.Vector
	offset        image.Point
	rectangle     image.Rectangle
	blocked       bool
	StandingImage *ebiten.Image
	Description   string
	animation     *StepAnimation
	previous      State
	current       State
	stamina       *Stamina
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
	p.position = view.Vector{X: 0.0, Y: float64(charaHeights[0])}

	// set the edge rectangle with the position and image's rectangle.
	b := p.StandingImage.Bounds().Size()
	b.X -= 3
	b.Y -= 3
	p.rectangle = image.Rectangle{
		Min: image.Point{X: int(p.position.X), Y: int(p.position.Y)},
		Max: image.Point{X: int(p.position.X) + b.X, Y: int(p.position.Y) + b.Y},
	}

	return nil
}

// Start starts dash!
func (p *Player) Start() {
	if p.stamina.GetStamina() > 0 {
		p.current = Dash
	} else {
		p.current = Walk
	}
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
func (p *Player) Update() error {
	p.updateState()
	p.updateStamina()
	p.updatePosition()
	err := p.playSe()
	if err != nil {
		log.Println("failed to play SE:", err)
		return err
	}
	return nil
}

func (p *Player) updateState() {
	// TODO: ユーザーのキー入力、キャラクターの位置、障害物との衝突有無などを総合的に判断するStateManageがほしい。
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
		// update state by user input
		if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.GamepadAxis(0, 1) <= -0.5 {
			if !p.lanes.IsTop() {
				if p.lanes.Ascend() {
					p.previous = p.current
					p.current = Ascending
				}
			}
			return
		}
		if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.GamepadAxis(0, 1) >= 0.5 {
			if !p.lanes.IsBottom() {
				if p.lanes.Descend() {
					p.previous = p.current
					p.current = Descending
				}
			}
			return
		}

		// update state by stamina
		if p.stamina.GetStamina() <= 0 {
			p.current = Walk
		}

		// update state by blocked status
		if p.blocked {
			if p.current == Walk {
				return
			}
			p.previous = p.current
			p.current = Walk
		} else {
			if p.current == Dash {
				return
			}
			p.previous = p.current
			if p.stamina.GetStamina() > 0 {
				p.current = Dash
			} else {
				p.current = Walk
			}
		}
	}
}

// update stamina by current state
func (p *Player) updateStamina() {
	switch p.current {
	case Dash:
		p.stamina.Consumes(2)
	case Walk:
		p.stamina.Consumes(1)
	case Ascending:
		p.stamina.Consumes(1)
	case Descending:
		p.stamina.Consumes(1)
	default:
		// nothing to do
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
	// TODO: view.Rectangleなくしたい
	vel := image.Point{int(p.velocity.X), int(p.velocity.Y)}
	p.rectangle.Min = p.rectangle.Min.Add(vel)
	p.rectangle.Max = p.rectangle.Max.Add(vel)
}

// Draw draws the character image.
func (p *Player) Draw(screen *ebiten.Image) error {
	// TODO: ダッシュ中とか奥義中とか状態に応じて多少前後しつつ、ほぼ画面中央に描画したい
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(view.ScreenWidth/4, p.position.Y)
	p.offset.X = (int)(p.position.X - view.ScreenWidth/4)
	return screen.DrawImage(p.animation.GetCurrentFrame(), op)
}

func (p *Player) playSe() error {
	if p.previous != Ascending && p.current == Ascending {
		return p.jumpSe.Play()
	}
	return nil
}

// GetPosition return the current position of this player.
func (p *Player) GetPosition() view.Vector {
	return p.position
}

// GetOffset returns the offset to draw other filed parts.
func (p *Player) GetOffset() image.Point {
	return p.offset
}

// GetVelocity returns the velocity of this playable character.
func (p *Player) GetVelocity() view.Vector {
	return p.velocity
}

// GetStamina returns the stamina value fo this character.
func (p *Player) GetStamina() int {
	return p.stamina.GetStamina()
}

// GetRectangle returns the edge rentangle of this player.
func (p *Player) GetRectangle() image.Rectangle {
	return p.rectangle
}

// BeBlocked puts the player in a position where the path is blocked by an obstacle.
func (p *Player) BeBlocked(blocked bool) {
	p.blocked = blocked
}

// Close closes the inner resources.
func (p *Player) Close() error {
	return p.jumpSe.Close()
}
