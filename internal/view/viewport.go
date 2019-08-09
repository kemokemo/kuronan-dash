package view

// Viewport is the view of the player's view.
type Viewport struct {
	x16    int
	y16    int
	x      int
	y      int
	maxX16 int
	maxY16 int
}

// SetSize sets the size of this Viewport.
func (p *Viewport) SetSize(w, h int) {
	p.x = w
	p.y = h
	p.maxX16 = p.x * 16
	p.maxY16 = p.y * 16
}

// Move moves the view of this Viewport.
func (p *Viewport) Move(d Direction) {
	x, y := getDirectionValue(d)
	p.x16 += x * p.x / 32
	p.y16 += y * p.y / 32
	p.x16 %= p.maxX16
	p.y16 %= p.maxY16
}

// Position returns the position of this Viewport.
func (p *Viewport) Position() (int, int) {
	return p.x16, p.y16
}
