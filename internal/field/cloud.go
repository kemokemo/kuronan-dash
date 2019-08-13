package field

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// Cloud is the cloud object to draw.
type Cloud struct {
	Image *ebiten.Image

	pos   view.Position
	speed ScrollSpeed
	mag   float32
	view  view.Viewport
}

// Initialize initializes this cloud with a image and the initial position.
func (c *Cloud) Initialize(img *ebiten.Image, pos view.Position) {
	c.Image = img
	c.pos = pos
	c.view = view.Viewport{}
	c.view.SetSize(1280, 768)
}

// SetSpeed sets the speed to scroll.
func (c *Cloud) SetSpeed(speed ScrollSpeed) {
	c.speed = speed
}

// SetMagnification sets the magnification for the scroll speed.
func (c *Cloud) SetMagnification(mag float32) {
	c.mag = mag
}

// Update updates the position of this cloud according to the spped.
func (c *Cloud) Update() {
	switch c.speed {
	case Normal:
		c.view.SetVelocity(0.4 * c.mag)
	case Slow:
		c.view.SetVelocity(0.2 * c.mag)
	}

	c.view.Move(view.Left)
}

// Draw draws this cloud.
func (c *Cloud) Draw(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	x, y := c.view.Position()

	op.GeoM.Translate(float64(c.pos.X), float64(c.pos.Y))
	op.GeoM.Translate(float64(x), float64(y))

	err := screen.DrawImage(c.Image, op)
	if err != nil {
		return err
	}
	return nil
}
