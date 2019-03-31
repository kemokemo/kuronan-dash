// Copy from github.com/hajimehoshi/ebiten/example/blocks

package scenes

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	mplus "github.com/hajimehoshi/go-mplusbitmap"

	"github.com/kemokemo/kuronan-dash/assets/music"

	chara "github.com/kemokemo/kuronan-dash/internal/character"
)

// Stage01Scene is the scene for the 1st stage game.
type Stage01Scene struct {
	state  state
	player *chara.Player
	disc   *music.Disc
}

// Initialize initializes all resources.
func (s *Stage01Scene) Initialize() error {
	s.disc = music.Stage01
	s.player = chara.Selected
	s.player.SetInitialPosition(chara.Position{X: 10, Y: 50})
	return nil
}

// Update updates the status of this scene.
func (s *Stage01Scene) Update(state *GameState) error {
	s.updateStatus(state)
	return nil
}

// Draw draws background and characters.
func (s *Stage01Scene) Draw(screen *ebiten.Image) {
	err := s.player.Draw(screen)
	if err != nil {
		log.Println("failed to draw a character:", err)
		return
	}
	text.Draw(screen, fmt.Sprintf("Now Playing: %s", s.disc.Name),
		mplus.Gothic12r, 12, 15, color.White)

	s.drawWithState(screen)
	// TODO: 衝突判定とSE再生
	err = s.checkCollision()
	if err != nil {
		log.Println("failed to check collisions:", err)
		return
	}
}

func (s *Stage01Scene) updateStatus(state *GameState) error {
	switch s.state {
	case beforeRun:
		if state.Input.StateForKey(ebiten.KeySpace) == 1 {
			s.state = running
		}
	case running:
		s.player.Move()
		if state.Input.StateForKey(ebiten.KeySpace) == 1 {
			s.state = pause
		}
		// TODO: とりあえずゲームオーバーの練習
		if s.player.Position.X+50 > ScreenWidth-50 && s.state != gameover {
			s.state = gameover
		}
	case pause:
		if state.Input.StateForKey(ebiten.KeySpace) == 1 {
			s.state = running
		}
	case gameover:
		if state.Input.StateForKey(ebiten.KeySpace) == 1 {
			state.SceneManager.GoTo(&TitleScene{})
		}
	default:
		s.player.Move()
		// unknown state
	}
	return nil
}

func (s *Stage01Scene) drawWithState(screen *ebiten.Image) {
	switch s.state {
	case beforeRun:
		text.Draw(screen, "Spaceキーを押してゲームを開始してね！", mplus.Gothic12r, ScreenWidth/2-100, ScreenHeight/2, color.White)
	case pause:
		text.Draw(screen, "一時停止中だよ！", mplus.Gothic12r, ScreenWidth/2-100, ScreenHeight/2, color.White)
	case gameover:
		text.Draw(screen, "ゲームオーバー！Spaceを押してタイトルへ戻ってね！", mplus.Gothic12r, ScreenWidth/2-100, ScreenHeight/2, color.White)
	default:
		// nothing to draw
	}
}

func (s *Stage01Scene) checkCollision() error {
	// TODO: 衝突判定の代わりにボタン入力
	if ebiten.IsKeyPressed(ebiten.KeyJ) {
		s.player.SetState(chara.Ascending)
	} else {
		s.player.SetState(chara.Dash)
	}
	err := s.player.PlaySe()
	if err != nil {
		return err
	}
	return nil
}

// Close stops music
func (s *Stage01Scene) Close() error {
	err := s.disc.Stop()
	if err != nil {
		return fmt.Errorf("failed to stop music:%v", err)
	}
	return nil
}

// StartMusic starts playing music
func (s *Stage01Scene) StartMusic() error {
	return s.disc.Play()
}

// StopMusic stops playing music
func (s *Stage01Scene) StopMusic() error {
	return s.disc.Stop()
}
