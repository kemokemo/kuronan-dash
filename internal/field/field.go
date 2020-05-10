package field

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// Field is the interface to draw the field.
type Field interface {
	// Initialize initializes the all of field parts.
	Initialize()

	// Update updates the internal state and position with the player's velocity.
	Update(v view.Vector)

	// DrawFarther draws the field parts farther than the player from the user's point of view.
	DrawFarther(screen *ebiten.Image) error

	// DrawCloser draws the field parts closer than the player from the user's point of view.
	DrawCloser(screen *ebiten.Image) error

	// IsCollidedWithObstacles returns whether the r is collided with this item.
	IsCollidedWithObstacles(r image.Rectangle) bool
}
