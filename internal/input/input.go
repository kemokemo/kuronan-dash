package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// InputChecker is the checker for user input.
type InputChecker interface {
	Update()
	TriggeredUp() bool
	TriggeredDown() bool
	TriggeredLeft() bool
	TriggeredRight() bool
	TriggeredStart() bool
	TriggeredPause() bool
	TriggeredAttack() bool
	TriggeredSpecial() bool
}

func TriggeredOne() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeySpace)
}

func TriggeredUp() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeyArrowUp)
}

func TriggeredDown() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeyArrowDown)
}

func TriggeredLeft() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeyArrowLeft)
}

func TriggeredRight() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeyArrowRight)
}
