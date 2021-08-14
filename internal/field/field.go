package field

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// Field is the interface to draw the field.
type Field interface {
	// Initialize initializes the all of field parts.
	Initialize()

	// Update updates the internal state and position with the player's velocity.
	Update(v view.Vector)

	// DrawFarther draws the field parts farther than the player from the user's point of view.
	DrawFarther(screen *ebiten.Image, pOffset image.Point)

	// DrawCloser draws the field parts closer than the player from the user's point of view.
	DrawCloser(screen *ebiten.Image, pOffset image.Point)

	// IsCollidedWithObstacles returns whether the r is collided with this item.
	IsCollidedWithObstacles(r image.Rectangle) bool

	// EatFoods determines if there is a conflict between the player and the food.
	// If it hits, it returns the stamina gained.
	EatFoods(r image.Rectangle) int
}
