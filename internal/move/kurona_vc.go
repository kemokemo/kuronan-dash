package move

import (
	"github.com/kemokemo/kuronan-dash/internal/view"
)

const (
	kuronaWalkMax             = 1.7
	kuronaDashMax             = 3.0
	kuronaDecelerateRate      = 0.9
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
func (kvc *KuronaVc) SetState(s State) {
	if kvc.currentState == s {
		kvc.elapsedX += elapsedStepX
		kvc.elapsedY += elapsedStepY
	} else {
		kvc.prevState = kvc.currentState
		kvc.currentState = s

		kvc.elapsedX = 1.0
		kvc.elapsedY = 0.0
	}
}

// GetVelocity returns the velocity to scroll the field parts and to update the character position.
func (kvc *KuronaVc) GetVelocity() (*view.Vector, *view.Vector, *view.Vector) {
	kvc.decideVbyState()
	kvc.updateVelocity()
	return kvc.scrollV, kvc.charaPosV, kvc.charaDrawV
}

func (kvc *KuronaVc) decideVbyState() {
	switch kvc.currentState {
	case Walk:
		kvc.decideVofWalk()
	case Dash:
		kvc.decideVofDash()
	case Ascending:
		kvc.deltaX = 0.6
		kvc.deltaY = kvc.jumpV0 + kvc.gravity*kvc.elapsedY
	case Descending:
		kvc.deltaX = 0.6
		if kvc.deltaY > 9.0 {
			kvc.deltaY = 9.0
		} else {
			kvc.deltaY = kvc.dropV0 + kvc.gravity*kvc.elapsedY
		}
	default:
		// Don't move
		kvc.deltaX = 0.0
		kvc.deltaY = 0.0
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
func (kvc *KuronaVc) updateVelocity() {
	kvc.charaPosV.X = kvc.deltaX
	kvc.charaPosV.Y = kvc.deltaY

	kvc.charaDrawV.X = 0.0
	kvc.charaDrawV.Y = kvc.deltaY

	kvc.scrollV.X = -kvc.deltaX
	kvc.scrollV.Y = 0.0
}
