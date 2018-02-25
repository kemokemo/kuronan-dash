// Copy from github.com/hajimehoshi/ebiten/example/blocks

package kuronandash

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	mplusbitmap "github.com/hajimehoshi/go-mplusbitmap"
)

type gameState int

const (
	beforeRun gameState = iota
	running
	pause
	gameover
)

// GameScene is the scene for the game.
type GameScene struct {
	state     gameState
	character *Character
	jukeBox   *JukeBox
}

// NewGameScene creates the new GameScene.
func NewGameScene() *GameScene {
	return &GameScene{
		state: beforeRun,
	}
}

// SetResources sets the resources like music, character images and so on.
func (s *GameScene) SetResources(j *JukeBox, c *Character) {
	s.jukeBox = j
	s.character = c
	s.character.SetInitialPosition(Position{X: 10, Y: 50})
	err := s.jukeBox.SelectDisc("hashire_kurona")
	if err != nil {
		log.Printf("Failed to select disc:%v", err)
	}
}

// Update updates the status of this scene.
func (s *GameScene) Update(state *GameState) error {
	s.updateStatus(state)
	return nil
}

// Draw draws background and characters. This function play music too.
func (s *GameScene) Draw(screen *ebiten.Image) {
	err := s.jukeBox.Play()
	if err != nil {
		log.Printf("Failed to play with JukeBox:%v", err)
		return
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Now Playing: %s", s.jukeBox.NowPlaying()))
	err = s.character.Draw(screen)
	if err != nil {
		log.Printf("Failed to draw character:%v", err)
		return
	}

	if s.state == gameover {
		text.Draw(screen, "ゲームオーバー: Spaceを押してタイトルへ", mplusbitmap.Gothic12r, ScreenWidth/2-100, ScreenHeight/2, color.White)
		return
	}

	// TODO: 衝突判定とSE再生
	err = s.checkCollision()
	if err != nil {
		log.Printf("Failed to check collision:%v", err)
		return
	}
}

func (s *GameScene) updateStatus(state *GameState) error {
	// TODO: とりあえずゲームオーバーの練習
	if s.state == gameover {
		if state.Input.StateForKey(ebiten.KeySpace) == 1 {
			state.SceneManager.GoTo(&TitleScene{})
			return nil
		}
		return nil
	}
	if s.character.position.X+50 > ScreenWidth-50 && s.state != gameover {
		s.state = gameover
		return nil
	}

	s.character.Move()
	return nil
}

func (s *GameScene) checkCollision() error {
	// TODO: 衝突判定の代わりにボタン入力
	if ebiten.IsKeyPressed(ebiten.KeyJ) {
		s.character.SetState(Ascending)
	} else {
		s.character.SetState(Dash)
	}
	err := s.character.PlaySe()
	if err != nil {
		return err
	}
	return nil
}
