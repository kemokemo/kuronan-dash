// Copy from github.com/hajimehoshi/ebiten/v2/example/blocks

package scenes

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/kemokemo/kuronan-dash/assets/fonts"
	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/assets/messages"
	"github.com/kemokemo/kuronan-dash/assets/music"

	vpad "github.com/kemokemo/ebiten-virtualpad"
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
	iChecker  input.InputChecker
	startBtn  vpad.TriggerButton
	pauseBtn  vpad.TriggerButton
	upBtn     vpad.TriggerButton
	downBtn   vpad.TriggerButton
	atkBtn    vpad.TriggerButton
	spBtn     vpad.TriggerButton
	pauseBg   *ebiten.Image
	pauseBgOp *ebiten.DrawImageOptions
}

// Initialize initializes all resources.
func (s *Stage01Scene) Initialize() error {
	s.goalX = 3000.0
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
	s.field.Initialize(lanes, s.goalX)

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

	laneRectArray := []image.Rectangle{}
	previousHeight := 0
	for index := range heights {
		laneRectArray = append(laneRectArray,
			image.Rectangle{
				Min: image.Point{X: 0, Y: previousHeight},
				Max: image.Point{X: view.ScreenWidth, Y: int(heights[index])},
			},
		)
		previousHeight = int(heights[index])
	}

	s.upBtn = vpad.NewTriggerButton(images.UpButton, vpad.JustPressed, vpad.SelectColor)
	bW, bH := images.UpButton.Size()
	s.upBtn.SetLocation(20, view.ScreenHeight-bH-45)
	s.downBtn = vpad.NewTriggerButton(images.DownButton, vpad.JustPressed, vpad.SelectColor)
	s.downBtn.SetLocation(20+bW+10, view.ScreenHeight-bH-45)
	s.atkBtn = vpad.NewTriggerButton(images.AttackButton, vpad.JustPressed, vpad.SelectColor)
	s.atkBtn.SetLocation(view.ScreenWidth-20-2*bW-10, view.ScreenHeight-bH-45)
	s.spBtn = vpad.NewTriggerButton(images.SpecialButton, vpad.JustPressed, vpad.SelectColor)
	s.spBtn.SetLocation(view.ScreenWidth-bW-20, view.ScreenHeight-bH-45)
	s.player.SetInputChecker(laneRectArray, s.upBtn, s.downBtn, s.atkBtn, s.spBtn)

	s.startBtn = vpad.NewTriggerButton(images.StartButton, vpad.JustPressed, vpad.SelectColor)
	s.startBtn.SetLocation(view.ScreenWidth/2-64, view.ScreenHeight/2-128)
	s.pauseBtn = vpad.NewTriggerButton(images.PauseButton, vpad.JustPressed, vpad.SelectColor)
	s.pauseBtn.SetLocation(view.ScreenWidth-58, 48)
	s.iChecker = &input.GameInputChecker{StartBtn: s.startBtn, PauseBtn: s.pauseBtn}

	s.pauseBg = images.PauseLayer
	s.pauseBgOp = &ebiten.DrawImageOptions{}

	return nil
}

