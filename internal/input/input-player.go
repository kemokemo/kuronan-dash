package input

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	vpad "github.com/kemokemo/ebiten-virtualpad"
)

type PlayerInputChecker struct {
	RectArray    []image.Rectangle
	UpBtn        vpad.TriggerButton
	DownBtn      vpad.TriggerButton
	currentIndex int
	mousePos     image.Point
	isUp         bool
	isDown       bool
}

func (i *PlayerInputChecker) Update() {
	i.isUp = false
	i.isDown = false

	// Mouse
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

	// Keyboard
	if inpututil.IsKeyJustReleased(ebiten.KeyUp) {
		i.isUp = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyDown) {
		i.isDown = true
	}

	// Button
	i.UpBtn.Update()
	if i.UpBtn.IsTriggered() {
		i.isUp = true
	}
	i.DownBtn.Update()
	if i.DownBtn.IsTriggered() {
		i.isDown = true
	}
}

func (i *PlayerInputChecker) TriggeredUp() bool {
	return i.isUp
}

func (i *PlayerInputChecker) TriggeredDown() bool {
	return i.isDown
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
