package field

import (
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// Food is the interface for filed food item. The character can eat them and restore stamina.
type Food interface {
	// IsCollided returns whether this obstacle is collided the arg.
	IsCollided(*view.HitRectangle) bool

	// Eat eats this food. This func returns the value to restore character's stamina and tension.
	Eat(soundPlayFlag bool) (stamina int, tension int)
}
