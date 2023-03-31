package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kemokemo/kuronan-dash/assets/images"
)

func NewCurtain() *Curtain {
	return &Curtain{
		img:              images.Curtain,
		op:               &ebiten.DrawImageOptions{},
		transitionFrames: 20,
	}
}

type Curtain struct {
	img              *ebiten.Image
	op               *ebiten.DrawImageOptions
	transitionFrames int
	counter          int
	changing         bool
	visibility       bool
	alpha            float32
}

// Start starts to change visible this image gradually.
// If arg is false, this will be hidden. If true, this will be appear.
func (c *Curtain) Start(visibility bool) {
	c.changing = true
	c.visibility = visibility
}

// Update updates visibility this image.
// Call after 'Start' function, until 'IsFinished' flag changes to true.
func (c *Curtain) Update() {
	if !c.changing {
		return
	}

	c.counter++
	if c.visibility {
		c.alpha = 0.05 * float32(c.counter)
	} else {
		c.alpha = 1.0 - (0.05 * float32(c.counter))
	}

	c.op.ColorScale.Reset()
	c.op.ColorScale.Scale(c.alpha, c.alpha, c.alpha, c.alpha)
}

func (c *Curtain) Draw(screen *ebiten.Image) {
	screen.DrawImage(c.img, c.op)
}

// IsFinished returns whether this image's transition of visibility finished.
func (c *Curtain) IsFinished() bool {
	if c.counter < c.transitionFrames {
		return false
	}

	c.counter = 0
	c.changing = false

	return true
}
