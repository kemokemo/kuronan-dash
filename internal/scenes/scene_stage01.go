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
	"github.com/kemokemo/kuronan-dash/internal/ui"
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
	sumTicks  float64
	msgWindow *ui.MessageWindow
	uiMsg     string
}

// Initialize initializes all resources.
func (s *Stage01Scene) Initialize() error {
	s.goalX = 1500.0
	s.timeLimit = 90
	s.time = s.timeLimit
	s.disc = music.Stage01

	s.player = chara.Selected
	lanes := field.NewLanes(field.PrairieLane)
	err := s.player.InitializeWithLanes(lanes)
	if err != nil {
		return err
	}

	s.field = &field.PrairieField{}
	s.field.Initialize(lanes)

	heights := lanes.GetLaneHeights()
	lowerLaneY := heights[len(heights)-1]
	s.msgWindow = ui.NewMessageWindow(
		300,
		int(lowerLaneY)+windowMargin+5,
		680,
		view.ScreenHeight-windowMargin*2-int(lowerLaneY),
		frameWidth)
	s.msgWindow.SetColors(
		color.RGBA{64, 64, 64, 255},
		color.RGBA{192, 192, 192, 255},
		color.RGBA{33, 228, 68, 255})

	return nil
}

// Update updates the status of this scene.
func (s *Stage01Scene) Update(state *GameState) {
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
			s.run()
		}
	case pause:
		if input.TriggeredOne() {
			s.state = run
			s.player.ReStart()
		}
	case stageClear:
		if input.TriggeredOne() {
			// TODO: goto next stage :-)
			err := state.SceneManager.GoTo(&TitleScene{})
			if err != nil {
				log.Println("failed to go to the 2nd stage: ", err)
			}
		}
	case gameOver:
		if input.TriggeredOne() {
			err := state.SceneManager.GoTo(&TitleScene{})
			if err != nil {
				log.Println("failed to go to the title screen: ", err)
			}
		}
	default:
		log.Println("unknown state of Stage01Scene:", s.state)
	}
}

// run works with 'run' state.
func (s *Stage01Scene) run() {
	s.sumTicks += 1.0 / ebiten.CurrentTPS()
	if s.sumTicks >= 1.0 {
		s.sumTicks = 0.0
		s.time--
	}

	isTimeUp := s.time <= 0
	isArriveGoal := s.player.GetPosition().X-view.DrawPosition > s.goalX

	if isArriveGoal {
		s.state = stageClear
		s.player.Pause()
	} else if !isArriveGoal && isTimeUp {
		s.state = gameOver
		s.player.Pause()
	} else {
		s.player.Update()
		s.field.Update(s.player.GetScrollVelocity())

		// TODO: プレイヤーの攻撃が障害物に当たっているか判定しつつ、当たっていればダメージを加える処理

		pRect := s.player.GetRectangle()
		s.player.BeBlocked(s.field.IsCollidedWithObstacles(pRect))
		s.player.Eat(s.field.EatFoods(pRect))
	}
}

// Draw draws background and characters.
func (s *Stage01Scene) Draw(screen *ebiten.Image) {
	s.field.DrawFarther(screen)
	s.player.Draw(screen)
	s.field.DrawCloser(screen)
	s.drawUI(screen)
	s.drawWithState(screen)
}

// description
func (s *Stage01Scene) drawUI(screen *ebiten.Image) {
	s.uiMsg = fmt.Sprintf("スタミナ: %v\nタイム: 　%v\nすすんだきょり: %.1f\nゴールのいち: 　%.1f",
		s.player.GetStamina(),
		s.time,
		s.player.GetPosition().X-view.DrawPosition,
		s.goalX,
	)
	s.msgWindow.DrawWindow(screen, s.uiMsg)

	text.Draw(screen, fmt.Sprintf("Now Playing: %s", s.disc.Name),
		fonts.GamerFontS, 12, view.ScreenHeight-10, color.White)
}

func (s *Stage01Scene) drawWithState(screen *ebiten.Image) {
	switch s.state {
	case wait:
		text.Draw(screen, messages.GameStart, fonts.GamerFontL, view.ScreenWidth/2-300, view.ScreenHeight/2, color.White)
	case pause:
		text.Draw(screen, messages.GamePause, fonts.GamerFontL, view.ScreenWidth/2-150, view.ScreenHeight/2, color.White)
	case stageClear:
		text.Draw(screen, messages.GameStageClear, fonts.GamerFontL, view.ScreenWidth/2-200, view.ScreenHeight/2-25, color.White)
		text.Draw(screen, messages.GameStageClear2, fonts.GamerFontL, view.ScreenWidth/2-300, view.ScreenHeight/2+25, color.White)
	case gameOver:
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
