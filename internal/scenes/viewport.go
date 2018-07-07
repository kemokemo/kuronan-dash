package scenes

import "image"

type viewport struct {
	x16    int
	y16    int
	bgSize image.Point
}

func (p *viewport) SetSize(w, h int) {
	p.bgSize = image.Point{X: w, Y: h}
}

func (p *viewport) Move() {
	maxX16 := p.bgSize.X * 16
	maxY16 := p.bgSize.Y * 16

	p.x16 += p.bgSize.X / 32
	p.y16 += p.bgSize.Y / 32
	p.x16 %= maxX16
	p.y16 %= maxY16
}

func (p *viewport) Position() (int, int) {
	return p.x16, p.y16
}
