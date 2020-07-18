package field

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// Onigiri is a kind of Food. Oishii yo!
type Onigiri struct {
	image    *ebiten.Image
	imgSize  image.Point
	position view.Vector
	rect     image.Rectangle
	stamina  int
	eaten    bool
}

// Initialize initializes the object.
//  args:
//   img: the image to draw
//   pos: the initial position
//   vel: the velocity to move this object
func (o *Onigiri) Initialize(img *ebiten.Image, pos view.Vector, vel view.Vector) {
	o.image = img
	o.position = pos
	o.stamina = 20
	o.eaten = false

	w, h := img.Size()
	o.imgSize = image.Point{w, h}
	o.rect = image.Rectangle{
		Min: image.Point{X: int(pos.X), Y: int(pos.Y)},
		Max: image.Point{X: int(pos.X) + w - offset, Y: int(pos.Y) + h - offset},
	}
}

// Update updates the position and velocity of this object.
//  args:
//   charaV: the velocity of the player character
func (o *Onigiri) Update(charaV view.Vector) {
	if o.eaten {
		return
	}
	o.position.X -= charaV.X

	o.rect.Min = image.Point{X: int(o.position.X), Y: int(o.position.Y)}
	o.rect.Max = image.Point{X: int(o.position.X) + o.imgSize.X - offset, Y: int(o.position.Y) + o.imgSize.Y - offset}
}

// Draw draws this object to the screen.
func (o *Onigiri) Draw(screen *ebiten.Image, offset image.Point) error {
	if o.eaten {
		return nil
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(o.position.X-float64(offset.X), o.position.Y-float64(offset.Y))
	return screen.DrawImage(o.image, op)
}

// IsCollided returns whether this obstacle is collided the arg.
func (o *Onigiri) IsCollided(r image.Rectangle) bool {
	// The food that is eaten is no longer subject to a hit decision.
	if o.eaten {
		return false
	}
	return o.rect.Overlaps(r)
}

// Eat eats this food. This func reteruns the value to restore character's stamina.
func (o *Onigiri) Eat() int {
	o.eaten = true
	return o.stamina
}
