package field

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/assets/se"
	"github.com/kemokemo/kuronan-dash/internal/anime"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// YakiManjuu is a kind of Food. Delicious!
type YakiManjuu struct {
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
func (y *YakiManjuu) Initialize(img *ebiten.Image, pos *view.Vector, kv float64) {
	// ignore img

	y.anime = anime.NewTimeAnimation(images.YakiManjuuAnimation, 1.0)
	y.stamina = 20
	y.tension = 5
	y.eaten = false
	y.sound = se.PickupItem

	y.op = &ebiten.DrawImageOptions{}
	y.op.GeoM.Translate(pos.X, pos.Y+FieldOffset)

	firstFrame := images.YakiManjuuAnimation[0]
	w := firstFrame.Bounds().Dx()
	h := firstFrame.Bounds().Dy()
	y.rect = view.NewHitRectangle(
		view.Vector{X: pos.X + rectOffset, Y: pos.Y + rectOffset},
		view.Vector{X: pos.X + float64(w) - rectOffset, Y: pos.Y + float64(h) - rectOffset})
}

// Update updates the position and velocity of this object.
//
//	args:
//	 scrollV: the velocity to scroll this field parts.
func (y *YakiManjuu) Update(scrollV *view.Vector) {
	if y.eaten {
		return
	}
	y.op.GeoM.Translate(scrollV.X, scrollV.Y)
	y.rect.Add(scrollV)
}

// Draw draws this object to the screen.
func (y *YakiManjuu) Draw(screen *ebiten.Image) {
	if y.eaten {
		return
	}
	screen.DrawImage(y.anime.GetCurrentFrame(), y.op)
}

// IsCollided returns whether this obstacle is collided the arg.
func (y *YakiManjuu) IsCollided(hr *view.HitRectangle) bool {
	// The food that is eaten is no longer subject to a hit decision.
	if y.eaten {
		return false
	}
	return y.rect.Overlaps(hr)
}

// Eat eats this food. This func returns the value to restore character's stamina and tension.
func (y *YakiManjuu) Eat(soundPlayFlag bool) (int, int) {
	y.eaten = true
	if soundPlayFlag {
		y.sound.Play()
	}
	return y.stamina, y.tension
}
