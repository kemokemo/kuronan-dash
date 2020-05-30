package view

// Viewport is the view of the player's view.
type Viewport struct {
	// position
	x, y int

	// width ,height and velocity to move
	w, h int
	v    float64

	// loop settings
	loop       bool
	maxW, maxH int
}

// SetSize sets the size of this Viewport.
func (p *Viewport) SetSize(w, h int) {
	p.w = w
	p.h = h
	p.maxW = w * 16
	p.maxH = h * 16
}

// SetVelocity sets the velocity of this viewport.
func (p *Viewport) SetVelocity(v float64) {
	p.v = v
}

// SetLoop sets the loop or not flag.
// If true, you can loop this view.
func (p *Viewport) SetLoop(loop bool) {
	p.loop = loop
}

// Move moves the view of this Viewport.
func (p *Viewport) Move(d Direction) {
	x, y := getDirectionValue(d)
	p.x += int(float64(x) * float64(p.w) * p.v / 32.0)
	p.y += int(float64(y) * float64(p.h) * p.v / 32.0)
	if p.loop {
		p.x %= p.maxW
		p.y %= p.maxH
	}
}

// Position returns the position of this Viewport.
func (p *Viewport) Position() (int, int) {
	return p.x, p.y
}
