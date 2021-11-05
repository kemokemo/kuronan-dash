package field

import (
	"github.com/kemokemo/kuronan-dash/internal/view"

	"github.com/hajimehoshi/ebiten/v2"
)

// Parts is the field parts.
type Parts struct {
	image *ebiten.Image
	op    *ebiten.DrawImageOptions
	kv    float64
}

// Initialize initializes the object.
//  args:
//   img: the image to draw
//   pos: the initial position
//   vel: the velocity at which this object moves autonomously.
//   kv: the factor to multiply the scroll speed when scrolling.
func (p *Parts) Initialize(img *ebiten.Image, pos *view.Vector, kv float64) {
	p.image = img
	p.kv = kv
	p.op = &ebiten.DrawImageOptions{}
	p.op.GeoM.Translate(pos.X, pos.Y)
}

// Update updates the position and velocity of this object.
//  args:
//   scrollV: the velocity to scroll this field parts.
func (p *Parts) Update(scrollV *view.Vector) {
	p.op.GeoM.Translate(p.kv*scrollV.X, p.kv*scrollV.Y)
}

// Draw draws this object to the screen.
func (p *Parts) Draw(screen *ebiten.Image) {
	screen.DrawImage(p.image, p.op)
}
