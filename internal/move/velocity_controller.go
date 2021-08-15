package move

import "github.com/kemokemo/kuronan-dash/internal/view"

const (
	elapsedStep = 0.2
)

// VelocityController calcurate velocity of characters.
type VelocityController interface {
	// GetVelocity returns the velocities to scroll screen, move character's internal position, move character's drawing position.
	GetVelocity(s State) (scrollV, charaPosV, charaDrawV *view.Vector)
}
