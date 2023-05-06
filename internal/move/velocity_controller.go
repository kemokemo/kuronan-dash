package move

import "github.com/kemokemo/kuronan-dash/internal/view"

const (
	elapsedStepX = 1.0
	elapsedStepY = 0.2
)

// VelocityController calcurate velocity of characters.
type VelocityController interface {
	// SetState sets the current state to this object.
	SetState(s State)

	// GetVelocity returns the velocities to scroll screen, move character's internal position, move character's drawing position.
	GetVelocity() (scrollV, charaPosV, charaDrawV *view.Vector)

	// GetDashMax returns the max velocity for dash state.
	GetDashMax() float64
}
