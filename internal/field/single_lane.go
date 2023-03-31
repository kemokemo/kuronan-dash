package field

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// SingleLane is the scrollable object to implement the lanes.
// Mainly for seamlessly updating and drawing locations, and drawing individual lanes.
type SingleLane struct {
	image *ebiten.Image
	op    *ebiten.DrawImageOptions
	pos   *view.Vector
	repos *view.Vector
	width float64
}

// Initialize initializes the object.
//  args:
//   img: the image to draw
//   pos: the initial position
//   vel: the velocity to move this object
func (l *SingleLane) Initialize(img *ebiten.Image, pos *view.Vector, vel *view.Vector) {
	l.image = img
	l.op = &ebiten.DrawImageOptions{}
	l.op.GeoM.Translate(pos.X, pos.Y)
	l.pos = &view.Vector{X: pos.X, Y: pos.Y}

	w := img.Bounds().Dx()
	l.width = float64(w)
	l.repos = &view.Vector{X: 2.0 * l.width, Y: 0.0}
}

// Update updates the position and velocity of this object.
//  args:
//   scrollV: the velocity to scroll field parts.
func (l *SingleLane) Update(scrollV *view.Vector) {
	l.pos.Add(scrollV)
	l.op.GeoM.Translate(scrollV.X, scrollV.Y)

	// If this lane moves off the screen, reposition it so that it becomes a candidate for drawing again.
	if l.pos.X <= -l.width {
		l.pos.Add(l.repos)
		l.op.GeoM.Translate(l.repos.X, l.repos.Y)
	}
}

// Draw draws this object to the screen.
func (l *SingleLane) Draw(screen *ebiten.Image) {
	screen.DrawImage(l.image, l.op)
}
