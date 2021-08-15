package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Scene is interface for the all scenes.
type Scene interface {
	Initialize() error
	Update(state *GameState)
	Draw(screen *ebiten.Image)
	StartMusic()
	StopMusic() error
}
