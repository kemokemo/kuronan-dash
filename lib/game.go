package kuronandash

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const sampleRate = 44100

type gameState int

const (
	normal gameState = iota
	running
	slow
)

// Game controls all things in the screen.
type Game struct {
	state     gameState
	character *Character
	jukeBox   *JukeBox
}

// Init loads resources.
func (g *Game) Init() error {
	context, err := audio.NewContext(sampleRate)
	if err != nil {
		return err
	}

	err = g.initCharacters(context)
	if err != nil {
		return err
	}

	g.jukeBox = NewJukeBox(context)
	err = g.jukeBox.InsertDiscs([]RequestCard{
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

	g.state = normal

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
	// First of all, updates all status.
	g.updateStatus()
	if ebiten.IsRunningSlowly() {
		return nil
	}
	// If not in slow mode, perform various playing and drawing.
	err := g.jukeBox.Play()
	if err != nil {
		return err
	}
	err = g.character.Draw(screen)
	if err != nil {
		return err
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Now Playing: %s", g.jukeBox.NowPlaying()))
	// TODO: 衝突判定とSE再生
	err = g.checkCollision()
	if err != nil {
		return err
	}
	return nil
}

func (g *Game) updateStatus() error {
	g.character.Move()

	// TODO: ボタン入力で代用
	// TODO: state変わったら曲を更新する
	if ebiten.IsKeyPressed(ebiten.Key0) {
		g.state = normal
		return nil
	} else if ebiten.IsKeyPressed(ebiten.Key1) {
		if g.state != slow {
			g.state = slow
			err := g.jukeBox.SelectDisc("shibugaki_no_kuroneko")
			if err != nil {
				return err
			}
			return nil
		}
	} else if ebiten.IsKeyPressed(ebiten.Key2) {
		if g.state != running {
			g.state = running
			err := g.jukeBox.SelectDisc("hashire_kurona")
			if err != nil {
				return err
			}
			return nil
		}
	} else {
		// nothig
	}
	return nil
}

func (g *Game) checkCollision() error {
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

func (g *Game) initCharacters(context *audio.Context) error {
	var err error
	g.character, err = NewCharacter(context, []string{
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
