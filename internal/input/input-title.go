package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type TitleInputChecker struct {
}

// memo:
// These implementations are generated by below.
// impl 'i *TitleInputChecker' InputChecker >> title-input.go

func (i *TitleInputChecker) Update() {}

func (i *TitleInputChecker) TriggeredUp() bool {
	return false
}

func (i *TitleInputChecker) TriggeredDown() bool {
	return false
}

func (i *TitleInputChecker) TriggeredLeft() bool {
	return false
}

func (i *TitleInputChecker) TriggeredRight() bool {
	return false
}

func (i *TitleInputChecker) TriggeredStart() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeySpace) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
}

func (i *TitleInputChecker) TriggeredPause() bool {
	return false
}

func (i *TitleInputChecker) TriggeredAttack() bool {
	return false
}

func (i *TitleInputChecker) TriggeredSpecial() bool {
	return false
}
