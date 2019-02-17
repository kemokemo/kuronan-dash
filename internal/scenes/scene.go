package scenes

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/kuronan-dash/internal/character"
	"github.com/kemokemo/kuronan-dash/internal/music"
)

// Scene is interface for the all scenes.
type Scene interface {
	SetResources(j *music.JukeBox, cm *character.Manager)
	Update(state *GameState) error
	Draw(screen *ebiten.Image)
}
