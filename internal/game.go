package kuronandash

import (
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/kemokemo/kuronan-dash/assets"

	"github.com/kemokemo/kuronan-dash/internal/character"
	"github.com/kemokemo/kuronan-dash/internal/scenes"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// Game controls all things in the screen.
type Game struct {
	scenes *scenes.SceneManager
	once   sync.Once
}

// NewGame returns a new game instance.
// Please call the Close method when you no longer use this instance.
func NewGame(ver string) (*Game, error) {
	err := assets.LoadAssets()
	if err != nil {
		return nil, err
	}

	err = character.NewPlayers()
	if err != nil {
		return nil, err
	}

	sm := scenes.NewSceneManager(ver)
	return &Game{
		scenes: sm,
	}, nil
}

func (g *Game) Update() error {
	g.once.Do(func() {
		err := g.scenes.GoTo(&scenes.TitleScene{})
		if err != nil {
			log.Println("failed to run title screen")
		}
	})
	g.scenes.Update()

	return nil
}

// Update is an implements to draw screens.
func (g *Game) Draw(screen *ebiten.Image) {
	g.scenes.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return view.ScreenWidth, view.ScreenHeight
}

// Close closes inner resources.
func (g *Game) Close() error {
	g.scenes.Close()
	return assets.CloseAssets()
}
