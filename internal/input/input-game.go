package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	vpad "github.com/kemokemo/ebiten-virtualpad"
)

type GameInputChecker struct {
	StartBtn  vpad.TriggerButton
	PauseBtn  vpad.TriggerButton
	UpBtn     vpad.TriggerButton
	DownBtn   vpad.TriggerButton
	AttackBtn vpad.TriggerButton
	SkillBtn  vpad.TriggerButton
}

func (gi *GameInputChecker) Update() {
	gi.StartBtn.Update()
	gi.PauseBtn.Update()
}

func (gi *GameInputChecker) TriggeredUp() bool {
	return gi.UpBtn.IsTriggered()
}

func (gi *GameInputChecker) TriggeredDown() bool {
	return gi.DownBtn.IsTriggered()
}

func (gi *GameInputChecker) TriggeredLeft() bool {
	return false
}

func (gi *GameInputChecker) TriggeredRight() bool {
	return false
}

func (gi *GameInputChecker) TriggeredStart() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeySpace) || gi.StartBtn.IsTriggered()
}

func (gi *GameInputChecker) TriggeredPause() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeySpace) || gi.PauseBtn.IsTriggered()
}

func (gi *GameInputChecker) TriggeredAttack() bool {
	return gi.AttackBtn.IsTriggered()
}

func (gi *GameInputChecker) TriggeredSkill() bool {
	return gi.SkillBtn.IsTriggered()
}
