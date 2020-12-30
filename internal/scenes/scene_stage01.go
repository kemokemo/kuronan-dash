// Copy from github.com/hajimehoshi/ebiten/v2/example/blocks

package scenes

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/kemokemo/kuronan-dash/assets/fonts"
	"github.com/kemokemo/kuronan-dash/assets/messages"
	"github.com/kemokemo/kuronan-dash/assets/music"

	chara "github.com/kemokemo/kuronan-dash/internal/character"
	"github.com/kemokemo/kuronan-dash/internal/field"
	"github.com/kemokemo/kuronan-dash/internal/input"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// Stage01Scene is the scene for the 1st stage game.
type Stage01Scene struct {
	state     gameState
	player    *chara.Player
	disc      *music.Disc
	field     field.Field
	goalX     float64
	timeLimit int // second
	time      int // second
}

// Initialize initializes all resources.
func (s *Stage01Scene) Initialize() error {
	s.goalX = 900.0
	s.timeLimit = 600
	s.time = s.timeLimit
	s.disc = music.Stage01
	s.player = chara.Selected
	err := s.player.InitializeWithLanesInfo(field.LaneHeights)
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
		if input.TriggeredOne() {
			s.state = run
			s.player.Start()
		}
	case run:
		if input.TriggeredOne() {
			s.state = pause
			s.player.Pause()
		} else {
			err = s.run()
		}
	case pause:
		if input.TriggeredOne() {
			s.state = run
			s.player.ReStart()
		}
	case stageClear:
		if input.TriggeredOne() {
			// TODO: goto next stage :-)
			state.SceneManager.GoTo(&TitleScene{})
		}
	case gameover:
		if input.TriggeredOne() {
			state.SceneManager.GoTo(&TitleScene{})
		}
	default:
		log.Println("unknown state of Stage01Scene:", s.state)
	}
	return err
}

// run works with 'run' state.
func (s *Stage01Scene) run() error {
	s.time--
	isTimeUp := s.time <= 0
	isArriveGoal := s.player.GetPosition().X > s.goalX

	var err error
	if isArriveGoal {
		s.state = stageClear
		s.player.Pause()
	} else if !isArriveGoal && isTimeUp {
		s.state = gameover
		s.player.Pause()
	} else {
		err = s.player.Update()
		if err != nil {
			log.Println("failed to update the player:", err)
			return err
		}
		s.field.Update(s.player.GetVelocity())

		// TODO: プレイヤーの攻撃が障害物に当たっているか判定しつつ、当たっていればダメージを加える処理

		s.player.BeBlocked(s.field.IsCollidedWithObstacles(s.player.GetRectangle()))
		s.player.Eat(s.field.EatFoods(s.player.GetRectangle()))
	}
	return err
}

// Draw draws background and characters.
func (s *Stage01Scene) Draw(screen *ebiten.Image) {
	pOffset := s.player.GetOffset()
	s.field.DrawFarther(screen, pOffset)
	s.player.Draw(screen)
	s.field.DrawCloser(screen, pOffset)
	s.drawUI(screen)
	s.drawWithState(screen)
}

// description
func (s *Stage01Scene) drawUI(screen *ebiten.Image) error {
	text.Draw(screen, fmt.Sprintf("Now Playing: %s", s.disc.Name),
		fonts.GamerFontS, 12, 35, color.White)

	text.Draw(screen, fmt.Sprintf("スタミナ: %v", s.player.GetStamina()),
		fonts.GamerFontS, 12, 60, color.White)

	text.Draw(screen, fmt.Sprintf("タイム: %v", s.time),
		fonts.GamerFontS, 160, 60, color.White)

	text.Draw(screen, fmt.Sprintf("すすんだきょり/ゴールいち: %v / %v", s.player.GetPosition().X, s.goalX),
		fonts.GamerFontS, 300, 60, color.White)

	return nil
}

func (s *Stage01Scene) drawWithState(screen *ebiten.Image) {
	switch s.state {
	case wait:
		text.Draw(screen, messages.GameStart, fonts.GamerFontL, view.ScreenWidth/2-250, view.ScreenHeight/2, color.White)
	case pause:
		text.Draw(screen, messages.GamePause, fonts.GamerFontL, view.ScreenWidth/2-150, view.ScreenHeight/2, color.White)
	case stageClear:
		text.Draw(screen, messages.GameStageClear, fonts.GamerFontL, view.ScreenWidth/2-200, view.ScreenHeight/2-25, color.White)
		text.Draw(screen, messages.GameStageClear2, fonts.GamerFontL, view.ScreenWidth/2-300, view.ScreenHeight/2+25, color.White)
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
func (s *Stage01Scene) StartMusic() {
	s.disc.Play()
}

// StopMusic stops playing music
func (s *Stage01Scene) StopMusic() error {
	return s.disc.Stop()
}
