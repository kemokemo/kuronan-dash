// Copy from github.com/hajimehoshi/ebiten/v2/example/blocks

package scenes

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	gauge "github.com/kemokemo/ebiten-gauge"
	progress "github.com/kemokemo/ebiten-progress"

	"github.com/kemokemo/kuronan-dash/assets"
	"github.com/kemokemo/kuronan-dash/assets/fonts"
	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/assets/messages"
	"github.com/kemokemo/kuronan-dash/assets/music"
	"github.com/kemokemo/kuronan-dash/assets/se"

	vpad "github.com/kemokemo/ebiten-virtualpad"
	chara "github.com/kemokemo/kuronan-dash/internal/character"
	"github.com/kemokemo/kuronan-dash/internal/field"
	"github.com/kemokemo/kuronan-dash/internal/input"
	"github.com/kemokemo/kuronan-dash/internal/move"
	"github.com/kemokemo/kuronan-dash/internal/ui"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// Stage01Scene is the scene for the 1st stage game.
type Stage01Scene struct {
	state                   gameState
	player                  *chara.Player
	disc                    *music.Disc
	clickSe                 *se.Player
	readyVoice              *se.Player
	goVoice                 *se.Player
	stageClearVoice         *se.Player
	field                   field.Field
	goalX                   float64
	timeLimit               int // second
	time                    int // second
	sumTicks                float64
	msgWindow               *ui.MessageWindow
	msgWindowTopY           int
	staminaGauge            *gauge.Gauge
	tensionGauge            *gauge.Gauge
	progMap                 *progress.Progress
	progPercent             int
	progMapBk               *ebiten.Image
	opMapBk                 *ebiten.DrawImageOptions
	iChecker                input.InputChecker
	startBtn                vpad.TriggerButton
	pauseBtn                vpad.TriggerButton
	upBtn                   vpad.TriggerButton
	downBtn                 vpad.TriggerButton
	atkBtn                  vpad.TriggerButton
	spBtn                   vpad.TriggerButton
	pauseBg                 *ebiten.Image
	pauseBgOp               *ebiten.DrawImageOptions
	curtain                 *Curtain
	isStarting              bool
	isClosing               bool
	resultEffects           *ResultEffect
	gameSoundControlCh      <-chan assets.GameSoundControl
	gameSoundCancellationCh chan struct{}
}

