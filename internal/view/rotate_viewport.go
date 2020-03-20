package view

// RotateViewport is the view of the player's view.
type RotateViewport struct {
	x16    int
	y16    int
	x      int
	y      int
	maxX16 int
	maxY16 int
	v      float64
}

// SetSize sets the size of this RotateViewport.
func (p *RotateViewport) SetSize(w, h int) {
	p.x = w
	p.y = h
	p.maxX16 = p.x * 16
	p.maxY16 = p.y * 16
}

// SetVelocity sets the velocity of this RotateViewport.
func (p *RotateViewport) SetVelocity(v float64) {
	p.v = v
}

// Move moves the view of this RotateViewport.
func (p *RotateViewport) Move(d Direction) {
	x, y := getDirectionValue(d)
	p.x16 += int(float64(x) * float64(p.x) * p.v / 32.0)
	p.y16 += int(float64(y) * float64(p.y) * p.v / 32.0)
	p.x16 %= p.maxX16
	p.y16 %= p.maxY16
}

// Position returns the position of this RotateViewport.
func (p *RotateViewport) Position() (int, int) {
	return p.x16, p.y16
}
