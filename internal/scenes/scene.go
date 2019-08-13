package scenes

import (
	"github.com/hajimehoshi/ebiten"
)

// Scene is interface for the all scenes.
type Scene interface {
	Initialize() error
	Update(state *GameState) error
	Draw(screen *ebiten.Image) error
	StartMusic() error
	StopMusic() error
}