// Initialize initializes all resources.
func (s *Stage01Scene) Initialize(gameSoundControlCh <-chan assets.GameSoundControl) error {
	s.gameSoundControlCh = gameSoundControlCh
	s.gameSoundCancellationCh = make(chan struct{})

	s.goalX = 4100.0
	s.timeLimit = 40
	s.time = s.timeLimit
	s.disc = music.Stage01
	s.clickSe = se.MenuSelect
	s.readyVoice = se.ReadyVoice
	s.goVoice = se.GoVoice
	s.stageClearVoice = se.StageClearVoice
	s.resultEffects = &ResultEffect{}
	s.resultEffects.Initialize()

	s.player = chara.Selected
	lanes := field.NewLanes(field.PrairieLane)
	err := s.player.InitializeWithLanes(lanes)
	if err != nil {
		return err
	}

	s.field = &field.PrairieField{}
	s.field.Initialize(lanes, s.goalX)

	heights := lanes.GetLaneHeights()
	s.msgWindowTopY = int(heights[len(heights)-1]) + windowMargin
	s.msgWindow = ui.NewMessageWindow(
		300,
		s.msgWindowTopY+5,
		680,
		view.ScreenHeight-windowMargin*2-(s.msgWindowTopY-windowMargin),
		frameWidth)
	s.msgWindow.SetColors(
		color.RGBA{64, 64, 64, 255},
		color.RGBA{192, 192, 192, 255},
		color.RGBA{33, 228, 68, 255})
	s.staminaGauge = gauge.NewGaugeWithScale(440, s.msgWindowTopY+57, s.player.GetMaxStamina(), color.RGBA{255, 255, 255, 255}, 2.0)
	s.staminaGauge.SetBlink(false)
	s.tensionGauge = gauge.NewGaugeWithScale(760, s.msgWindowTopY+57, s.player.GetMaxTension(), color.RGBA{248, 169, 0, 255}, 2.0)
	s.progMap = progress.NewProgress(s.player.MapIcon, 400-16, float64(s.msgWindowTopY+24), view.ScreenWidth-800+16)
	s.progPercent = 0
	s.progMapBk = images.MapBackground
	opMapBk := &ebiten.DrawImageOptions{}
	opMapBk.GeoM.Translate(400, float64(s.msgWindowTopY+21))
	s.opMapBk = opMapBk

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
	bW := images.UpButton.Bounds().Dx()
	bH := images.UpButton.Bounds().Dy()
	s.upBtn.SetLocation(20, view.ScreenHeight-bH-45)
	s.upBtn.SetTriggerButton([]ebiten.Key{ebiten.KeyArrowUp})
	s.downBtn = vpad.NewTriggerButton(images.DownButton, vpad.JustPressed, vpad.SelectColor)
	s.downBtn.SetLocation(20+bW+10, view.ScreenHeight-bH-45)
	s.downBtn.SetTriggerButton([]ebiten.Key{ebiten.KeyArrowDown})
	s.atkBtn = vpad.NewTriggerButton(images.AttackButton, vpad.JustPressed, vpad.SelectColor)
	s.atkBtn.SetLocation(view.ScreenWidth-20-2*bW-10, view.ScreenHeight-bH-45)
	s.atkBtn.SetTriggerButton([]ebiten.Key{ebiten.KeyA})
	s.spBtn = vpad.NewTriggerButton(images.SkillButton, vpad.JustPressed, vpad.SelectColor)
	s.spBtn.SetLocation(view.ScreenWidth-bW-20, view.ScreenHeight-bH-45)
	s.spBtn.SetTriggerButton([]ebiten.Key{ebiten.KeyS})
	s.player.SetInputChecker(laneRectArray, s.upBtn, s.downBtn, s.atkBtn, s.spBtn)

	s.startBtn = vpad.NewTriggerButton(images.StartButton, vpad.JustReleased, vpad.SelectColor)
	s.startBtn.SetLocation(view.ScreenWidth/2-64, view.ScreenHeight/2-128)
	s.startBtn.SetTriggerButton([]ebiten.Key{ebiten.KeySpace})
	s.pauseBtn = vpad.NewTriggerButton(images.PauseButton, vpad.JustReleased, vpad.SelectColor)
	s.pauseBtn.SetLocation(view.ScreenWidth-98, 10)
	s.pauseBtn.SetTriggerButton([]ebiten.Key{ebiten.KeySpace})
	s.iChecker = &input.GameInputChecker{StartBtn: s.startBtn, PauseBtn: s.pauseBtn}

	s.pauseBg = images.PauseLayer
	s.pauseBgOp = &ebiten.DrawImageOptions{}

	s.curtain = NewCurtain()
	s.isStarting = false
	s.isClosing = false

	return nil
}

// Update updates the status of this scene and play sounds.
func (s *Stage01Scene) Update(state *GameState) {
	if s.isStarting || s.isClosing {
		s.curtain.Update()

		if s.curtain.IsFinished() {
			if s.isClosing && s.state == stageClear {
				// goto next stage :-)
				err := state.SceneManager.GoTo(&TitleScene{})
				if err != nil {
					log.Println("failed to go to the 2nd stage: ", err)
				}
			} else if s.isClosing && s.state == gameOver {
				err := state.SceneManager.GoTo(&TitleScene{})
				if err != nil {
					log.Println("failed to go to the title screen: ", err)
				}
			} else if s.isStarting {
				s.isStarting = false
			}
		}
		return
	}

	// s.upBtnとs.downBtnは、s.iChecker内でUpdate()されるのでここではしない
	s.iChecker.Update()

	switch s.state {
	case wait:
		if s.iChecker.TriggeredStart() {
			s.clickSe.Play()
			s.state = readyCall
			s.readyVoice.Play()
			s.sumTicks = 0
		}
	case readyCall:
		if !s.readyVoice.IsPlaying() && s.isFullTicks(0.8) {
			s.state = goCall
			s.goVoice.Play()
			s.sumTicks = 0
		}
	case goCall:
		if !s.goVoice.IsPlaying() && s.isFullTicks(0.5) {
			s.state = run
			s.player.Start()
			s.disc.Play()
		}
	case run:
		if s.iChecker.TriggeredPause() {
			s.state = pause
			s.player.Pause()
			s.disc.Pause()
		} else {
			s.run()
		}
	case skillEffect:
		pState := s.player.Update()
		if pState != move.SkillEffect {
			s.state = run
		}
	case pause:
		if s.iChecker.TriggeredStart() {
			s.state = run
			s.player.ReStart()
			s.disc.Play()
		}
	case stageClear:
		s.resultEffects.Update()
		if s.iChecker.TriggeredStart() {
			s.clickSe.Play()
			s.isClosing = true
			s.curtain.Start(true)
		}
	case gameOver:
		if s.iChecker.TriggeredStart() {
			s.clickSe.Play()
			s.isClosing = true
			s.curtain.Start(true)
		}
	default:
		log.Println("unknown state of Stage01Scene:", s.state)
	}
}

