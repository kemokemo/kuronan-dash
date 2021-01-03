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
	//   vel: the velocity to move this object
	Initialize(img *ebiten.Image, pos, vel view.Vector)

	// Update updates the position and velocity of this object.
	//  args:
	//   charaV: the velocity of the player character
	Update(charaV view.Vector)

	// Draw draws this object to the screen.
	Draw(screen *ebiten.Image)
}
