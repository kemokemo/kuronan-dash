package character

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kemokemo/kuronan-dash/assets/se"
	"github.com/kemokemo/kuronan-dash/internal/anime"
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
	position   *view.Vector
	pDriver    *move.PlayerDriver
	scrollV    *view.Vector
	charaPosV  *view.Vector
	charaDrawV *view.Vector
	rect       *view.HitRectangle

	// Initialization is required before starting the stage.
	lanes    move.Lanes
	blocked  bool
	previous move.State
	current  move.State
	stamina  *Stamina
}

// InitializeWithLanesInfo sets the lanes information.
// The player can run on the lane or move between lanes based on the lane drawing height information received in the argument.
func (p *Player) InitializeWithLanesInfo(heights []float64) error {
	p.blocked = false
	p.previous = move.Walk
	p.current = move.Walk
	p.stamina.Initialize()

	cH := []float64{}
	_, h := p.StandingImage.Size()
	for i := 0; i < len(heights); i++ {
		cH = append(cH, heights[i]-float64(h))
	}

	p.lanes = move.Lanes{}
	err := p.lanes.SetHeights(cH)
	if err != nil {
		return err
	}

	// set the player at the top lane.
	initialY := float64(cH[0])
	offset := 3.0

	p.position = &view.Vector{X: view.DrawPosition, Y: initialY}
	p.charaPosV = &view.Vector{X: 0.0, Y: 0.0}
	p.scrollV = &view.Vector{X: 0.0, Y: 0.0}
	p.pDriver = move.NewPlayerDriver()
	p.op = &ebiten.DrawImageOptions{}
	p.op.GeoM.Translate(view.DrawPosition, initialY)

	w, h := p.StandingImage.Size()
	p.rect = view.NewHitRectangle(
		view.Vector{X: view.DrawPosition + offset, Y: initialY + offset},
		view.Vector{X: view.DrawPosition + float64(w) - offset, Y: initialY + float64(h) - offset})

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
func (p *Player) Update() error {
	err := p.updateState()
	if err != nil {
		return err
	}

	p.stamina.Consumes(p.current)

	p.pDriver.Update(p.current)
	p.scrollV, p.charaPosV, p.charaDrawV = p.pDriver.GetVelocity()

	p.animation.AddStep(p.charaPosV.X)
	p.position.Add(p.charaPosV)
	p.op.GeoM.Translate(p.charaDrawV.X, p.charaDrawV.Y)
	p.rect.Add(p.charaDrawV)

	return nil
}

func (p *Player) updateState() error {
	var err error
	// TODO: ユーザーのキー入力、キャラクターの位置、障害物との衝突有無などを総合的に判断するStateManageがほしい。
	switch p.current {
	case move.Pause:
		return err
	case move.Ascending, move.Descending:
		// TODO: I really want to go back to the previous movement before the ascending or descending motion.
		if p.lanes.IsReachedTarget(p.position.Y, p.charaPosV.Y) {
			p.current = move.Dash
		}
	default:
		// update state by user input
		if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.GamepadAxis(0, 1) <= -0.5 {
			if !p.lanes.IsTop() {
				if p.lanes.Ascend() {
					p.previous = p.current
					p.current = move.Ascending
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
					p.current = move.Descending
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
			p.current = move.Walk
		}

		// update state by blocked status
		if p.blocked {
			if p.current == move.Walk {
				return err
			}
			p.previous = p.current
			p.current = move.Walk
		} else {
			if p.current == move.Dash {
				return err
			}
			p.previous = p.current
			if p.stamina.GetStamina() > 0 {
				p.current = move.Dash
			} else {
				p.current = move.Walk
			}
		}
	}
	return err
}

// Draw draws the character image.
func (p *Player) Draw(screen *ebiten.Image) {
	// TODO: ダッシュ中とか奥義中とか状態に応じて多少前後しつつ、ほぼ画面中央に描画したい
	screen.DrawImage(p.animation.GetCurrentFrame(), p.op)
}

// GetPosition return the current position of this player.
func (p *Player) GetPosition() *view.Vector {
	return p.position
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
func (p *Player) BeBlocked(blocked bool) {
	p.blocked = blocked
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
