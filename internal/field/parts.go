package field

import (
	"github.com/kemokemo/kuronan-dash/internal/view"

	"github.com/hajimehoshi/ebiten/v2"
)

// Parts is the field parts.
type Parts struct {
	image    *ebiten.Image
	op       *ebiten.DrawImageOptions
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
	p.velocity = vel

	p.op = &ebiten.DrawImageOptions{}
	p.op.GeoM.Translate(pos.X, pos.Y)
}

// Update updates the position and velocity of this object.
//  args:
//   charaV: the velocity of the player character
func (p *Parts) Update(charaV view.Vector) {
	// Calculate relative speed with player only in horizontal direction
	p.op.GeoM.Translate(p.velocity.X-charaV.X, p.velocity.Y)
}

// Draw draws this object to the screen.
func (p *Parts) Draw(screen *ebiten.Image) {
	screen.DrawImage(p.image, p.op)
}
