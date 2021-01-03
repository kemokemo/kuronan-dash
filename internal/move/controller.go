package move

import (
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// Controller is the interface to control movement of the characters.
type Controller interface {
	// SetPosAndAccel sets the initial position and the accel of this controller.
	SetPosAndAccel(initial, accel view.Vector)
	// Update updates the velocity.
	Update(s State)
	// GetVelocity returns the velocity of the controller.
	GetVelocity() view.Vector
}
