package move

import "github.com/kemokemo/kuronan-dash/internal/view"

// 獅子丸:
// 黒菜と独楽の中間性能。バランスがいいよ。
const (
	shishimaruWalkMax               = 1.5
	shishimaruDashMax               = 2.6
	shishimaruSpWalkMax             = 3.5
	shishimaruSpMax                 = 4.7
	shishimaruDecelerateRate        = 0.8
	shishimaruInitialVelocityWalk   = 0.07
	shishimaruInitialVelocityDash   = 0.2
	shishimaruInitialVelocitySpWalk = 0.22
	shishimaruInitialVelocitySp     = 0.25
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

func (vc *ShishimaruVc) SetState(s State) {
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
func (vc *ShishimaruVc) GetVelocity() (*view.Vector, *view.Vector, *view.Vector) {
	vc.decideVbyState()
	vc.updateVelocity()
	return vc.scrollV, vc.charaPosV, vc.charaDrawV
}

func (vc *ShishimaruVc) decideVbyState() {
	switch vc.currentState {
	case Walk:
		vc.decideVofWalk()
	case Dash:
		vc.decideVofDash()
	case SkillDash:
		vc.decideVofSpDash()
	case SkillWalk:
		vc.decideVofSpWalk()
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

func (vc *ShishimaruVc) decideVofSpDash() {
	vc.deltaX += shishimaruInitialVelocitySp * vc.elapsedX
	if vc.deltaX > shishimaruSpMax {
		vc.deltaX = shishimaruSpMax
	}
	vc.deltaY = 0.0
}

func (vc *ShishimaruVc) decideVofSpWalk() {
	vc.deltaX += shishimaruInitialVelocitySpWalk * vc.elapsedX
	if vc.deltaX > shishimaruSpWalkMax {
		vc.deltaX = shishimaruSpWalkMax
	}
	vc.deltaY = 0.0
}

// updateVelocity updates all velocities. Please pass me the data for charaPosV.
func (vc *ShishimaruVc) updateVelocity() {
	vc.charaPosV.X = vc.deltaX
	vc.charaPosV.Y = vc.deltaY

	vc.charaDrawV.X = 0.0
	vc.charaDrawV.Y = vc.deltaY

	vc.scrollV.X = -vc.deltaX
	vc.scrollV.Y = 0.0
}
