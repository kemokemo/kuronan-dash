package field

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// ScrollableObject is the object can move
type ScrollableObject interface {
	// Initialize initializes the object.
	//  args:
	//   img: the image to draw
	//   pos: the initial position
	//   kv: the factor to multiply the scroll speed when scrolling.
	Initialize(img *ebiten.Image, pos *view.Vector, kv float64)

	// Update updates the position and velocity of this object.
	//  args:
	//   scrollV: the velocity to scroll field parts.
	Update(scrollV *view.Vector)

	// Draw draws this object to the screen.
	Draw(screen *ebiten.Image)
}

type MovableObject interface {
	// SetVelocity sets velocity at which this object will move by itself.
	SetVelocity(vel *view.Vector)
}

// ScrollInfo is the info to initialize the ScrollableObject/
type ScrollInfo struct {
	img *ebiten.Image
	pos *view.Vector
	vel *view.Vector
}
