package field

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// Mountain is the mountain parts to move and draw.
type Mountain struct {
	Image *ebiten.Image

	pos   view.Position
	v0    float32
	speed ScrollSpeed
	view  view.Viewport
}

// Initialize initializes this cloud with a image and the initial position.
// The v0 argument is the initial scroll speed for the Normal speed.
func (m *Mountain) Initialize(img *ebiten.Image, pos view.Position, v0 float32) {
	m.Image = img
	m.pos = pos
	m.v0 = v0
	m.view = view.Viewport{}
	m.view.SetSize(1280, 768)
}

// SetSpeed sets the speed to scroll.
func (m *Mountain) SetSpeed(speed ScrollSpeed) {
	m.speed = speed
}

// Update updates the position of this cloud according to the spped.
func (m *Mountain) Update() {
	switch m.speed {
	case Normal:
		m.view.SetVelocity(m.v0)
	case Slow:
		m.view.SetVelocity(m.v0 * 0.5)
	}

	m.view.Move(view.Left)
}

// Draw draws this cloud.
func (m *Mountain) Draw(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	x16, y16 := m.view.Position()
	offsetX, offsetY := float64(x16)/16, float64(y16)/16

	op.GeoM.Translate(float64(m.pos.X), float64(m.pos.Y))
	op.GeoM.Translate(offsetX, offsetY)

	err := screen.DrawImage(m.Image, op)
	if err != nil {
		return err
	}
	return nil
}
