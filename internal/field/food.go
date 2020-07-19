package field

import "image"

// Food is the interface for filed food item. The character can eat them and restore stamina.
type Food interface {
	// IsCollided returns whether this obstacle is collided the arg.
	IsCollided(r image.Rectangle) bool

	// Eat eats this food. This func reteruns the value to restore character's stamina.
	Eat() int
}
