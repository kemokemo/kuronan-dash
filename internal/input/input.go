package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func TriggeredOne() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeySpace)
}
