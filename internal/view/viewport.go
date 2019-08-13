package view

// Viewport is the view of the player's view.
type Viewport struct {
	x16    int
	y16    int
	x      int
	y      int
	maxX16 int
	maxY16 int
	v      float32
}

// SetSize sets the size of this Viewport.
func (p *Viewport) SetSize(w, h int) {
	p.x = w
	p.y = h
	p.maxX16 = p.x * 16
	p.maxY16 = p.y * 16
}

// SetVelocity sets the velocity of this viewport.
func (p *Viewport) SetVelocity(v float32) {
	p.v = v
}

// Move moves the view of this Viewport.
func (p *Viewport) Move(d Direction) {
	x, y := getDirectionValue(d)
	p.x16 += int(float32(x) * float32(p.x) * p.v / 32.0)
	p.y16 += int(float32(y) * float32(p.y) * p.v / 32.0)
	p.x16 %= p.maxX16
	p.y16 %= p.maxY16
}

// Position returns the position of this Viewport.
func (p *Viewport) Position() (int, int) {
	return p.x16, p.y16
}
