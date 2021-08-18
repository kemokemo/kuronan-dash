package move

import "github.com/kemokemo/kuronan-dash/internal/view"

// NewKomaVc returns a new VelocityController for Kurona.
func NewKomaVc() *KomaVc {
	return &KomaVc{
		scrollV:    &view.Vector{X: 0.0, Y: 0.0},
		charaPosV:  &view.Vector{X: 0.0, Y: 0.0},
		charaDrawV: &view.Vector{X: 0.0, Y: 0.0},
		gravity:    1.2,
		jumpV0:     -9.7,
		dropV0:     0.2,
	}
}

// KomaVc is VelocityController of Kurona. Please create via 'NewKomaVc' method.
type KomaVc struct {
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

func (kvc *KomaVc) SetState(s State) {
	kvc.prevState = kvc.currentState
	kvc.currentState = s

	if kvc.prevState == s {
		kvc.elapsed += elapsedStep
	} else {
		kvc.elapsed = 0.0
	}
}

// GetVelocity returns the velocity to scroll the field parts and to update the character position.
func (mvc *KomaVc) GetVelocity() (*view.Vector, *view.Vector, *view.Vector) {
	mvc.decideVbyState()
	mvc.updateVelocity()
	return mvc.scrollV, mvc.charaPosV, mvc.charaDrawV
}

// TODO: キャラクターごとに個性を出す部分
// ダッシュから歩きに変わる時のX方向の速度の落ち方、上がり方もキャラごとに特性が出せたら素敵。
func (mvc *KomaVc) decideVbyState() {
	switch mvc.currentState {
	case Walk:
		mvc.deltaX = 1.0
		mvc.deltaY = 0.0
	case Dash:
		mvc.deltaX = 2.0
		mvc.deltaY = 0.0
	case Ascending:
		mvc.deltaX = 0.6
		mvc.deltaY = mvc.jumpV0 + mvc.gravity*mvc.elapsed
	case Descending:
		mvc.deltaX = 0.6
		mvc.deltaY = mvc.dropV0 + mvc.gravity*mvc.elapsed
	default:
		// Don't move
		mvc.deltaX = 0.0
		mvc.deltaY = 0.0
	}
}

// updateVelocity updates all velocities. Please pass me the data for charaPosV.
func (mvc *KomaVc) updateVelocity() {
	mvc.charaPosV.X = mvc.deltaX
	mvc.charaPosV.Y = mvc.deltaY

	mvc.charaDrawV.X = 0.0
	mvc.charaDrawV.Y = mvc.deltaY

	mvc.scrollV.X = -mvc.deltaX
	mvc.scrollV.Y = 0.0
}
