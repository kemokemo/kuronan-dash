package move

import "github.com/kemokemo/kuronan-dash/internal/view"

// 独楽:
// 移動速度はやや遅め。その代わり、障害物に当たっても速度が落ちにくい。
const (
	komaWalkMax             = 1.2
	komaDashMax             = 2.2
	komaDecelerateRate      = 0.2
	komaInitialVelocityWalk = 0.03
	komaInitialVelocityDash = 0.1
)

// NewKomaVc returns a new VelocityController for Kurona.
func NewKomaVc() *KomaVc {
	return &KomaVc{
		scrollV:    &view.Vector{X: 0.0, Y: 0.0},
		charaPosV:  &view.Vector{X: 0.0, Y: 0.0},
		charaDrawV: &view.Vector{X: 0.0, Y: 0.0},
		gravity:    1.2,
		jumpV0:     -9.7,
		dropV0:     0.5,
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
	elapsedX       float64
	elapsedY       float64
	deltaX, deltaY float64
}

func (kvc *KomaVc) SetState(s State) {
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
func (mvc *KomaVc) GetVelocity() (*view.Vector, *view.Vector, *view.Vector) {
	mvc.decideVbyState()
	mvc.updateVelocity()
	return mvc.scrollV, mvc.charaPosV, mvc.charaDrawV
}

func (mvc *KomaVc) decideVbyState() {
	switch mvc.currentState {
	case Walk:
		mvc.decideVofWalk()
	case Dash:
		mvc.decideVofDash()
	case Ascending:
		mvc.deltaX = 0.6
		mvc.deltaY = mvc.jumpV0 + mvc.gravity*mvc.elapsedY
	case Descending:
		mvc.deltaX = 0.6
		if mvc.deltaY > 9.0 {
			mvc.deltaY = 9.0
		} else {
			mvc.deltaY = mvc.dropV0 + mvc.gravity*mvc.elapsedY
		}
	default:
		// Don't move
		mvc.deltaX = 0.0
		mvc.deltaY = 0.0
	}
}

// ダッシュから歩きに変わる時のX方向の速度の落ち方、上がり方もキャラごとに特性を出す。
func (vc *KomaVc) decideVofWalk() {
	if vc.prevState == Dash && vc.deltaX > komaWalkMax {
		// 減速処理
		vc.deltaX -= komaDecelerateRate * vc.elapsedX
	} else {
		vc.deltaX += komaInitialVelocityWalk * vc.elapsedX
		if vc.deltaX > komaWalkMax {
			vc.deltaX = komaWalkMax
		}
	}

	vc.deltaY = 0.0
}

func (vc *KomaVc) decideVofDash() {
	vc.deltaX += komaInitialVelocityDash * vc.elapsedX
	if vc.deltaX > komaDashMax {
		vc.deltaX = komaDashMax
	}
	vc.deltaY = 0.0
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
