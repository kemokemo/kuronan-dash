// Copy from github.com/hajimehoshi/ebiten/example/blocks

package scenes

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"

	"github.com/kemokemo/kuronan-dash/assets/fonts"
	"github.com/kemokemo/kuronan-dash/assets/messages"
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
	goalX  float64
}

// Initialize initializes all resources.
func (s *Stage01Scene) Initialize() error {
	s.goalX = 600.0
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
	var err error
	switch s.state {
	case wait:
		if state.Input.StateForKey(ebiten.KeySpace) == 1 {
			s.state = run
			s.player.Start()
		}
	case run:
		err = s.run(state)
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
		log.Println("unknown state of Stage01Scene:", s.state)
	}
	return err
}

// run works with 'run' state.
func (s *Stage01Scene) run(state *GameState) error {
	var err error
	if state.Input.StateForKey(ebiten.KeySpace) == 1 {
		s.state = pause
		s.player.Pause()
	} else if s.player.GetPosition().X > view.ScreenWidth+s.goalX && s.state != gameover {
		// TODO: ゴールとプレイヤーの衝突有無や、経過時間が制限時間以内かをチェックした結果でゲームオーバー判定を行う
		s.state = gameover
		s.player.Pause()
	} else {
		// 位置の更新
		err = s.player.Update()
		if err != nil {
			log.Println("failed to update the player:", err)
		}
		s.field.Update(s.player.GetVelocity())

		// TODO: プレイヤーの攻撃が障害物に当たっているか判定しつつ、当たっていればダメージを加える処理

		s.player.BeBlocked(s.field.IsCollidedWithObstacles(s.player.GetRectangle()))
		s.player.Eat(s.field.EatFoods(s.player.GetRectangle()))
	}
	return err
}

// Draw draws background and characters.
func (s *Stage01Scene) Draw(screen *ebiten.Image) error {
	pOffset := s.player.GetOffset()
	err := s.field.DrawFarther(screen, pOffset)
	if err != nil {
		return fmt.Errorf("failed to draw the farther field parts,%v", err)
	}

	err = s.player.Draw(screen)
	if err != nil {
		return fmt.Errorf("failed to draw a character,%v", err)
	}

	err = s.field.DrawCloser(screen, pOffset)
	if err != nil {
		return fmt.Errorf("failed to draw the closer field parts,%v", err)
	}

	text.Draw(screen, fmt.Sprintf("Now Playing: %s", s.disc.Name),
		fonts.GamerFontS, 12, 35, color.White)

	text.Draw(screen, fmt.Sprintf("スタミナ: %v", s.player.GetStamina()),
		fonts.GamerFontS, 12, 60, color.White)

	s.drawWithState(screen)
	return nil
}

func (s *Stage01Scene) drawWithState(screen *ebiten.Image) {
	switch s.state {
	case wait:
		text.Draw(screen, messages.GameStart, fonts.GamerFontL, view.ScreenWidth/2-250, view.ScreenHeight/2, color.White)
	case pause:
		text.Draw(screen, messages.GamePause, fonts.GamerFontL, view.ScreenWidth/2-150, view.ScreenHeight/2, color.White)
	case gameover:
		text.Draw(screen, messages.GameOver, fonts.GamerFontL, view.ScreenWidth/2-420, view.ScreenHeight/2, color.White)
	default:
		// nothing to draw
	}
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