func (s *Stage01Scene) isFullTicks(num float64) bool {
	s.sumTicks += ebiten.ActualTPS()
	if s.sumTicks >= 3600*num {
		s.sumTicks = 0.0
		return true
	}
	return false
}

// run works with 'run' state.
func (s *Stage01Scene) run() {
	if s.isFullTicks(1) {
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
		pState := s.player.Update()
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
		s.progPercent = int((s.player.GetPosition().X - view.DrawPosition) * 100 / s.goalX)
		s.progMap.SetPercent(s.progPercent)

		if pState == move.SkillEffect {
			s.state = skillEffect
		}
	}
}

// Draw draws background and characters.
func (s *Stage01Scene) Draw(screen *ebiten.Image) {
	s.field.DrawFarther(screen)
	s.player.Draw(screen)
	s.field.DrawCloser(screen)
	s.drawUI(screen)
	s.drawWithState(screen)

	tOp := &text.DrawOptions{}
	tOp.GeoM.Translate(10, view.ScreenHeight-15)
	tOp.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, fmt.Sprintf("FPS: %3.1f", ebiten.ActualFPS()), fonts.GamerFontSS, tOp)

	if s.isStarting || s.isClosing {
		s.curtain.Draw(screen)
	} else if s.state == stageClear {
		s.resultEffects.Draw(screen)
	}
}

// description
func (s *Stage01Scene) drawUI(screen *ebiten.Image) {
	s.msgWindow.DrawWindow(screen, "")

	startTextOp := &text.DrawOptions{}
	startTextOp.GeoM.Translate(330, float64(s.msgWindowTopY+20))
	startTextOp.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, "スタート", fonts.GamerFontS, startTextOp)
	screen.DrawImage(s.progMapBk, s.opMapBk)
	s.progMap.Draw(screen)

	goalTextOp := &text.DrawOptions{}
	goalTextOp.GeoM.Translate(910, float64(s.msgWindowTopY+20))
	goalTextOp.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, "ゴール", fonts.GamerFontS, goalTextOp)

	staminaTextOp := &text.DrawOptions{}
	staminaTextOp.GeoM.Translate(350, float64(s.msgWindowTopY+62))
	staminaTextOp.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, "スタミナ:", fonts.GamerFontM, staminaTextOp)
	s.staminaGauge.Draw(screen)

	tensionTextOp := &text.DrawOptions{}
	tensionTextOp.GeoM.Translate(660, float64(s.msgWindowTopY+62))
	tensionTextOp.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, "テンション:", fonts.GamerFontM, tensionTextOp)
	s.tensionGauge.Draw(screen)

	remainTimeTextOp := &text.DrawOptions{}
	remainTimeTextOp.GeoM.Translate(350, float64(s.msgWindowTopY+105))
	remainTimeTextOp.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, fmt.Sprintf("のこりタイム: %v", s.time), fonts.GamerFontM, remainTimeTextOp)

	s.upBtn.Draw(screen)
	s.downBtn.Draw(screen)
	s.atkBtn.Draw(screen)
	s.spBtn.Draw(screen)
}

