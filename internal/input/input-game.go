package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GameInputChecker struct{}

func (gi *GameInputChecker) Update() {}

func (gi *GameInputChecker) TriggeredUp() bool {
	return false
}

func (gi *GameInputChecker) TriggeredDown() bool {
	return false
}

func (gi *GameInputChecker) TriggeredLeft() bool {
	return false
}

func (gi *GameInputChecker) TriggeredRight() bool {
	return false
}

func (gi *GameInputChecker) TriggeredStart() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeySpace) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
}

func (gi *GameInputChecker) TriggeredPause() bool {
	// todo: マウスの右クリックでPauseじゃなくて、別途Pauseボタンを設けて、それを押したらという動作にしたい
	return inpututil.IsKeyJustReleased(ebiten.KeySpace) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight)
}

func (gi *GameInputChecker) TriggeredAttack() bool {
	return false
}

func (gi *GameInputChecker) TriggeredSpecial() bool {
	return false
}
