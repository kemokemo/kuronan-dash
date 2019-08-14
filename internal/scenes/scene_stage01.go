// Copy from github.com/hajimehoshi/ebiten/example/blocks

package scenes

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	mplus "github.com/hajimehoshi/go-mplusbitmap"

	"github.com/kemokemo/kuronan-dash/assets/music"

	chara "github.com/kemokemo/kuronan-dash/internal/character"
	"github.com/kemokemo/kuronan-dash/internal/field"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// Stage01Scene is the scene for the 1st stage game.
type Stage01Scene struct {
	state  gameState
	player *chara.Player
	disc   *music.Disc
	field  field.Field
}

// Initialize initializes all resources.
func (s *Stage01Scene) Initialize() error {
	s.disc = music.Stage01
	s.player = chara.Selected
	err := s.player.SetLanes(field.LaneHeights)
	if err != nil {
		return err
	}
	s.field = &field.PrairieField{}
	s.field.Initialize()
	return nil
}

// Update updates the status of this scene.
func (s *Stage01Scene) Update(state *GameState) error {
	switch s.state {
	case wait:
		if state.Input.StateForKey(ebiten.KeySpace) == 1 {
			s.state = run
			s.player.Start()
		}
	case run:
		if state.Input.StateForKey(ebiten.KeySpace) == 1 {
			s.state = pause
			s.player.Pause()
		} else if s.player.Position.X+50 > view.ScreenWidth-50 && s.state != gameover {
			// TODO: とりあえずゲームオーバーの練習
			s.state = gameover
			s.player.Stop()
		} else {
			s.player.Update()
			if s.player.GetState() == chara.Dash {
				s.field.SetScrollSpeed(field.Normal)
			} else {
				s.field.SetScrollSpeed(field.Slow)
			}
			s.field.Update()
		}
	case pause:
		if state.Input.StateForKey(ebiten.KeySpace) == 1 {
			s.state = run
			s.player.ReStart()
		}
	case gameover:
		if state.Input.StateForKey(ebiten.KeySpace) == 1 {
			state.SceneManager.GoTo(&TitleScene{})
		}
	default:
		// unknown state
	}
	return nil
}

// Draw draws background and characters.
func (s *Stage01Scene) Draw(screen *ebiten.Image) error {
	err := s.field.Draw(screen)
	if err != nil {
		return fmt.Errorf("failed to draw field parts,%v", err)
	}

	err = s.player.Draw(screen)
	if err != nil {
		return fmt.Errorf("failed to draw a character,%v", err)
	}

	text.Draw(screen, fmt.Sprintf("Now Playing: %s", s.disc.Name),
		mplus.Gothic12r, 12, 15, color.White)

	s.drawWithState(screen)
	// TODO: 衝突判定とSE再生
	err = s.checkCollision()
	if err != nil {
		return fmt.Errorf("failed to check collisions,%v", err)
	}
	return nil
}

func (s *Stage01Scene) drawWithState(screen *ebiten.Image) {
	switch s.state {
	case wait:
		text.Draw(screen, "Spaceキーを押してゲームを開始してね！", mplus.Gothic12r, view.ScreenWidth/2-100, view.ScreenHeight/2, color.White)
	case pause:
		text.Draw(screen, "一時停止中だよ！", mplus.Gothic12r, view.ScreenWidth/2-100, view.ScreenHeight/2, color.White)
	case gameover:
		text.Draw(screen, "ゲームオーバー！Spaceを押してタイトルへ戻ってね！", mplus.Gothic12r, view.ScreenWidth/2-100, view.ScreenHeight/2, color.White)
	default:
		// nothing to draw
	}
}

func (s *Stage01Scene) checkCollision() error {
	// TODO: プレイヤーと障害物との衝突判定などをするよ
	return nil
}

// Close stops music
func (s *Stage01Scene) Close() error {
	err := s.disc.Stop()
	if err != nil {
		return fmt.Errorf("failed to stop music,%v", err)
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
