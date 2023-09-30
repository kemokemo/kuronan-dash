package field

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/assets/se"
	"github.com/kemokemo/kuronan-dash/internal/anime"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// IkariYaki is a kind of Food. Delicious!
type IkariYaki struct {
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
func (i *IkariYaki) Initialize(img *ebiten.Image, pos *view.Vector, kv float64) {
	// ignore img

	i.anime = anime.NewTimeAnimation(images.IkariYakiAnimation, 1.0)
	i.stamina = 50
	i.tension = 15
	i.eaten = false
	i.sound = se.PickupItem

	i.op = &ebiten.DrawImageOptions{}
	i.op.GeoM.Translate(pos.X, pos.Y+FieldOffset)

	firstFrame := images.IkariYakiAnimation[0]
	w := firstFrame.Bounds().Dx()
	h := firstFrame.Bounds().Dy()
	i.rect = view.NewHitRectangle(
		view.Vector{X: pos.X + rectOffset, Y: pos.Y + rectOffset},
		view.Vector{X: pos.X + float64(w) - rectOffset, Y: pos.Y + float64(h) - rectOffset})
}

// Update updates the position and velocity of this object.
//
//	args:
//	 scrollV: the velocity to scroll this field parts.
func (i *IkariYaki) Update(scrollV *view.Vector) {
	if i.eaten {
		return
	}
	i.op.GeoM.Translate(scrollV.X, scrollV.Y)
	i.rect.Add(scrollV)
}

// Draw draws this object to the screen.
func (i *IkariYaki) Draw(screen *ebiten.Image) {
	if i.eaten {
		return
	}
	screen.DrawImage(i.anime.GetCurrentFrame(), i.op)
}

// IsCollided returns whether this obstacle is collided the arg.
func (i *IkariYaki) IsCollided(hr *view.HitRectangle) bool {
	// The food that is eaten is no longer subject to a hit decision.
	if i.eaten {
		return false
	}
	return i.rect.Overlaps(hr)
}

// Eat eats this food. This func returns the value to restore character's stamina.
func (i *IkariYaki) Eat() (int, int) {
	i.eaten = true
	i.sound.Play()
	return i.stamina, i.tension
}
