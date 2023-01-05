package move

import "github.com/kemokemo/kuronan-dash/internal/view"

// 独楽:
// 移動速度はやや遅め。その代わり、障害物に当たっても速度が落ちにくい。
const (
	komaWalkMax               = 1.2
	komaDashMax               = 2.2
	komaSpWalkMax             = 3.0
	komaSpMax                 = 4.0
	komaDecelerateRate        = 0.2
	komaInitialVelocityWalk   = 0.03
	komaInitialVelocityDash   = 0.1
	komaInitialVelocitySpWalk = 0.12
	komaInitialVelocitySp     = 0.15
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

func (vc *KomaVc) SetState(s State) {
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
func (vc *KomaVc) GetVelocity() (*view.Vector, *view.Vector, *view.Vector) {
	vc.decideVbyState()
	vc.updateVelocity()
	return vc.scrollV, vc.charaPosV, vc.charaDrawV
}

func (vc *KomaVc) decideVbyState() {
	switch vc.currentState {
	case Walk:
		vc.decideVofWalk()
	case Dash:
		vc.decideVofDash()
	case SkillDash:
		vc.decideVofSkillDash()
	case SkillWalk:
		vc.decideVofSkillWalk()
	case Ascending, SkillAscending:
		if vc.currentState == Ascending {
			vc.deltaX = 0.6
		} else if vc.currentState == SkillAscending {
			vc.deltaX = 1.2
		}
		vc.deltaY = vc.jumpV0 + vc.gravity*vc.elapsedY
	case Descending, SkillDescending:
		if vc.currentState == Descending {
			vc.deltaX = 0.6
		} else if vc.currentState == SkillDescending {
			vc.deltaX = 1.2
		}
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

func (vc *KomaVc) decideVofSkillDash() {
	vc.deltaX += komaInitialVelocitySp * vc.elapsedX
	if vc.deltaX > komaSpMax {
		vc.deltaX = komaSpMax
	}
	vc.deltaY = 0.0
}

func (vc *KomaVc) decideVofSkillWalk() {
	vc.deltaX += komaInitialVelocitySpWalk * vc.elapsedX
	if vc.deltaX > komaSpWalkMax {
		vc.deltaX = komaSpWalkMax
	}
	vc.deltaY = 0.0
}

// updateVelocity updates all velocities. Please pass me the data for charaPosV.
func (vc *KomaVc) updateVelocity() {
	vc.charaPosV.X = vc.deltaX
	vc.charaPosV.Y = vc.deltaY

	vc.charaDrawV.X = 0.0
	vc.charaDrawV.Y = vc.deltaY

	vc.scrollV.X = -vc.deltaX
	vc.scrollV.Y = 0.0
}
