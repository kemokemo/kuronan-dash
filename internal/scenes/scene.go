package scenes

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/kuronan-dash/internal/music"
	"github.com/kemokemo/kuronan-dash/internal/objects"
)

// Scene is interface for the all scenes.
type Scene interface {
	SetResources(j *music.JukeBox, cm *objects.CharacterManager)
	Update(state *GameState) error
	Draw(screen *ebiten.Image)
}
