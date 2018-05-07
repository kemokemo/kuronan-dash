package kuronandash

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/kuronan-dash/lib/music"
	"github.com/kemokemo/kuronan-dash/lib/objects"
	"github.com/kemokemo/kuronan-dash/lib/scenes"
	"github.com/kemokemo/kuronan-dash/lib/util"
)

// Game controls all things in the screen.
type Game struct {
	sceneManager *scenes.SceneManager
	input        util.Input
	charaManager *objects.CharacterManager
	jukeBox      *music.JukeBox
}

// NewGame returns a new game instance.
// Please call the Close method when you no longer use this instance.
func NewGame() (*Game, error) {
	g := Game{}
	var err error
	g.jukeBox, err = music.NewJukeBox()
	if err != nil {
		return nil, err
	}
	g.charaManager, err = objects.NewCharacterManager()
	if err != nil {
		return nil, err
	}

	g.sceneManager, err = scenes.NewSceneManager()
	if err != nil {
		return nil, err
	}
	g.sceneManager.SetResources(g.jukeBox, g.charaManager)
	g.sceneManager.GoTo(&scenes.TitleScene{})
	return &g, nil
}

// Close closes inner resources.
func (g *Game) Close() error {
	var err, e error
	e = g.jukeBox.Close()
	if e != nil {
		err = fmt.Errorf("%v %v", err, e)
	}
	e = g.charaManager.Close()
	if e != nil {
		err = fmt.Errorf("%v %v", err, e)
	}
	return err
}

// Update is an implements to draw screens.
func (g *Game) Update(screen *ebiten.Image) error {
	g.input.Update()
	if err := g.sceneManager.Update(&g.input); err != nil {
		return err
	}
	// First of all, updates all status.
	if ebiten.IsRunningSlowly() {
		return nil
	}
	g.sceneManager.Draw(screen)
	return nil
}
