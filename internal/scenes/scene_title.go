// Copy from github.com/hajimehoshi/ebiten/v2/example/blocks

package scenes

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	vpad "github.com/kemokemo/ebiten-virtualpad"
	"github.com/kemokemo/kuronan-dash/assets/fonts"
	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/assets/music"
	"github.com/kemokemo/kuronan-dash/assets/se"
	"github.com/kemokemo/kuronan-dash/internal/input"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// TitleScene is the scene for title.
type TitleScene struct {
	bg            *ebiten.Image
	disc          *music.Disc
	clickSe       *se.Player
	titleCall     *se.Player
	verPos        view.Vector
	titlePos      view.Vector
	msgPos        view.Vector
	iChecker      input.InputChecker
	vChecker      input.VolumeChecker
	startTitleBtn vpad.TriggerButton
	volumeBtn     vpad.SelectButton
	curtain       *Curtain
	isStarting    bool
	isClosing     bool
}

// Initialize initializes all resources.
func (s *TitleScene) Initialize() error {
	s.bg = images.TitleBackground
	s.disc = music.Title
	s.clickSe = se.MenuSelect
	s.titleCall = se.TitleCall
	s.verPos = view.Vector{X: 10, Y: float64(view.ScreenHeight) - 15}
	s.titlePos = view.Vector{
		X: float64(view.ScreenWidth/2) - 200,
		Y: 80}
	s.msgPos = view.Vector{
		X: float64(view.ScreenWidth/2) - 200,
		Y: float64(view.ScreenHeight/2) + 50}
	s.startTitleBtn = vpad.NewTriggerButton(images.StartTitleButton, vpad.JustReleased, vpad.SelectColor)
	s.startTitleBtn.SetLocation(view.ScreenWidth/2-64, view.ScreenHeight/2-30)
	s.startTitleBtn.SetTriggerButton([]ebiten.Key{ebiten.KeySpace})
	s.volumeBtn = vpad.NewSelectButton(images.VolumeOffButton, vpad.JustPressed, vpad.SelectColor)
	s.volumeBtn.SetLocation(view.ScreenWidth-58, 10)
	s.volumeBtn.SetSelectImage(images.VolumeOnButton)
	s.volumeBtn.SetSelectKeys([]ebiten.Key{ebiten.KeyV})
	s.iChecker = &input.TitleInputChecker{StartBtn: s.startTitleBtn}
	s.vChecker = &input.VolumeInputChecker{VolumeBtn: s.volumeBtn}

	s.curtain = NewCurtain()
	s.isStarting = false
	s.isClosing = false

	return nil
}

// Update updates the status of this scene.
func (s *TitleScene) Update(state *GameState) {
	if s.isStarting || s.isClosing {
		s.curtain.Update()

		if s.curtain.IsFinished() {
			if s.isClosing {
				err := state.SceneManager.GoTo(&SelectScene{})
				if err != nil {
					log.Println("failed to go to the select scene: ", err)
				}
			} else if s.isStarting {
				s.isStarting = false
			}
		}

		return
	}

	s.updateVolume()

	if !s.titleCall.IsPlaying() {
		s.disc.SetVolume(0.5)
	}

	s.iChecker.Update()
	if s.iChecker.TriggeredStart() {
		s.isClosing = true
		s.curtain.Start(true)
		s.clickSe.Play()
		return
	}
}

func (s *TitleScene) updateVolume() {
	s.vChecker.Update()

	if s.vChecker.JustVolumeOn() {
		s.setVolume(true)
		s.disc.Play()
	} else if s.vChecker.JustVolumeOff() {
		s.setVolume(false)
	}
}

func (s *TitleScene) setVolume(flag bool) {
	s.disc.SetVolumeFlag(flag)
	s.clickSe.SetVolumeFlag(flag)
	s.titleCall.SetVolumeFlag(flag)
}

// Draw draws background and characters.
func (s *TitleScene) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(s.bg, op)
	text.Draw(screen, fmt.Sprintf("FPS: %3.1f", ebiten.ActualFPS()), fonts.GamerFontSS, 10, view.ScreenHeight-15, color.White)
	text.Draw(screen, versionInfo, fonts.GamerFontSS, view.ScreenWidth-180, view.ScreenHeight-15, color.White)
	s.startTitleBtn.Draw(screen)
	s.volumeBtn.Draw(screen)

	if s.isStarting || s.isClosing {
		s.curtain.Draw(screen)
	}
}

// StartMusic starts playing music
func (s *TitleScene) StartMusic(isVolumeOn bool) {
	s.volumeBtn.SetSelectState(isVolumeOn)
	if isVolumeOn {
		s.disc.SetVolume(0.3)
		s.disc.Play()
		s.titleCall.Play()
	}
	s.isStarting = true
	s.curtain.Start(false)
}

// StopMusic stops playing music and sound effects
func (s *TitleScene) StopMusic() error {
	var err, e error
	e = s.titleCall.Close()
	if e != nil {
		err = fmt.Errorf("%v, %v", err, e)
	}
	e = s.disc.Stop()
	if e != nil {
		err = fmt.Errorf("%v, %v", err, e)
	}

	return err
}

func (s *TitleScene) IsVolumeOn() bool {
	return s.vChecker.IsVolumeOn()
}
