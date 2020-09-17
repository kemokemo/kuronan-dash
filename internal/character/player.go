package character

import (
	"fmt"
	"image"

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
		dropSe:        se.Drop,
		stamina:       NewStamina(130, 6),
	}

	Koma = &Player{
		StandingImage: images.KomaStanding,
		Description:   messages.DescKoma,
		animation:     NewStepAnimation(images.KomaAnimation, 5),
		jumpSe:        se.Jump,
		dropSe:        se.Drop,
		stamina:       NewStamina(160, 11),
	}

	Shishimaru = &Player{
		StandingImage: images.ShishimaruStanding,
		Description:   messages.DescShishimaru,
		animation:     NewStepAnimation(images.ShishimaruAnimation, 5),
		jumpSe:        se.Jump,
		dropSe:        se.Drop,
		stamina:       NewStamina(200, 17),
	}

	Selected = Kurona

	return nil
}

// Player is a player character.
type Player struct {
	// Specified at creation and not changed
	StandingImage *ebiten.Image
	Description   string
	animation     *StepAnimation
	jumpSe        *se.Player
	dropSe        *se.Player

	// Update each time based on the internal status and other information
	position  view.Vector
	velocity  view.Vector
	rectangle image.Rectangle
	offset    image.Point

	// Initialization is required before starting the stage.
	lanes    Lanes
	blocked  bool
	previous State
	current  State
	stamina  *Stamina
}

// InitilizeWithLanesInfo sets the lanes information.
// The player can run on the lane or move between lanes based on the lane drawing height information received in the argument.
func (p *Player) InitilizeWithLanesInfo(heights []float64) error {
	p.blocked = false
	p.previous = Walk
	p.current = Walk
	p.stamina.Initialize()

	cH := []float64{}
	_, h := p.StandingImage.Size()
	for i := 0; i < len(heights); i++ {
		cH = append(cH, heights[i]-float64(h))
	}

	p.lanes = Lanes{}
	err := p.lanes.SetHeights(cH)
	if err != nil {
		return err
	}

	// set the player at the top lane.
	p.position = view.Vector{X: 0.0, Y: float64(cH[0])}

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

// Start starts playing.
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
func (p *Player) Update() error {
	err := p.updateState()
	if err != nil {
		return err
	}
	p.updateStamina()
	p.updatePosition()

	return nil
}

func (p *Player) updateState() error {
	var err error
	// TODO: ユーザーのキー入力、キャラクターの位置、障害物との衝突有無などを総合的に判断するStateManageがほしい。
	switch p.current {
	case Pause:
		return err
	case Ascending, Descending:
		// TODO: I really want to go back to the previous movement before the ascending or descending motion.
		if p.lanes.IsReachedTarget(p.position.Y) {
			p.current = Dash
		}
	default:
		// update state by user input
		if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.GamepadAxis(0, 1) <= -0.5 {
			if !p.lanes.IsTop() {
				if p.lanes.Ascend() {
					p.previous = p.current
					p.current = Ascending
					err = p.jumpSe.Play()
					if err != nil {
						err = fmt.Errorf("failed to play se: %v", err)
					}
				}
			}
			return err
		}
		if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.GamepadAxis(0, 1) >= 0.5 {
			if !p.lanes.IsBottom() {
				if p.lanes.Descend() {
					p.previous = p.current
					p.current = Descending
					err = p.dropSe.Play()
					if err != nil {
						err = fmt.Errorf("failed to play se: %v", err)
					}
				}
			}
			return err
		}

		// update state by stamina
		if p.stamina.GetStamina() <= 0 {
			p.current = Walk
		}

		// update state by blocked status
		if p.blocked {
			if p.current == Walk {
				return err
			}
			p.previous = p.current
			p.current = Walk
		} else {
			if p.current == Dash {
				return err
			}
			p.previous = p.current
			if p.stamina.GetStamina() > 0 {
				p.current = Dash
			} else {
				p.current = Walk
			}
		}
	}
	return err
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

// Eat eats foods and restore stamina value by argument value.
func (p *Player) Eat(stamina int) {
	p.stamina.Restore(stamina)
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
