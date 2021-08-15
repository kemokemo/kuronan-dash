package move

import "github.com/kemokemo/kuronan-dash/internal/view"

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
	prevState      State
	elapsed        float64
	deltaX, deltaY float64
}

// GetVelocity returns the velocity to scroll the field parts and to update the character position.
func (svc *ShishimaruVc) GetVelocity(s State) (*view.Vector, *view.Vector, *view.Vector) {
	if svc.prevState == s {
		svc.elapsed += elapsedStep
	} else {
		svc.prevState = s
		svc.elapsed = 0.0
	}

	svc.decideVbyState(s)
	svc.updateVelocity()

	return svc.scrollV, svc.charaPosV, svc.charaDrawV
}

// TODO: キャラクターごとに個性を出す部分
// ダッシュから歩きに変わる時のX方向の速度の落ち方、上がり方もキャラごとに特性が出せたら素敵。
func (svc *ShishimaruVc) decideVbyState(s State) {
	switch s {
	case Walk:
		svc.deltaX = 1.0
		svc.deltaY = 0.0
	case Dash:
		svc.deltaX = 2.0
		svc.deltaY = 0.0
	case Ascending:
		svc.deltaX = 0.6
		svc.deltaY = svc.jumpV0 + svc.gravity*svc.elapsed
	case Descending:
		svc.deltaX = 0.6
		svc.deltaY = svc.dropV0 + svc.gravity*svc.elapsed
	default:
		// Don't move
		svc.deltaX = 0.0
		svc.deltaY = 0.0
	}
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
