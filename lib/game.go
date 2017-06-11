package kuronandash

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

// Game controls all things in the screen.
type Game struct {
	character Character
}

// Init loads resources.
func (g *Game) Init() error {
	g.character = Character{
		ImagesPaths: []string{
			"images/game/koma_00.png",
			"images/game/koma_01.png",
			"images/game/koma_02.png",
			"images/game/koma_03.png",
		},
	}

	err := g.character.Init()
	if err != nil {
		log.Println("Failed to load assets.", err)
		return err
	}
	g.character.SetInitialPosition(Position{X: 10, Y: 10})

	return nil
}

// Update is an implements to draw screens.
func (g *Game) Update(screen *ebiten.Image) error {
	// First of all, all characters move.
	g.character.Move()
	if ebiten.IsRunningSlowly() {
		return nil
	}

	// If not in slow mode, perform various drawing.
	err := g.character.Draw(screen)
	if err != nil {
		return err
	}
	return nil
}
