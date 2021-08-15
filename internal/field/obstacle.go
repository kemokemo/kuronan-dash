package field

import (
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// Obstacle is the obstacles falling on the field.
type Obstacle interface {
	// SetHardness sets the hardness of this obstacle.
	SetHardness(hardness float64)

	// Attack attacks this obstacle.
	// The damage value reduces this obstacle's hardness.
	Attack(damage float64)

	// IsBroken returns whether this obstacle was broken.
	// The broken state means that the hardness is 0 or less.
	IsBroken() bool

	// IsCollided returns whether this obstacle is collided the arg.
	IsCollided(hr *view.HitRectangle) bool
}
