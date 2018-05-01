package kuronandash

import (
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
	err := g.loadsMusic()
	if err != nil {
		return err
	}

	err = g.loadsCharacters()
	if err != nil {
		return err
	}
	return nil
}

func (g *Game) loadsMusic() error {
	g.jukeBox = music.NewJukeBox()
	err := g.jukeBox.InsertDiscs([]music.RequestCard{
		music.RequestCard{
			FilePath: "_assets/music/shibugaki_no_kuroneko.mp3",
		},
		music.RequestCard{
			FilePath: "_assets/music/hashire_kurona.mp3",
		},
	})
	if err != nil {
		return err
	}
	err = g.jukeBox.SelectDisc("shibugaki_no_kuroneko")
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

// Close closes own resources.
func (g *Game) Close() error {
	return g.jukeBox.Close()
}

// Update is an implements to draw screens.
func (g *Game) Update(screen *ebiten.Image) error {
	if g.sceneManager == nil {
		g.sceneManager = &scenes.SceneManager{}
		g.sceneManager.SetResources(g.jukeBox, g.character)
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
