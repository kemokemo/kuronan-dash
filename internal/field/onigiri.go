package field

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/assets/se"
	"github.com/kemokemo/kuronan-dash/internal/anime"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// Onigiri is a kind of Food. Delicious!
type Onigiri struct {
	anime   *anime.TimeAnimation
	op      *ebiten.DrawImageOptions
	rect    *view.HitRectangle
	stamina int
	tension int
	eaten   bool
	sound   *se.Player
}

// Initialize initializes the object.
//
//	args:
//	 img: the image to draw (ignore to use animation)
//	 pos: the initial position
//	 vel: the velocity to move this object
func (o *Onigiri) Initialize(img *ebiten.Image, pos *view.Vector, kv float64) {
	// ignore img

	o.anime = anime.NewTimeAnimation(images.OnigiriAnimation, 1.0)
	o.stamina = 5
	o.tension = 0
	o.eaten = false
	o.sound = se.PickupItem

	o.op = &ebiten.DrawImageOptions{}
	o.op.GeoM.Translate(pos.X, pos.Y+FieldOffset)

	firstFrame := images.OnigiriAnimation[0]
	w := firstFrame.Bounds().Dx()
	h := firstFrame.Bounds().Dy()
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
	screen.DrawImage(o.anime.GetCurrentFrame(), o.op)
}

// IsCollided returns whether this obstacle is collided the arg.
func (o *Onigiri) IsCollided(hr *view.HitRectangle) bool {
	// The food that is eaten is no longer subject to a hit decision.
	if o.eaten {
		return false
	}
	return o.rect.Overlaps(hr)
}

// Eat eats this food. This func returns the value to restore character's stamina and tension.
func (o *Onigiri) Eat(soundPlayFlag bool) (int, int) {
	o.eaten = true
	if soundPlayFlag {
		o.sound.Play()
	}
	return o.stamina, o.tension
}
