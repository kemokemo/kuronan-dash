package move

import "github.com/kemokemo/kuronan-dash/internal/view"

// memo:
// When you add implementation of Driver interface to the PlayerDriver, Let's rock below command.
// $ impl 'pd *PlayerDriver' move.Driver >> player_driver.go

// todo: キャラごとに異なる初速度、加速度特性にしたいよ。
// ダッシュから歩きに変わる時のX方向の速度の落ち方、上がり方もキャラごとに特性が出せたら素敵。
var (
	gravity = 1.3
	jumpV0  = -7.2
	dropV0  = 1.2
)

// NewPlayerDriver returns a new PlayerDriver.
func NewPlayerDriver() *PlayerDriver {
	return &PlayerDriver{
		scrollV:    &view.Vector{X: 0.0, Y: 0.0},
		charaPosV:  &view.Vector{X: 0.0, Y: 0.0},
		charaDrawV: &view.Vector{X: 0.0, Y: 0.0},
	}
}

// PlayerDriver is various velocity controller to drive the player and the field parts.
type PlayerDriver struct {
	scrollV    *view.Vector
	charaPosV  *view.Vector
	charaDrawV *view.Vector
	elapsed    float64
	prevState  State
}

// Update updates the velocity.
func (pd *PlayerDriver) Update(s State) {
	if pd.prevState == s {
		pd.elapsed += 0.1
	} else {
		pd.prevState = s
		pd.elapsed = 0.0
	}

	switch s {
	case Walk:
		pd.updateVelocity(1.0, 0.0)
	case Dash:
		pd.updateVelocity(2.0, 0.0)
	case Ascending:
		pd.updateVelocity(0.6, jumpV0+gravity*pd.elapsed)
	case Descending:
		pd.updateVelocity(0.6, dropV0+gravity*pd.elapsed)
	default:
		// Don't move
		pd.updateVelocity(0.0, 0.0)
	}
}

// updateVelocity updates all velocities. Please pass me the data for charaPosV.
func (pd *PlayerDriver) updateVelocity(x, y float64) {
	pd.charaPosV.X = x
	pd.charaPosV.Y = y

	pd.charaDrawV.X = 0.0
	pd.charaDrawV.Y = y

	pd.scrollV.X = -x
	pd.scrollV.Y = 0.0
}

// GetVelocity returns the velocity to scroll the field parts and to update the character position.
func (pd *PlayerDriver) GetVelocity() (scrollV, charaPosV, charaDrawV *view.Vector) {
	return pd.scrollV, pd.charaPosV, pd.charaDrawV
}
