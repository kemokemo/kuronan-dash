package ui

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
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
func NewFrameWindow(x, y, width, height, frameWidth int) (*FrameWindow, error) {
	fw := FrameWindow{
		rect: image.Rectangle{
			Min: image.Point{X: x, Y: y},
			Max: image.Point{X: x + width, Y: y + height},
		},
	}
	var err error
	fw.innerImg, err = ebiten.NewImage(width, height, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	err = fw.innerImg.Fill(color.White)
	if err != nil {
		return nil, err
	}
	fw.innerOp = &ebiten.DrawImageOptions{}
	fw.innerOp.GeoM.Translate(float64(x), float64(y))

	if frameWidth > 0 {
		fw.frameImg, err = ebiten.NewImage(width+frameWidth*2, height+frameWidth*2, ebiten.FilterDefault)
		if err != nil {
			return nil, err
		}
		err = fw.frameImg.Fill(color.White)
		if err != nil {
			return nil, err
		}
		fw.frameDarkOp = &ebiten.DrawImageOptions{}
		fw.frameDarkOp.GeoM.Translate(float64(x-frameWidth), float64(y-frameWidth))
		fw.frameLightOp = &ebiten.DrawImageOptions{}
		fw.frameLightOp.GeoM.Translate(float64(x-frameWidth), float64(y-frameWidth))
	}
	return &fw, nil
}

// GetWindowRect returns the rectangle of this window.
func (w *FrameWindow) GetWindowRect() image.Rectangle {
	return w.rect
}

// SetColors sets the colors of the window's inner region and the frame's
// normal color.
// If you need to blink the frame, please use the SetBlinkFrame method.
func (w *FrameWindow) SetColors(inner, frameDark, frameLight color.RGBA) {
	w.innerOp.ColorM.Scale(colorScale(inner))
	if w.frameDarkOp != nil {
		w.frameDarkOp.ColorM.Scale(colorScale(frameDark))
	}
	if w.frameLightOp != nil {
		w.frameLightOp.ColorM.Scale(colorScale(frameLight))
	}
}

// SetBlink sets the flag to blink the frame.
func (w *FrameWindow) SetBlink(enableBlink bool) {
	w.enableBlink = enableBlink
}

// DrawWindow draws this window.
func (w *FrameWindow) DrawWindow(screen *ebiten.Image) {
	var err error
	if w.frameImg != nil {
		err = screen.DrawImage(w.frameImg, w.getFrameOp())
		if err != nil {
			log.Println("failed to draw the frame image", err)
		}
	}
	err = screen.DrawImage(w.innerImg, w.innerOp)
	if err != nil {
		log.Println("failed to draw the inner image", err)
	}
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

func colorScale(clr color.Color) (rf, gf, bf, af float64) {
	r, g, b, a := clr.RGBA()
	if a == 0 {
		return 0, 0, 0, 0
	}

	rf = float64(r) / float64(a)
	gf = float64(g) / float64(a)
	bf = float64(b) / float64(a)
	af = float64(a) / 0xffff
	return
}
