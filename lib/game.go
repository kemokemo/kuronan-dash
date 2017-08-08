package kuronandash

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
)

const sampleRate = 44100

// Game controls all things in the screen.
type Game struct {
	context   *audio.Context
	character *Character
	jukeBox   *JukeBox
}

// Init loads resources.
func (g *Game) Init() error {
	var err error
	g.context, err = audio.NewContext(sampleRate)
	if err != nil {
		return err
	}
	err = g.initCharacters()
	if err != nil {
		return err
	}
	g.jukeBox, err = NewJukeBox(g.context, "assets/music/hasire_kuroneko.mp3")
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
	err := g.context.Update()
	if err != nil {
		return err
	}
	err = g.jukeBox.Play(g.getCurrentBGM())
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
	// TODO: 衝突判定とSE再生
	err = g.CheckCollision()
	if err != nil {
		return err
	}
	return nil
}

// CheckCollision check the collision between a character and other objects.
func (g *Game) CheckCollision() error {
	// TODO: 衝突判定の代わりにボタン入力
	if ebiten.IsKeyPressed(ebiten.KeyJ) {
		g.character.SetState(Ascending)
	} else {
		g.character.SetState(Dash)
	}
	err := g.character.PlaySe()
	if err != nil {
		return err
	}
	return nil
}

func (g *Game) getCurrentBGM() string {
	// TODO: return a BGM name of the current game stage
	return ""
}

func (g *Game) initCharacters() error {
	var err error
	g.character, err = NewCharacter(g.context, []string{
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
