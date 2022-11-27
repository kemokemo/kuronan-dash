package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	vpad "github.com/kemokemo/ebiten-virtualpad"
)

type SelectInputChecker struct {
	GoBtn vpad.TriggerButton
}

func (i *SelectInputChecker) Update() {
	i.GoBtn.Update()
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
	return i.GoBtn.IsTriggered()
}

func (i *SelectInputChecker) TriggeredPause() bool {
	return false
}

func (i *SelectInputChecker) TriggeredAttack() bool {
	return false
}

func (i *SelectInputChecker) TriggeredSkill() bool {
	return false
}
