package move

import (
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// 黒菜:
// 移動速度が速い。その代わり、障害物に当るとすぐに速度が落ちちゃう。
const (
	kuronaWalkMax             = 1.7
	kuronaDashMax             = 3.0
	kuronaDecelerateRate      = 1.2
	kuronaInitialVelocityWalk = 0.1
	kuronaInitialVelocityDash = 0.35
)

// NewKuronaVc returns a new VelocityController for Kurona.
func NewKuronaVc() *KuronaVc {
	return &KuronaVc{
		scrollV:    &view.Vector{X: 0.0, Y: 0.0},
		charaPosV:  &view.Vector{X: 0.0, Y: 0.0},
		charaDrawV: &view.Vector{X: 0.0, Y: 0.0},
		gravity:    1.2,
		jumpV0:     -10.3,
		dropV0:     0.2,
	}
}

// KuronaVc is VelocityController of Kurona. Please create via 'NewKuronaVc' method.
type KuronaVc struct {
	scrollV        *view.Vector
	charaPosV      *view.Vector
	charaDrawV     *view.Vector
	gravity        float64
	jumpV0         float64
	dropV0         float64
	currentState   State
	prevState      State
	elapsedX       float64
	elapsedY       float64
	deltaX, deltaY float64
}

// Only when state changed, prev and current states are updated.
func (vc *KuronaVc) SetState(s State) {
	if vc.currentState == s {
		vc.elapsedX += elapsedStepX
		vc.elapsedY += elapsedStepY
	} else {
		vc.prevState = vc.currentState
		vc.currentState = s

		vc.elapsedX = 1.0
		vc.elapsedY = 0.0
	}
}

// GetVelocity returns the velocity to scroll the field parts and to update the character position.
func (vc *KuronaVc) GetVelocity() (*view.Vector, *view.Vector, *view.Vector) {
	vc.decideVbyState()
	vc.updateVelocity()
	return vc.scrollV, vc.charaPosV, vc.charaDrawV
}

func (vc *KuronaVc) decideVbyState() {
	switch vc.currentState {
	case Walk:
		vc.decideVofWalk()
	case Dash:
		vc.decideVofDash()
	case Ascending:
		vc.deltaX = 0.6
		vc.deltaY = vc.jumpV0 + vc.gravity*vc.elapsedY
	case Descending:
		vc.deltaX = 0.6
		if vc.deltaY > 9.0 {
			vc.deltaY = 9.0
		} else {
			vc.deltaY = vc.dropV0 + vc.gravity*vc.elapsedY
		}
	default:
		// Don't move
		vc.deltaX = 0.0
		vc.deltaY = 0.0
	}
}

// ダッシュから歩きに変わる時のX方向の速度の落ち方、上がり方もキャラごとに特性を出す。
func (vc *KuronaVc) decideVofWalk() {
	if vc.prevState == Dash && vc.deltaX > kuronaWalkMax {
		// 減速処理
		vc.deltaX -= kuronaDecelerateRate * vc.elapsedX
	} else {
		vc.deltaX += kuronaInitialVelocityWalk * vc.elapsedX
		if vc.deltaX > kuronaWalkMax {
			vc.deltaX = kuronaWalkMax
		}
	}

	vc.deltaY = 0.0
}

func (vc *KuronaVc) decideVofDash() {
	vc.deltaX += kuronaInitialVelocityDash * vc.elapsedX
	if vc.deltaX > kuronaDashMax {
		vc.deltaX = kuronaDashMax
	}
	vc.deltaY = 0.0
}

// updateVelocity updates all velocities. Please pass me the data for charaPosV.
func (vc *KuronaVc) updateVelocity() {
	vc.charaPosV.X = vc.deltaX
	vc.charaPosV.Y = vc.deltaY

	vc.charaDrawV.X = 0.0
	vc.charaDrawV.Y = vc.deltaY

	vc.scrollV.X = -vc.deltaX
	vc.scrollV.Y = 0.0
}
