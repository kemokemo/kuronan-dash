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
	v0       *view.Vector
}

// Initialize initializes the object.
//  args:
//   img: the image to draw
//   pos: the initial position
//   vel: the velocity to move this object
func (p *Parts) Initialize(img *ebiten.Image, pos, vel *view.Vector) {
	p.image = img
	p.v0 = &view.Vector{X: vel.X, Y: vel.Y}
	p.op = &ebiten.DrawImageOptions{}
	p.op.GeoM.Translate(pos.X, pos.Y)
}

// Update updates the position and velocity of this object.
//  args:
//   scrollV: the velocity to scroll this field parts.
func (p *Parts) Update(scrollV *view.Vector) {
	p.op.GeoM.Translate(p.v0.X+scrollV.X, p.v0.Y+scrollV.Y)
}

// Draw draws this object to the screen.
func (p *Parts) Draw(screen *ebiten.Image) {
	screen.DrawImage(p.image, p.op)
}
