package field

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kemokemo/kuronan-dash/assets/se"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// Onigiri is a kind of Food. Delicious!
type Onigiri struct {
	image   *ebiten.Image
	op      *ebiten.DrawImageOptions
	rect    *view.HitRectangle
	stamina int
	eaten   bool
	sound   *se.Player
}

// Initialize initializes the object.
//
//	args:
//	 img: the image to draw
//	 pos: the initial position
//	 vel: the velocity to move this object
func (o *Onigiri) Initialize(img *ebiten.Image, pos *view.Vector, kv float64) {
	o.image = img
	o.stamina = 5
	o.eaten = false
	o.sound = se.PickupItem

	o.op = &ebiten.DrawImageOptions{}
	o.op.GeoM.Translate(pos.X, pos.Y+FieldOffset)

	w := img.Bounds().Dx()
	h := img.Bounds().Dy()
	o.rect = view.NewHitRectangle(
		view.Vector{X: pos.X + rectOffset, Y: pos.Y + rectOffset},
		view.Vector{X: pos.X + float64(w) - rectOffset, Y: pos.Y + float64(h) - rectOffset})
}

// Update updates the position and velocity of this object.
//
//	args:
//	 scrollV: the velocity to scroll this field parts.
func (o *Onigiri) Update(scrollV *view.Vector) {
	if o.eaten {
		return
	}
	o.op.GeoM.Translate(scrollV.X, scrollV.Y)
	o.rect.Add(scrollV)
}

// Draw draws this object to the screen.
func (o *Onigiri) Draw(screen *ebiten.Image) {
	if o.eaten {
		return
	}
	screen.DrawImage(o.image, o.op)
}

// IsCollided returns whether this obstacle is collided the arg.
func (o *Onigiri) IsCollided(hr *view.HitRectangle) bool {
	// The food that is eaten is no longer subject to a hit decision.
	if o.eaten {
		return false
	}
	return o.rect.Overlaps(hr)
}

// Eat eats this food. This func returns the value to restore character's stamina.
func (o *Onigiri) Eat() int {
	o.eaten = true
	o.sound.Play()
	return o.stamina
}
