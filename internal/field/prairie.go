package field

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// Prairie is the Prairie parts to move and draw.
type Prairie struct {
	Image *ebiten.Image

	pos   view.Position
	v0    float32
	speed ScrollSpeed
	view  view.Viewport
}

// Initialize initializes this cloud with a image and the initial position.
// The v0 argument is the initial scroll speed for the Normal speed.
func (p *Prairie) Initialize(img *ebiten.Image, pos view.Position, v0 float32) {
	p.Image = img
	p.pos = pos
	p.v0 = v0
	p.view = view.Viewport{}
	p.view.SetSize(1280, 768)
}

// SetSpeed sets the speed to scroll.
func (p *Prairie) SetSpeed(speed ScrollSpeed) {
	p.speed = speed
}

// Update updates the position of this cloud according to the spped.
func (p *Prairie) Update() {
	switch p.speed {
	case Normal:
		p.view.SetVelocity(p.v0)
	case Slow:
		p.view.SetVelocity(p.v0 * 0.5)
	}

	p.view.Move(view.Left)
}

// Draw draws this cloud.
func (p *Prairie) Draw(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	x16, y16 := p.view.Position()
	offsetX, offsetY := float64(x16)/16, float64(y16)/16

	op.GeoM.Translate(float64(p.pos.X), float64(p.pos.Y))
	op.GeoM.Translate(offsetX, offsetY)

	err := screen.DrawImage(p.Image, op)
	if err != nil {
		return err
	}
	return nil
}
