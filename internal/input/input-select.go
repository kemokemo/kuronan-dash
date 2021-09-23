package input

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type SelectInputChecker struct {
	RectArray     []image.Rectangle
	previousIndex int
	currentIndex  int
	mousePos      image.Point
	inArea        bool
	selected      bool
}

func (i *SelectInputChecker) Update() {
	i.inArea = false
	i.mousePos.X, i.mousePos.Y = ebiten.CursorPosition()
	for index := range i.RectArray {
		if !i.mousePos.In(i.RectArray[index]) {
			continue
		}
		i.previousIndex = i.currentIndex
		i.currentIndex = index
		i.inArea = true
		break
	}

	i.selected = false
	if i.inArea && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		i.selected = true
	}
}

func (i *SelectInputChecker) TriggeredUp() bool {
	return false
}

func (i *SelectInputChecker) TriggeredDown() bool {
	return false
}

func (i *SelectInputChecker) TriggeredLeft() bool {
	return i.previousIndex > i.currentIndex || inpututil.IsKeyJustReleased(ebiten.KeyLeft)
}

func (i *SelectInputChecker) TriggeredRight() bool {
	return i.previousIndex < i.currentIndex || inpututil.IsKeyJustReleased(ebiten.KeyRight)
}

func (i *SelectInputChecker) TriggeredStart() bool {
	return i.selected || inpututil.IsKeyJustReleased(ebiten.KeySpace)
}

func (i *SelectInputChecker) TriggeredPause() bool {
	return false
}

func (i *SelectInputChecker) TriggeredAttack() bool {
	return false
}

func (i *SelectInputChecker) TriggeredSpecial() bool {
	return false
}