// Update updates the status of this scene.
func (s *Stage01Scene) Update(state *GameState) {
	// s.upBtnとs.downBtnは、s.iChecker内でUpdate()されるのでここではしない
	s.iChecker.Update()

	switch s.state {
	case wait:
		if s.iChecker.TriggeredStart() {
			s.state = run
			s.player.Start()
			s.disc.Play()
		}
	case run:
		if s.iChecker.TriggeredPause() {
			s.state = pause
			s.player.Pause()
			s.disc.Pause()
		} else if s.player.StartSpEffect() {
			s.state = specialEffect
		} else {
			s.run()
		}
	case pause:
		if s.iChecker.TriggeredStart() {
			s.state = run
			s.player.ReStart()
			s.disc.Play()
		}
	case specialEffect:
		s.player.UpdateSpecialEffect()
		if s.player.FinishSpEffect() {
			s.state = run
		}
	case stageClear:
		if s.iChecker.TriggeredStart() {
			// TODO: goto next stage :-)
			err := state.SceneManager.GoTo(&TitleScene{})
			if err != nil {
				log.Println("failed to go to the 2nd stage: ", err)
			}
		}
	case gameOver:
		if s.iChecker.TriggeredStart() {
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
	s.sumTicks += ebiten.CurrentTPS()
	if s.sumTicks >= 3600 {
		s.sumTicks = 0.0
		s.time--
	}

	isTimeUp := s.time <= 0
	isArriveGoal := s.player.GetPosition().X-view.DrawPosition > s.goalX

	if isArriveGoal {
		s.state = stageClear
		s.player.Pause()
		s.disc.Pause()
	} else if !isArriveGoal && isTimeUp {
		s.state = gameOver
		s.player.Pause()
		s.disc.Pause()
	} else {
		s.player.Update()
		s.field.Update(s.player.GetScrollVelocity())

		isAtk, aRect, power := s.player.IsAttacked()
		if isAtk {
			collided, broken := s.field.AttackObstacles(aRect, power)
			if collided > 0 {
				s.player.ConsumeStaminaByAttack(collided)
			}
			if broken > 0 {
				s.player.AddTension(broken)
			}
		}

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
	s.uiMsg = fmt.Sprintf("スタミナ: %v\nテンション: %v\nタイム: 　%v\nすすんだきょり: %.1f\nゴールのいち: 　%.1f",
		s.player.GetStamina(),
		s.player.GetTension(),
		s.time,
		s.player.GetPosition().X-view.DrawPosition,
		s.goalX,
	)
	s.msgWindow.DrawWindow(screen, s.uiMsg)
}

func (s *Stage01Scene) drawWithState(screen *ebiten.Image) {
	s.upBtn.Draw(screen)
	s.downBtn.Draw(screen)
	s.atkBtn.Draw(screen)
	s.spBtn.Draw(screen)

	// TODO: StartとPauseのボタンは見えてないだけで、該当する場所を押せばボタンはトリガーされる。弊害がありそうなら処置する。
	switch s.state {
	case wait:
		screen.DrawImage(s.pauseBg, s.pauseBgOp)
		text.Draw(screen, messages.GameStart, fonts.GamerFontL, view.ScreenWidth/2-280, view.ScreenHeight/2+30, color.White)
		s.startBtn.Draw(screen)
	case pause:
		text.Draw(screen, fmt.Sprintf("Now Playing: %s", s.disc.Name),
			fonts.GamerFontS, 12, view.ScreenHeight-10, color.White)
		screen.DrawImage(s.pauseBg, s.pauseBgOp)
		text.Draw(screen, messages.GamePause, fonts.GamerFontL, view.ScreenWidth/2-150, view.ScreenHeight/2+30, color.White)
		s.startBtn.Draw(screen)
	case run:
		s.pauseBtn.Draw(screen)
		text.Draw(screen, fmt.Sprintf("Now Playing: %s", s.disc.Name),
			fonts.GamerFontS, 12, view.ScreenHeight-10, color.White)
	case specialEffect:
		s.player.DrawSpecialEffect(screen)
	case stageClear:
		text.Draw(screen, fmt.Sprintf("Now Playing: %s", s.disc.Name),
			fonts.GamerFontS, 12, view.ScreenHeight-10, color.White)
		screen.DrawImage(s.pauseBg, s.pauseBgOp)
		text.Draw(screen, messages.GameStageClear, fonts.GamerFontL, view.ScreenWidth/2-200, view.ScreenHeight/2-134, color.White)
		text.Draw(screen, messages.GameStageClear2, fonts.GamerFontL, view.ScreenWidth/2-500, view.ScreenHeight/2+30, color.White)
		s.startBtn.Draw(screen)
	case gameOver:
		text.Draw(screen, fmt.Sprintf("Now Playing: %s", s.disc.Name),
			fonts.GamerFontS, 12, view.ScreenHeight-10, color.White)
		screen.DrawImage(s.pauseBg, s.pauseBgOp)
		text.Draw(screen, messages.GameOver, fonts.GamerFontL, view.ScreenWidth/2-420, view.ScreenHeight/2, color.White)
		s.startBtn.Draw(screen)
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
	// start music when game state is changed from 'wait'.
}

// StopMusic stops playing music
func (s *Stage01Scene) StopMusic() error {
	return s.disc.Stop()
}
