package field

import (
	"github.com/kemokemo/kuronan-dash/internal/view"

	"github.com/hajimehoshi/ebiten"
)

// Rock is the interface of the field part.
type Rock struct {
	image    *ebiten.Image
	position view.Vector
	velocity view.Vector
}

// Initialize initializes the object.
//  args:
//   img: the image to draw
//   pos: the initial position
//   vel: the velocity to move this object
func (r *Rock) Initialize(img *ebiten.Image, pos, vel view.Vector) {
	r.image = img
	r.position = pos
	r.velocity = vel
}

// Update updates the position and velocity of this object.
//  args:
//   charaV: the velocity of the player character
func (r *Rock) Update(charaV view.Vector) {
	r.position = r.position.Add(r.velocity)
	// Calculate relative speed with player only in horizontal direction
	r.position.X -= charaV.X
}

// Draw draws this object to the screen.
func (r *Rock) Draw(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(r.position.X, r.position.Y)

	return screen.DrawImage(r.image, op)
}
