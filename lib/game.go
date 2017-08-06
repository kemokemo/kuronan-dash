package kuronandash

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

// Game controls all things in the screen.
type Game struct {
	character *Character
	jukeBox   *JukeBox
}

// Init loads resources.
func (g *Game) Init() error {
	err := g.initCharacters()
	if err != nil {
		return err
	}
	g.jukeBox, err = NewJukeBox("assets/music/hasire_kuroneko.mp3")
	if err != nil {
		return err
	}
	return nil
}

// Close closes own resources.
func (g *Game) Close() error {
	return g.jukeBox.Close()
}

// Update is an implements to draw screens.
func (g *Game) Update(screen *ebiten.Image) error {
	// First of all, all characters move.
	g.character.Move()
	err := g.jukeBox.Play("")
	if err != nil {
		return err
	}
	if ebiten.IsRunningSlowly() {
		return nil
	}
	// If not in slow mode, perform various drawing.
	err = g.character.Draw(screen)
	if err != nil {
		return err
	}
	return nil
}

func (g *Game) initCharacters() error {
	var err error
	g.character, err = NewCharacter([]string{
		"assets/images/koma_00.png",
		"assets/images/koma_01.png",
		"assets/images/koma_02.png",
		"assets/images/koma_03.png",
	})
	if err != nil {
		log.Println("Failed to load assets.", err)
		return err
	}
	g.character.SetInitialPosition(Position{X: 10, Y: 10})
	return nil
}
