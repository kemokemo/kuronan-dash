// Copy from github.com/hajimehoshi/ebiten/v2/example/blocks

package scenes

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	gauge "github.com/kemokemo/ebiten-gauge"
	"github.com/kemokemo/kuronan-dash/assets/fonts"
	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/assets/messages"
	"github.com/kemokemo/kuronan-dash/assets/music"
	"github.com/kemokemo/kuronan-dash/assets/se"

	vpad "github.com/kemokemo/ebiten-virtualpad"
	chara "github.com/kemokemo/kuronan-dash/internal/character"
	"github.com/kemokemo/kuronan-dash/internal/field"
	"github.com/kemokemo/kuronan-dash/internal/input"
	"github.com/kemokemo/kuronan-dash/internal/ui"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// Stage01Scene is the scene for the 1st stage game.
type Stage01Scene struct {
	state           gameState
	player          *chara.Player
	disc            *music.Disc
	readyVoice      *se.Player
	goVoice         *se.Player
	stageClearVoice *se.Player
	field           field.Field
	goalX           float64
	timeLimit       int // second
	time            int // second
	sumTicks        float64
	msgWindow       *ui.MessageWindow
	staminaGauge    *gauge.Gauge
	tensionGauge    *gauge.Gauge
	uiMsg           string
	iChecker        input.InputChecker
	vChecker        input.VolumeChecker
	startBtn        vpad.TriggerButton
	pauseBtn        vpad.TriggerButton
	volumeBtn       vpad.SelectButton
	upBtn           vpad.TriggerButton
	downBtn         vpad.TriggerButton
	atkBtn          vpad.TriggerButton
	spBtn           vpad.TriggerButton
	pauseBg         *ebiten.Image
	pauseBgOp       *ebiten.DrawImageOptions
}

// Initialize initializes all resources.
func (s *Stage01Scene) Initialize() error {
	s.goalX = 3900.0
	s.timeLimit = 90
	s.time = s.timeLimit
	s.disc = music.Stage01
	s.readyVoice = se.ReadyVoice
	s.goVoice = se.GoVoice
	s.stageClearVoice = se.StageClearVoice

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
	s.staminaGauge = gauge.NewGaugeWithColor(405, int(lowerLaneY)+windowMargin+15, s.player.GetMaxStamina(), color.RGBA{255, 255, 255, 255})
	s.staminaGauge.SetBlink(false)
	s.tensionGauge = gauge.NewGaugeWithColor(405, int(lowerLaneY)+windowMargin+35, s.player.GetMaxTension(), color.RGBA{248, 169, 0, 255})

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
	s.upBtn.SetTriggerButton([]ebiten.Key{ebiten.KeyArrowUp})
	s.downBtn = vpad.NewTriggerButton(images.DownButton, vpad.JustPressed, vpad.SelectColor)
	s.downBtn.SetLocation(20+bW+10, view.ScreenHeight-bH-45)
	s.downBtn.SetTriggerButton([]ebiten.Key{ebiten.KeyArrowDown})
	s.atkBtn = vpad.NewTriggerButton(images.AttackButton, vpad.JustPressed, vpad.SelectColor)
	s.atkBtn.SetLocation(view.ScreenWidth-20-2*bW-10, view.ScreenHeight-bH-45)
	s.atkBtn.SetTriggerButton([]ebiten.Key{ebiten.KeyA})
	s.spBtn = vpad.NewTriggerButton(images.SpecialButton, vpad.JustPressed, vpad.SelectColor)
	s.spBtn.SetLocation(view.ScreenWidth-bW-20, view.ScreenHeight-bH-45)
	s.spBtn.SetTriggerButton([]ebiten.Key{ebiten.KeyS})
	s.player.SetInputChecker(laneRectArray, s.upBtn, s.downBtn, s.atkBtn, s.spBtn)

	s.startBtn = vpad.NewTriggerButton(images.StartButton, vpad.JustReleased, vpad.SelectColor)
	s.startBtn.SetLocation(view.ScreenWidth/2-64, view.ScreenHeight/2-128)
	s.startBtn.SetTriggerButton([]ebiten.Key{ebiten.KeySpace})
	s.volumeBtn = vpad.NewSelectButton(images.VolumeOffButton, vpad.JustPressed, vpad.SelectColor)
	s.volumeBtn.SetLocation(view.ScreenWidth-58, 10)
	s.volumeBtn.SetSelectImage(images.VolumeOnButton)
	s.volumeBtn.SetSelectKeys([]ebiten.Key{ebiten.KeyV})
	s.pauseBtn = vpad.NewTriggerButton(images.PauseButton, vpad.JustReleased, vpad.SelectColor)
	s.pauseBtn.SetLocation(view.ScreenWidth-98, 10)
	s.pauseBtn.SetTriggerButton([]ebiten.Key{ebiten.KeySpace})
	s.iChecker = &input.GameInputChecker{StartBtn: s.startBtn, PauseBtn: s.pauseBtn}
	s.vChecker = input.NewVolumeChecker(s.volumeBtn, true)

	s.pauseBg = images.PauseLayer
	s.pauseBgOp = &ebiten.DrawImageOptions{}

	return nil
}

