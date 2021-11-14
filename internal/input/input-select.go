package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type SelectInputChecker struct {
}

func (i *SelectInputChecker) Update() {
}

func (i *SelectInputChecker) TriggeredUp() bool {
	return false
}

func (i *SelectInputChecker) TriggeredDown() bool {
	return false
}

func (i *SelectInputChecker) TriggeredLeft() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeyLeft)
}

func (i *SelectInputChecker) TriggeredRight() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeyRight)
}

func (i *SelectInputChecker) TriggeredStart() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeySpace)
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
