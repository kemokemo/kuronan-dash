package ui

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// FrameWindow is a struct to draw a window with frame.
type FrameWindow struct {
	frameImg     *ebiten.Image
	innerImg     *ebiten.Image
	innerOp      *ebiten.DrawImageOptions
	frameDarkOp  *ebiten.DrawImageOptions
	frameLightOp *ebiten.DrawImageOptions
	rect         image.Rectangle
	counter      int
	enableBlink  bool
}

// NewFrameWindow returns a FrameWindow.
//
// The width and height are used for the inner region excluding the frame.
// If 0 is set to the frameWidth, the frame will not be drawn.
func NewFrameWindow(x, y, width, height, frameWidth int) *FrameWindow {
	fw := FrameWindow{
		rect: image.Rectangle{
			Min: image.Point{X: x, Y: y},
			Max: image.Point{X: x + width, Y: y + height},
		},
	}

	fw.innerImg = ebiten.NewImage(width, height)
	fw.innerImg.Fill(color.White)
	fw.innerOp = &ebiten.DrawImageOptions{}
	fw.innerOp.GeoM.Translate(float64(x), float64(y))

	if frameWidth > 0 {
		fw.frameImg = ebiten.NewImage(width+frameWidth*2, height+frameWidth*2)
		fw.frameImg.Fill(color.White)
		fw.frameDarkOp = &ebiten.DrawImageOptions{}
		fw.frameDarkOp.GeoM.Translate(float64(x-frameWidth), float64(y-frameWidth))
		fw.frameLightOp = &ebiten.DrawImageOptions{}
		fw.frameLightOp.GeoM.Translate(float64(x-frameWidth), float64(y-frameWidth))
	}
	return &fw
}

// GetWindowRect returns the rectangle of this window.
func (w *FrameWindow) GetWindowRect() image.Rectangle {
	return w.rect
}

// SetColors sets the colors of the window's inner region and the frame's
// normal color.
// If you need to blink the frame, please use the SetBlinkFrame method.
func (w *FrameWindow) SetColors(inner, frameDark, frameLight color.RGBA) {
	w.innerOp.ColorScale.Scale(colorScale(inner))
	if w.frameDarkOp != nil {
		w.frameDarkOp.ColorScale.Scale(colorScale(frameDark))
	}
	if w.frameLightOp != nil {
		w.frameLightOp.ColorScale.Scale(colorScale(frameLight))
	}
}

// SetBlink sets the flag to blink the frame.
func (w *FrameWindow) SetBlink(enableBlink bool) {
	w.enableBlink = enableBlink
}

// DrawWindow draws this window.
func (w *FrameWindow) DrawWindow(screen *ebiten.Image) {
	if w.frameImg != nil {
		screen.DrawImage(w.frameImg, w.getFrameOp())
	}
	screen.DrawImage(w.innerImg, w.innerOp)
}

func (w *FrameWindow) getFrameOp() *ebiten.DrawImageOptions {
	if !w.enableBlink {
		return w.frameDarkOp
	}

	w.counter++
	switch {
	case w.counter <= 30:
		return w.frameDarkOp
	case 30 < w.counter && w.counter <= 60:
		return w.frameLightOp
	case 60 < w.counter:
		w.counter = 0
		return w.frameDarkOp
	default:
		return w.frameDarkOp
	}
}

func colorScale(clr color.Color) (rf, gf, bf, af float32) {
	r, g, b, a := clr.RGBA()
	if a == 0 {
		return 0, 0, 0, 0
	}

	rf = float32(r) / float32(a)
	gf = float32(g) / float32(a)
	bf = float32(b) / float32(a)
	af = float32(a) / 0xffff
	return
}