func (s *Stage01Scene) drawWithState(screen *ebiten.Image) {
	// StartとPauseのボタンは見えてないだけで、該当する場所を押せばボタンはトリガーされる。弊害がありそうなら処置する。
	switch s.state {
	case wait:
		screen.DrawImage(s.pauseBg, s.pauseBgOp)
		tOp := &text.DrawOptions{}
		tOp.GeoM.Translate(view.ScreenWidth/2-300, view.ScreenHeight/2)
		tOp.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, messages.GameStart, fonts.GamerFontL, tOp)
		s.startBtn.Draw(screen)
	case readyCall:
		tOp := &text.DrawOptions{}
		tOp.GeoM.Translate(view.ScreenWidth/2-40, view.ScreenHeight/2-80)
		tOp.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, messages.GameReady, fonts.GamerFontL, tOp)
	case goCall:
		tOp := &text.DrawOptions{}
		tOp.GeoM.Translate(view.ScreenWidth/2-30, view.ScreenHeight/2-80)
		tOp.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, messages.GameGo, fonts.GamerFontL, tOp)
	case pause:
		musicTextOp := &text.DrawOptions{}
		musicTextOp.GeoM.Translate(10, 20)
		musicTextOp.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, fmt.Sprintf("Music: %s", s.disc.Name), fonts.GamerFontS, musicTextOp)

		screen.DrawImage(s.pauseBg, s.pauseBgOp)

		tOp := &text.DrawOptions{}
		tOp.GeoM.Translate(view.ScreenWidth/2-150, view.ScreenHeight/2)
		tOp.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, messages.GamePause, fonts.GamerFontL, tOp)

		s.startBtn.Draw(screen)
	case run:
		s.pauseBtn.Draw(screen)

		tOp := &text.DrawOptions{}
		tOp.GeoM.Translate(10, 20)
		tOp.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, fmt.Sprintf("Music: %s", s.disc.Name), fonts.GamerFontS, tOp)
	case stageClear:
		musicTextOp := &text.DrawOptions{}
		musicTextOp.GeoM.Translate(10, 20)
		musicTextOp.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, fmt.Sprintf("Music: %s", s.disc.Name), fonts.GamerFontS, musicTextOp)

		screen.DrawImage(s.pauseBg, s.pauseBgOp)

		tOp1 := &text.DrawOptions{}
		tOp1.GeoM.Translate(view.ScreenWidth/2-220, view.ScreenHeight/2-220)
		tOp1.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, messages.GameStageClear, fonts.GamerFontL, tOp1)

		tOp2 := &text.DrawOptions{}
		tOp2.GeoM.Translate(view.ScreenWidth/2-260, view.ScreenHeight/2+30)
		tOp2.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, messages.GameStageClear2, fonts.GamerFontL, tOp2)

		tOp3 := &text.DrawOptions{}
		tOp3.GeoM.Translate(view.ScreenWidth/2-180, view.ScreenHeight/2+75)
		tOp3.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, messages.GameStageClear3, fonts.GamerFontL, tOp3)

		s.startBtn.Draw(screen)
	case gameOver:
		musicTextOp := &text.DrawOptions{}
		musicTextOp.GeoM.Translate(10, 20)
		musicTextOp.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, fmt.Sprintf("Music: %s", s.disc.Name), fonts.GamerFontS, musicTextOp)

		screen.DrawImage(s.pauseBg, s.pauseBgOp)

		tOp := &text.DrawOptions{}
		tOp.GeoM.Translate(view.ScreenWidth/2-400, view.ScreenHeight/2+20)
		tOp.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, messages.GameOver, fonts.GamerFontL, tOp)

		s.startBtn.Draw(screen)
	default:
		// nothing to draw
	}
}

func (s *Stage01Scene) Start(gameSoundState bool) {
	s.setVolume(gameSoundState)

	s.isStarting = true
	s.curtain.Start(false)

	go s.playSounds()
}

func (s *Stage01Scene) setVolume(flag bool) {
	s.disc.SetVolumeFlag(flag)
	s.clickSe.SetVolumeFlag(flag)
	s.readyVoice.SetVolumeFlag(flag)
	s.goVoice.SetVolumeFlag(flag)
	s.stageClearVoice.SetVolumeFlag(flag)
	s.player.SetVolumeFlag(flag)
}

func (s *Stage01Scene) playSounds() {
	for {
		select {
		case sControl := <-s.gameSoundControlCh:
			switch sControl {
			case assets.PauseGameSound:
				s.disc.Pause()
			case assets.StartGameSound:
				// このSceneでは最初音楽を再生せずに、ユーザーがStartボタンを押すのを待つ。
				// そのため、gameStateがwaitの時にはSceneManagerからStartGameSound指示が
				// 来ても無視する。
				if s.state != wait {
					s.disc.Play()
				}
			case assets.StopGameSound:
				s.disc.Stop()
			case assets.SoundOn:
				s.setVolume(true)
				if s.state != pause {
					s.disc.Play()
				}
			case assets.SoundOff:
				s.setVolume(false)
			default:
				log.Println("unknown game sound control type, ", s)
			}
		case <-s.gameSoundCancellationCh:
			return
		}
	}
}

func (s *Stage01Scene) Close() error {

	s.disc.Stop()
	close(s.gameSoundCancellationCh)
	return nil
}
