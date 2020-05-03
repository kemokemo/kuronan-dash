package field

import (
	"github.com/kemokemo/kuronan-dash/internal/view"

	"github.com/hajimehoshi/ebiten"
)

const offset = 2

// Rock is the interface of the field part.
type Rock struct {
	image    *ebiten.Image
	position view.Vector
	rect     view.Rectangle
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
	r.rect = view.Rectangle{
		LeftBottom: view.Vector{X: pos.X, Y: pos.Y},
		RightTop:   view.Vector{X: pos.X + float64(w-offset), Y: pos.Y + float64(h-offset)},
	}
}

// Update updates the position and velocity of this object.
//  args:
//   charaV: the velocity of the player character
func (r *Rock) Update(charaV view.Vector) {
	r.position = r.position.Add(r.velocity)
	// Calculate relative speed with player only in horizontal direction
	r.position.X -= charaV.X
}

// Draw draws this object to the screen.
func (r *Rock) Draw(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(r.position.X, r.position.Y)

	return screen.DrawImage(r.image, op)
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

// IsColleded returns whether this obstacle is collided the arg.
func (r *Rock) IsCollided(rect view.Rectangle) bool {
	return r.rect.IsCollided(rect)
}
