package move

import (
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// NewKuronaVc returns a new VelocityController for Kurona.
func NewKuronaVc() *KuronaVc {
	return &KuronaVc{
		scrollV:    &view.Vector{X: 0.0, Y: 0.0},
		charaPosV:  &view.Vector{X: 0.0, Y: 0.0},
		charaDrawV: &view.Vector{X: 0.0, Y: 0.0},
		gravity:    1.2,
		jumpV0:     -10.3,
		dropV0:     0.5,
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
	elapsed        float64
	deltaX, deltaY float64
}

func (kvc *KuronaVc) SetState(s State) {
	kvc.prevState = kvc.currentState
	kvc.currentState = s

	if kvc.prevState == s {
		kvc.elapsed += elapsedStep
	} else {
		kvc.elapsed = 0.0
	}
}

// GetVelocity returns the velocity to scroll the field parts and to update the character position.
func (kvc *KuronaVc) GetVelocity() (*view.Vector, *view.Vector, *view.Vector) {
	kvc.decideVbyState()
	kvc.updateVelocity()
	return kvc.scrollV, kvc.charaPosV, kvc.charaDrawV
}

// TODO: キャラクターごとに個性を出す部分
// ダッシュから歩きに変わる時のX方向の速度の落ち方、上がり方もキャラごとに特性が出せたら素敵。
func (kvc *KuronaVc) decideVbyState() {
	switch kvc.currentState {
	case Walk:
		kvc.deltaX = 1.0
		kvc.deltaY = 0.0
	case Dash:
		kvc.deltaX = 2.0
		kvc.deltaY = 0.0
	case Ascending:
		kvc.deltaX = 0.6
		kvc.deltaY = kvc.jumpV0 + kvc.gravity*kvc.elapsed
	case Descending:
		kvc.deltaX = 0.6
		kvc.deltaY = kvc.dropV0 + kvc.gravity*kvc.elapsed
		// todo: 試験的に上限の落下速度を導入
		if kvc.deltaY > 9.0 {
			kvc.deltaY = 9.0
		}
	default:
		// Don't move
		kvc.deltaX = 0.0
		kvc.deltaY = 0.0
	}
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
