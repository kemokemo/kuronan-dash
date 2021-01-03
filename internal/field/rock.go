package field

import (
	"github.com/kemokemo/kuronan-dash/internal/view"

	"github.com/hajimehoshi/ebiten/v2"
)

// Rock is the one of obstacles.
type Rock struct {
	image    *ebiten.Image
	op       *ebiten.DrawImageOptions
	rect     *view.HitRectangle
	velocity view.Vector
	hardness float64
}

// Initialize initializes the object.
//  args:
//   img: the image to draw
//   pos: the initial position
//   vel: the velocity to move this object
func (r *Rock) Initialize(img *ebiten.Image, pos, vel view.Vector) {
	r.image = img
	r.velocity = vel

	r.op = &ebiten.DrawImageOptions{}
	r.op.GeoM.Translate(pos.X, pos.Y-float64(fieldOffset))

	w, h := img.Size()
	r.rect = view.NewHitRectangle(
		view.Vector{X: pos.X + rectOffset, Y: pos.Y + rectOffset},
		view.Vector{X: pos.X + float64(w) - rectOffset, Y: pos.Y + float64(h) - rectOffset})
}

// Update updates the position and velocity of this object.
//  args:
//   charaV: the velocity of the player character
func (r *Rock) Update(charaV view.Vector) {
	// Calculate relative speed with player only in horizontal direction
	vX := r.velocity.X - charaV.X
	r.op.GeoM.Translate(vX, r.velocity.Y)
	r.rect.Add(view.Vector{X: vX, Y: r.velocity.Y})
}

// Draw draws this object to the screen.
func (r *Rock) Draw(screen *ebiten.Image) {
	screen.DrawImage(r.image, r.op)
}

// SetHardness sets the hardness of this obstacle.
func (r *Rock) SetHardness(hardness float64) {
	r.hardness = hardness
}

// Attack attacks this obstacle.
// The damage value reduces this obstacle's hardness.
func (r *Rock) Attack(damage float64) {
	r.hardness -= damage
}

// IsBroken returns whether this obstacle was broken.
// The broken state means that the hardness is 0 or less.
func (r *Rock) IsBroken() bool {
	return r.hardness <= 0
}

// IsCollided returns whether this obstacle is collided the arg.
func (r *Rock) IsCollided(rect *view.HitRectangle) bool {
	return r.rect.Overlaps(rect)
}
