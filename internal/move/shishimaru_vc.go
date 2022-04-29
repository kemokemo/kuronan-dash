package move

import "github.com/kemokemo/kuronan-dash/internal/view"

// 獅子丸:
// 黒菜と独楽の中間性能。バランスがいいよ。
const (
	shishimaruWalkMax             = 1.5
	shishimaruDashMax             = 2.6
	shishimaruDecelerateRate      = 0.8
	shishimaruInitialVelocityWalk = 0.07
	shishimaruInitialVelocityDash = 0.2
)

// NewShishimaruVc returns a new VelocityController for Kurona.
func NewShishimaruVc() *ShishimaruVc {
	return &ShishimaruVc{
		scrollV:    &view.Vector{X: 0.0, Y: 0.0},
		charaPosV:  &view.Vector{X: 0.0, Y: 0.0},
		charaDrawV: &view.Vector{X: 0.0, Y: 0.0},
		gravity:    1.2,
		jumpV0:     -10.0,
		dropV0:     0.3,
	}
}

// ShishimaruVc is VelocityController of Kurona. Please create via 'NewShishimaruVc' method.
type ShishimaruVc struct {
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

func (svc *ShishimaruVc) SetState(s State) {
	if svc.currentState == s {
		svc.elapsedX += elapsedStepX
		svc.elapsedY += elapsedStepY
	} else {
		svc.prevState = svc.currentState
		svc.currentState = s

		svc.elapsedX = 1.0
		svc.elapsedY = 0.0
	}
}

// GetVelocity returns the velocity to scroll the field parts and to update the character position.
func (svc *ShishimaruVc) GetVelocity() (*view.Vector, *view.Vector, *view.Vector) {
	svc.decideVbyState()
	svc.updateVelocity()
	return svc.scrollV, svc.charaPosV, svc.charaDrawV
}

func (svc *ShishimaruVc) decideVbyState() {
	switch svc.currentState {
	case Walk:
		svc.decideVofWalk()
	case Dash:
		svc.decideVofDash()
	case Ascending:
		svc.deltaX = 0.6
		svc.deltaY = svc.jumpV0 + svc.gravity*svc.elapsedY
	case Descending:
		svc.deltaX = 0.6
		if svc.deltaY > 9.0 {
			svc.deltaY = 9.0
		} else {
			svc.deltaY = svc.dropV0 + svc.gravity*svc.elapsedY
		}
	default:
		// Don't move
		svc.deltaX = 0.0
		svc.deltaY = 0.0
	}
}

// ダッシュから歩きに変わる時のX方向の速度の落ち方、上がり方もキャラごとに特性を出す。
func (vc *ShishimaruVc) decideVofWalk() {
	if vc.prevState == Dash && vc.deltaX > shishimaruWalkMax {
		// 減速処理
		vc.deltaX -= shishimaruDecelerateRate * vc.elapsedX
	} else {
		vc.deltaX += shishimaruInitialVelocityWalk * vc.elapsedX
		if vc.deltaX > shishimaruWalkMax {
			vc.deltaX = shishimaruWalkMax
		}
	}

	vc.deltaY = 0.0
}

func (vc *ShishimaruVc) decideVofDash() {
	vc.deltaX += shishimaruInitialVelocityDash * vc.elapsedX
	if vc.deltaX > shishimaruDashMax {
		vc.deltaX = shishimaruDashMax
	}
	vc.deltaY = 0.0
}

// updateVelocity updates all velocities. Please pass me the data for charaPosV.
func (svc *ShishimaruVc) updateVelocity() {
	svc.charaPosV.X = svc.deltaX
	svc.charaPosV.Y = svc.deltaY

	svc.charaDrawV.X = 0.0
	svc.charaDrawV.Y = svc.deltaY

	svc.scrollV.X = -svc.deltaX
	svc.scrollV.Y = 0.0
}
