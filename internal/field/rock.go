package field

import (
	"github.com/kemokemo/kuronan-dash/assets/se"
	"github.com/kemokemo/kuronan-dash/internal/view"

	"github.com/hajimehoshi/ebiten/v2"
)

// Rock is the one of obstacles.
type Rock struct {
	image    *ebiten.Image
	op       *ebiten.DrawImageOptions
	rect     *view.HitRectangle
	hardness float64
	broken   bool
	sound    *se.Player
}

// Initialize initializes the object.
//
//	args:
//	 img: the image to draw
//	 pos: the initial position
//	 vel: the velocity to move this object
func (r *Rock) Initialize(img *ebiten.Image, pos *view.Vector, kv float64) {
	r.image = img
	r.hardness = 2
	r.sound = se.BreakRock

	r.op = &ebiten.DrawImageOptions{}
	r.op.GeoM.Translate(pos.X, pos.Y+FieldOffset)

	w := img.Bounds().Dx()
	h := img.Bounds().Dy()
	r.rect = view.NewHitRectangle(
		view.Vector{X: pos.X + rectOffset, Y: pos.Y + rectOffset},
		view.Vector{X: pos.X + float64(w) - rectOffset, Y: pos.Y + float64(h) - rectOffset})
}

// Update updates the position and velocity of this object.
//
//	args:
//	 scrollV: the velocity to scroll this field parts.
func (r *Rock) Update(scrollV *view.Vector) {
	if r.broken {
		return
	}
	r.op.GeoM.Translate(scrollV.X, scrollV.Y)
	r.rect.Add(scrollV)
}

// Draw draws this object to the screen.
func (r *Rock) Draw(screen *ebiten.Image) {
	if r.broken {
		return
	}
	screen.DrawImage(r.image, r.op)
}

// SetHardness sets the hardness of this obstacle.
func (r *Rock) SetHardness(hardness float64) {
	r.hardness = hardness
}

// Attack attacks this obstacle.
// The damage value reduces this obstacle's hardness.
func (r *Rock) Attack(damage float64, soundPlayFlag bool) {
	r.hardness -= damage
	if r.hardness <= 0 {
		r.broken = true
		if soundPlayFlag {
			r.sound.Play()
		}
	}
}

// IsBroken returns whether this obstacle was broken.
// The broken state means that the hardness is 0 or less.
func (r *Rock) IsBroken() bool {
	return r.broken
}

// IsCollided returns whether this obstacle is collided the arg.
func (r *Rock) IsCollided(rect *view.HitRectangle) bool {
	if r.broken {
		return false
	}
	return r.rect.Overlaps(rect)
}
