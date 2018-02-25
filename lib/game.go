package kuronandash

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
)

const sampleRate = 44100

// Game controls all things in the screen.
type Game struct {
	sceneManager *SceneManager
	input        Input
	character    *Character
	jukeBox      *JukeBox
}

// Init loads resources.
func (g *Game) Init() error {
	context, err := audio.NewContext(sampleRate)
	if err != nil {
		return err
	}

	err = g.loadsMusic(context)
	if err != nil {
		return err
	}

	err = g.loadsCharacters(context)
	if err != nil {
		return err
	}
	return nil
}

func (g *Game) loadsMusic(context *audio.Context) error {
	g.jukeBox = NewJukeBox(context)
	err := g.jukeBox.InsertDiscs([]RequestCard{
		RequestCard{
			FilePath: "assets/music/shibugaki_no_kuroneko.mp3",
		},
		RequestCard{
			FilePath: "assets/music/hashire_kurona.mp3",
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

func (g *Game) loadsCharacters(context *audio.Context) error {
	var err error
	g.character, err = NewCharacter(context, []string{
		"assets/images/character/koma_00.png",
		"assets/images/character/koma_01.png",
		"assets/images/character/koma_02.png",
		"assets/images/character/koma_03.png",
	})
	if err != nil {
		log.Println("Failed to load assets.", err)
		return err
	}
	g.character.SetInitialPosition(Position{X: 10, Y: 10})
	return nil
}

// Close closes own resources.
func (g *Game) Close() error {
	return g.jukeBox.Close()
}

// Update is an implements to draw screens.
func (g *Game) Update(screen *ebiten.Image) error {
	if g.sceneManager == nil {
		g.sceneManager = &SceneManager{}
		g.sceneManager.SetResources(g.jukeBox, g.character)
		g.sceneManager.GoTo(&TitleScene{})
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
