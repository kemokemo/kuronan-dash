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
	AttackBtn    vpad.TriggerButton
	SkillBtn     vpad.TriggerButton
	DoubleClk    *DoubleClick
	currentIndex int
	mousePos     image.Point
	isUp         bool
	isDown       bool
	isAttack     bool
	tIDs         []ebiten.TouchID
	touchPos     image.Point
}

func (i *PlayerInputChecker) Update() {
	i.AttackBtn.Update()
	i.SkillBtn.Update()

	// 使用頻度が高そうな操作系からチェック。上下移動があればこの関数の処理を終える。
	i.isUp = false
	i.isDown = false

	// Button
	i.UpBtn.Update()
	if i.UpBtn.IsTriggered() {
		i.isUp = true
	}
	i.DownBtn.Update()
	if i.DownBtn.IsTriggered() {
		i.isDown = true
	}
	if i.isUp || i.isDown {
		return
	}

	i.isAttack = false
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
		}
	}
	if i.isUp || i.isDown {
		return
	}
	i.isAttack = inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight)

	// Touches
	i.tIDs = inpututil.AppendJustPressedTouchIDs(nil)
	for _, tID := range i.tIDs {
		i.touchPos.X, i.touchPos.Y = ebiten.TouchPosition(tID)
		for rIndex := range i.RectArray {
			if !i.touchPos.In(i.RectArray[rIndex]) || rIndex == i.currentIndex {
				continue
			}
			if i.currentIndex > rIndex {
				i.isUp = true
				i.currentIndex--
			} else {
				i.isDown = true
				i.currentIndex++
			}
		}
	}
	if i.isUp || i.isDown {
		return
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
	return i.AttackBtn.IsTriggered() || i.isAttack
}

func (i *PlayerInputChecker) TriggeredSkill() bool {
	return i.SkillBtn.IsTriggered() || i.DoubleClk.Triggered()
}
