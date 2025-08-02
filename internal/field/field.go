package field

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// Field is the interface to draw the field.
type Field interface {
	// Initialize initializes the all of field parts.
	Initialize(lanes *Lanes, goalX float64)

	// Update updates the all field parts position with the scroll velocity.
	Update(scroll *view.Vector)

	// DrawFarther draws the field parts farther than the player from the user's point of view.
	DrawFarther(screen *ebiten.Image)

	// DrawCloser draws the field parts closer than the player from the user's point of view.
	DrawCloser(screen *ebiten.Image)

	// IsCollidedWithObstacles returns whether the r is collided with this item.
	IsCollidedWithObstacles(hr *view.HitRectangle) bool

	// EatFoods determines if there is a conflict between the player and the food.
	// If it hits, it returns the stamina and tension gained.
	EatFoods(hr *view.HitRectangle, soundPlayFlag bool) (stamina int, tension int)

	// AttackObstacles attacks obstacles.
	// If any obstacle is collided or broken, returns the number of collided or broken obstacles.
	AttackObstacles(hr *view.HitRectangle, power float64, soundPlayFlag bool) (int, int)
}

const (
	// FieldOffset is the offset to make it look like it's on the lane.
	FieldOffset = 1.8

	// rectOffset is the offset for field parts to check collision with player
	rectOffset = 2.0
)
