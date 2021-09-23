package input

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PlayerInputChecker struct {
	RectArray    []image.Rectangle
	currentIndex int
	mousePos     image.Point
	isUp         bool
	isDown       bool
}

func (i *PlayerInputChecker) Update() {
	i.isUp = false
	i.isDown = false

	i.mousePos.X, i.mousePos.Y = ebiten.CursorPosition()
	for index := range i.RectArray {
		if !i.mousePos.In(i.RectArray[index]) {
			continue
		}

		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && index != i.currentIndex {
			if i.currentIndex > index {
				i.isUp = true
				i.currentIndex--
			} else {
				i.isDown = true
				i.currentIndex++
			}
			return
		}
	}

}

func (i *PlayerInputChecker) TriggeredUp() bool {
	return i.isUp || inpututil.IsKeyJustReleased(ebiten.KeyUp)
}

func (i *PlayerInputChecker) TriggeredDown() bool {
	return i.isDown || inpututil.IsKeyJustReleased(ebiten.KeyDown)
}

func (i *PlayerInputChecker) TriggeredLeft() bool {
	return false
}

func (i *PlayerInputChecker) TriggeredRight() bool {
	return false
}

func (i *PlayerInputChecker) TriggeredStart() bool {
	return false
}

func (i *PlayerInputChecker) TriggeredPause() bool {
	return false
}

func (i *PlayerInputChecker) TriggeredAttack() bool {
	// todo
	return false
}

func (i *PlayerInputChecker) TriggeredSpecial() bool {
	// todo
	return false
}
