package ui

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// FrameWindow is a struct to draw a window with frame.
type FrameWindow struct {
	rect            image.Rectangle
	width           int
	heigt           int
	innerColor      color.RGBA
	frameWidth      int
	frameDarkColor  color.RGBA
	frameLightColor color.RGBA
	counter         int
	enableBlink     bool
}

// NewFrameWindow returns a FrameWindow.
//
// The width and height are used for the inner region excluding the frame.
// If 0 is set to the frameWidth, the frame will not be drawn.
func NewFrameWindow(x, y, width, heigt, frameWidth int) *FrameWindow {
	return &FrameWindow{
		rect: image.Rectangle{
			Min: image.Point{X: x, Y: y},
			Max: image.Point{X: x + width, Y: y + heigt},
		},
		width: width, heigt: heigt,
		frameWidth: frameWidth,
	}
}

// GetWindowRect returns the rectangle of this window.
func (w *FrameWindow) GetWindowRect() image.Rectangle {
	return w.rect
}

// SetColors sets the colors of the window's inner region and the frame's
// normal color.
// If you need to blink the frame, please use the SetBlinkFrame method.
func (w *FrameWindow) SetColors(inner, frameDark, frameLight color.RGBA) {
	w.innerColor = inner
	w.frameDarkColor = frameDark
	w.frameLightColor = frameLight
}

// SetBlink sets the flag to blink the frame.
func (w *FrameWindow) SetBlink(enableBlink bool) {
	w.enableBlink = enableBlink
}

// DrawWindow draws this window.
func (w *FrameWindow) DrawWindow(screen *ebiten.Image) {
	if w.frameWidth > 0 {
		ebitenutil.DrawRect(screen,
			float64(w.rect.Min.X-w.frameWidth), float64(w.rect.Min.Y-w.frameWidth),
			float64(w.width+w.frameWidth*2), float64(w.heigt+w.frameWidth*2),
			w.getFrameColor())
	}
	ebitenutil.DrawRect(screen,
		float64(w.rect.Min.X), float64(w.rect.Min.Y),
		float64(w.width), float64(w.heigt),
		w.innerColor)
}

func (w *FrameWindow) getFrameColor() color.RGBA {
	if !w.enableBlink {
		return w.frameDarkColor
	}

	w.counter++
	switch {
	case w.counter <= 30:
		return w.frameDarkColor
	case 30 < w.counter && w.counter <= 60:
		return w.frameLightColor
	case 60 < w.counter:
		w.counter = 0
		return w.frameDarkColor
	default:
		return w.frameDarkColor
	}
}