// Update updates the status of this scene and play sounds.
func (s *Stage01Scene) Update(state *GameState) {
	s.updateVolume()

	// s.upBtnとs.downBtnは、s.iChecker内でUpdate()されるのでここではしない
	s.iChecker.Update()

	switch s.state {
	case wait:
		if s.iChecker.TriggeredStart() {
			s.state = readyCall
			s.readyVoice.Play()
		}
	case readyCall:
		if !s.readyVoice.IsPlaying() {
			s.state = goCall
			s.goVoice.Play()
		}
	case goCall:
		if !s.goVoice.IsPlaying() {
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

// updateVolume updates the volume on/off state of music and sounds.
// If you add some sounds, please add this logic.
func (s *Stage01Scene) updateVolume() {
	s.vChecker.Update()

	if s.vChecker.JustVolumeOn() {
		s.setVolume(true)
		s.disc.Play()
	} else if s.vChecker.JustVolumeOff() {
		s.setVolume(false)
	}
}

func (s *Stage01Scene) setVolume(flag bool) {
	s.disc.SetVolumeFlag(flag)
	s.readyVoice.SetVolumeFlag(flag)
	s.goVoice.SetVolumeFlag(flag)
	s.stageClearVoice.SetVolumeFlag(flag)
	s.player.SetVolumeFlag(flag)
}

// run works with 'run' state.
func (s *Stage01Scene) run() {
	// todo: do not use CurrentTPS for game (check the comment of func)
	// I will use time.Elapsed for count.
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
		s.stageClearVoice.Play()
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

		s.staminaGauge.Update(float64(s.player.GetStamina()))
		s.tensionGauge.Update(float64(s.player.GetTension()))
	}
}

// Draw draws background and characters.
func (s *Stage01Scene) Draw(screen *ebiten.Image) {
	s.field.DrawFarther(screen)
	s.player.Draw(screen)
	s.field.DrawCloser(screen)
	s.drawUI(screen)
	s.drawWithState(screen)

	// Let's make sure that the volume can be changed at any time.
	s.volumeBtn.Draw(screen)
}

// description
func (s *Stage01Scene) drawUI(screen *ebiten.Image) {
	s.uiMsg = fmt.Sprintf("スタミナ :\nテンション:\nタイム: 　%v\nすすんだきょり: %.1f\nゴールのいち: 　%.1f",
		s.time,
		s.player.GetPosition().X-view.DrawPosition,
		s.goalX,
	)
	s.msgWindow.DrawWindow(screen, s.uiMsg)
	s.staminaGauge.Draw(screen)
	s.tensionGauge.Draw(screen)

	s.upBtn.Draw(screen)
	s.downBtn.Draw(screen)
	s.atkBtn.Draw(screen)
	s.spBtn.Draw(screen)
}

func (s *Stage01Scene) drawWithState(screen *ebiten.Image) {
	// TODO: StartとPauseのボタンは見えてないだけで、該当する場所を押せばボタンはトリガーされる。弊害がありそうなら処置する。
	switch s.state {
	case wait:
		screen.DrawImage(s.pauseBg, s.pauseBgOp)
		text.Draw(screen, messages.GameStart, fonts.GamerFontL, view.ScreenWidth/2-280, view.ScreenHeight/2+30, color.White)
		s.startBtn.Draw(screen)
	case readyCall:
		text.Draw(screen, messages.GameReady, fonts.GamerFontL, view.ScreenWidth/2-30, view.ScreenHeight/2+30, color.White)
	case goCall:
		text.Draw(screen, messages.GameGo, fonts.GamerFontL, view.ScreenWidth/2-20, view.ScreenHeight/2+30, color.White)
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
func (s *Stage01Scene) StartMusic(isVolumeOn bool) {
	s.volumeBtn.SetSelectState(isVolumeOn)

	s.setVolume(isVolumeOn)
	// when the game state is changed to 'run', the music starts. not now.
}

// StopMusic stops playing music and sound effects
func (s *Stage01Scene) StopMusic() error {
	var err, e error
	e = s.stageClearVoice.Close()
	if e != nil {
		err = fmt.Errorf("%v, %v", err, e)
	}
	e = s.disc.Stop()
	if e != nil {
		err = fmt.Errorf("%v, %v", err, e)
	}

	return err
}

func (s *Stage01Scene) IsVolumeOn() bool {
	return s.vChecker.IsVolumeOn()
}
