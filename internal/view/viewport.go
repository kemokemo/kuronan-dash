package view

import "math"

// Viewport is the view of the player's view.
type Viewport struct {
	// position
	x, y int

	// width ,height and velocity to move
	w, h int
	v    int

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
	p.v = int(math.Abs(v))
}

// SetLoop sets the loop or not flag.
// If true, you can loop this view.
func (p *Viewport) SetLoop(loop bool) {
	p.loop = loop
}

// Move moves the view of this Viewport.
func (p *Viewport) Move(d Direction) {
	x, y := getDirectionValue(d)
	p.x += x * p.w * p.v / 32
	p.y += y * p.h * p.v / 32
	if p.loop {
		p.x %= p.maxW
		p.y %= p.maxH
	}
}

// Position returns the position of this Viewport.
func (p *Viewport) Position() (int, int) {
	return p.x, p.y
}
