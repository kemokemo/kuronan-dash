// Copy from github.com/hajimehoshi/ebiten/example/blocks

package kuronandash

import (
	"fmt"
	_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type gameState int

const (
	normal gameState = iota
	running
	slow
)

type GameScene struct {
	state     gameState
	character *Character
	jukeBox   *JukeBox
}

func NewGameScene() *GameScene {
	return &GameScene{
		state: normal,
	}
}

func (s *GameScene) SetResources(j *JukeBox, c *Character) {
	s.jukeBox = j
	s.character = c
	err := s.jukeBox.SelectDisc("hashire_kurona")
	if err != nil {
		log.Printf("Failed to select disc:%v", err)
	}
}

func (s *GameScene) Update(state *GameState) error {
	s.updateStatus()
	return nil
}

func (s *GameScene) Draw(screen *ebiten.Image) {
	err := s.jukeBox.Play()
	if err != nil {
		log.Printf("Failed to play with JukeBox:%v", err)
		return
	}
	err = s.character.Draw(screen)
	if err != nil {
		log.Printf("Failed to draw character:%v", err)
		return
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Now Playing: %s", s.jukeBox.NowPlaying()))
	// TODO: 衝突判定とSE再生
	err = s.checkCollision()
	if err != nil {
		log.Printf("Failed to check collision:%v", err)
		return
	}
}

func (s *GameScene) updateStatus() error {
	s.character.Move()

	// TODO: ボタン入力で代用
	// TODO: state変わったら曲を更新する
	if ebiten.IsKeyPressed(ebiten.Key0) {
		s.state = normal
		return nil
	} else if ebiten.IsKeyPressed(ebiten.Key1) {
		if s.state != slow {
			s.state = slow
			err := s.jukeBox.SelectDisc("shibugaki_no_kuroneko")
			if err != nil {
				return err
			}
			return nil
		}
	} else if ebiten.IsKeyPressed(ebiten.Key2) {
		if s.state != running {
			s.state = running
			err := s.jukeBox.SelectDisc("hashire_kurona")
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
