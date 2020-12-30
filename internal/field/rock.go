package field

import (
	"image"

	"github.com/kemokemo/kuronan-dash/internal/view"

	"github.com/hajimehoshi/ebiten/v2"
)

// Position offset to make it look like it's on the lane
const offset = 2

// Rock is the one of obstacles.
type Rock struct {
	image    *ebiten.Image
	imgSize  image.Point
	position view.Vector
	rect     image.Rectangle
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
	r.position = pos
	r.velocity = vel

	w, h := img.Size()
	r.imgSize = image.Point{w, h}
	r.rect = image.Rectangle{
		Min: image.Point{X: int(pos.X), Y: int(pos.Y)},
		Max: image.Point{X: int(pos.X) + w - offset, Y: int(pos.Y) + h - offset},
	}
}

// Update updates the position and velocity of this object.
//  args:
//   charaV: the velocity of the player character
func (r *Rock) Update(charaV view.Vector) {
	r.position = r.position.Add(r.velocity)
	// Calculate relative speed with player only in horizontal direction
	r.position.X -= charaV.X

	r.rect.Min = image.Point{X: int(r.position.X), Y: int(r.position.Y)}
	r.rect.Max = image.Point{X: int(r.position.X) + r.imgSize.X - offset, Y: int(r.position.Y) + r.imgSize.Y - offset}
}

// Draw draws this object to the screen.
func (r *Rock) Draw(screen *ebiten.Image, offset image.Point) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(r.position.X-float64(offset.X), r.position.Y-float64(offset.Y))
	screen.DrawImage(r.image, op)
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
func (r *Rock) IsCollided(rect image.Rectangle) bool {
	return r.rect.Overlaps(rect)
}
