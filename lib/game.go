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
	character    *objects.Character
	jukeBox      *music.JukeBox
}

// Init loads resources.
func (g *Game) Init() error {
	var err error
	g.jukeBox, err = music.NewJukeBox()
	if err != nil {
		return err
	}
	err = g.jukeBox.SelectDisc(music.Title)
	if err != nil {
		return err
	}

	err = g.loadsCharacters()
	if err != nil {
		return err
	}
	return nil
}

func (g *Game) loadsCharacters() error {
	var err error
	g.charaManager, err = objects.NewCharacterManager()
	if err != nil {
		return err
	}
	g.character = g.charaManager.GetSelectedCharacter()
	g.character.SetInitialPosition(objects.Position{X: 10, Y: 10})
	return nil
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
	if g.sceneManager == nil {
		g.sceneManager = &scenes.SceneManager{}
		g.sceneManager.SetResources(g.jukeBox, g.charaManager)
		g.sceneManager.GoTo(&scenes.TitleScene{})
	}

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
