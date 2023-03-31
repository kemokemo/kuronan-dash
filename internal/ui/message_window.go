package ui

import (
	"image"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/kemokemo/kuronan-dash/assets/fonts"
	"golang.org/x/image/font"
)

// MessageWindow is a struct to draw a window with frame.
type MessageWindow struct {
	frameImg     *ebiten.Image
	innerImg     *ebiten.Image
	innerOp      *ebiten.DrawImageOptions
	frameDarkOp  *ebiten.DrawImageOptions
	frameLightOp *ebiten.DrawImageOptions
	rect         image.Rectangle
	counter      int
	enableBlink  bool
	font         font.Face
	fontSize     int
	lineSpacing  int
}

// NewMessageWindow returns a MessageWindow.
//
// The width and height are used for the inner region excluding the frame.
// If 0 is set to the frameWidth, the frame will not be drawn.
func NewMessageWindow(x, y, width, height, frameWidth int) *MessageWindow {
	mw := MessageWindow{
		rect: image.Rectangle{
			Min: image.Point{X: x, Y: y},
			Max: image.Point{X: x + width, Y: y + height},
		},
	}

	mw.innerImg = ebiten.NewImage(width, height)
	mw.innerImg.Fill(color.White)
	mw.innerOp = &ebiten.DrawImageOptions{}
	mw.innerOp.GeoM.Translate(float64(x), float64(y))

	mw.frameImg = ebiten.NewImage(width+frameWidth*2, height+frameWidth*2)
	mw.frameImg.Fill(color.White)
	mw.frameDarkOp = &ebiten.DrawImageOptions{}
	mw.frameDarkOp.GeoM.Translate(float64(x-frameWidth), float64(y-frameWidth))
	mw.frameLightOp = &ebiten.DrawImageOptions{}
	mw.frameLightOp.GeoM.Translate(float64(x-frameWidth), float64(y-frameWidth))
	mw.font = fonts.GamerFontM
	mw.fontSize = 20
	mw.lineSpacing = 2

	return &mw
}

func (mw *MessageWindow) SetSizeAndSpacing(fontSize, lineSpacing int) {
	mw.fontSize = fontSize
	mw.lineSpacing = lineSpacing
}

// GetWindowRect returns the rectangle of this window.
func (mw *MessageWindow) GetWindowRect() image.Rectangle {
	return mw.rect
}

// SetColors sets the colors of the window's inner region and the frame's
// normal color.
// If you need to blink the frame, please use the SetBlinkFrame method.
func (mw *MessageWindow) SetColors(inner, frameDark, frameLight color.RGBA) {
	mw.innerOp.ColorScale.Scale(colorScale(inner))
	if mw.frameDarkOp != nil {
		mw.frameDarkOp.ColorScale.Scale(colorScale(frameDark))
	}
	if mw.frameLightOp != nil {
		mw.frameLightOp.ColorScale.Scale(colorScale(frameLight))
	}
}

// SetBlink sets the flag to blink the frame.
func (mw *MessageWindow) SetBlink(enableBlink bool) {
	mw.enableBlink = enableBlink
}

// DrawWindow draws this window.
func (mw *MessageWindow) DrawWindow(screen *ebiten.Image, msg string) {
	if mw.frameImg != nil {
		screen.DrawImage(mw.frameImg, mw.getFrameOp())
	}
	screen.DrawImage(mw.innerImg, mw.innerOp)

	mw.drawMessage(screen, msg)
}

func (mw *MessageWindow) drawMessage(screen *ebiten.Image, msg string) {
	rowRunesNum := (mw.rect.Max.X - mw.rect.Min.X) / 15

	startPoint := mw.takeTextPosition()
	lineNum := 1
	x := startPoint.X
	y := startPoint.Y
	rowMsg := ""

	rows := strings.Split(msg, "\n")
	for _, row := range rows {
		runes := []rune(row)

		for i := 0; i < len(runes); i += (rowRunesNum - 1) {
			y = y + mw.fontSize + mw.lineSpacing
			if i+rowRunesNum < len(runes) {
				rowMsg = string(runes[i : i+rowRunesNum-1])
			} else {
				rowMsg = string(runes[i:])
			}
			text.Draw(screen, rowMsg, mw.font, x, y, color.White)

			lineNum++
		}
	}
}

func (mw *MessageWindow) takeTextPosition() image.Point {
	x := mw.rect.Min.X + 15
	y := mw.rect.Min.Y + 7
	return image.Point{X: x, Y: y}
}

func (mw *MessageWindow) getFrameOp() *ebiten.DrawImageOptions {
	if !mw.enableBlink {
		return mw.frameDarkOp
	}

	mw.counter++
	switch {
	case mw.counter <= 30:
		return mw.frameDarkOp
	case 30 < mw.counter && mw.counter <= 60:
		return mw.frameLightOp
	case 60 < mw.counter:
		mw.counter = 0
		return mw.frameDarkOp
	default:
		return mw.frameDarkOp
	}
}
