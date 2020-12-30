package field

import (
	"image"

	"github.com/kemokemo/kuronan-dash/internal/view"

	"github.com/hajimehoshi/ebiten/v2"
)

// Parts is the field parts.
type Parts struct {
	image    *ebiten.Image
	position view.Vector
	velocity view.Vector
}

// Initialize initializes the object.
//  args:
//   img: the image to draw
//   pos: the initial position
//   vel: the velocity to move this object
func (p *Parts) Initialize(img *ebiten.Image, pos, vel view.Vector) {
	p.image = img
	p.position = pos
	p.velocity = vel
}

// Update updates the position and velocity of this object.
//  args:
//   charaV: the velocity of the player character
func (p *Parts) Update(charaV view.Vector) {
	p.position = p.position.Add(p.velocity)
	// Calculate relative speed with player only in horizontal direction
	p.position.X -= charaV.X
}

// Draw draws this object to the screen.
func (p *Parts) Draw(screen *ebiten.Image, offset image.Point) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.position.X-float64(offset.X), p.position.Y-float64(offset.Y))
	screen.DrawImage(p.image, op)
}
