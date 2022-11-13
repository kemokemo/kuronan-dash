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
	"github.com/kemokemo/kuronan-dash/assets/messages"
	"github.com/kemokemo/kuronan-dash/assets/music"
	"github.com/kemokemo/kuronan-dash/assets/se"
	"github.com/kemokemo/kuronan-dash/internal/input"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// TitleScene is the scene for title.
type TitleScene struct {
	bg        *ebiten.Image
	disc      *music.Disc
	titleCall *se.Player
	verPos    view.Vector
	titlePos  view.Vector
	msgPos    view.Vector
	iChecker  input.InputChecker
	vChecker  input.VolumeChecker
	volumeBtn vpad.SelectButton
}

// Initialize initializes all resources.
func (s *TitleScene) Initialize() error {
	s.bg = images.TitleBackground
	s.disc = music.Title
	s.titleCall = se.TitleCall
	s.verPos = view.Vector{X: 10, Y: float64(view.ScreenHeight) - 15}
	s.titlePos = view.Vector{
		X: float64(view.ScreenWidth/2) - 200,
		Y: 80}
	s.msgPos = view.Vector{
		X: float64(view.ScreenWidth/2) - 200,
		Y: float64(view.ScreenHeight/2) + 50}
	s.volumeBtn = vpad.NewSelectButton(images.VolumeOnButton, vpad.JustPressed, vpad.SelectColor)
	s.volumeBtn.SetLocation(view.ScreenWidth-58, 10)
	s.iChecker = &input.TitleInputChecker{}
	s.vChecker = &input.VolumeInputChecker{VolumeBtn: s.volumeBtn}

	return nil
}

// Update updates the status of this scene.
func (s *TitleScene) Update(state *GameState) {
	s.updateVolume()

	if !s.titleCall.IsPlaying() {
		s.disc.SetVolume(0.8)
	}

	s.iChecker.Update()
	if s.iChecker.TriggeredStart() {
		err := state.SceneManager.GoTo(&SelectScene{})
		if err != nil {
			log.Println("failed to go to the select scene: ", err)
		}
	}
}

func (s *TitleScene) updateVolume() {
	s.vChecker.Update()

	if s.vChecker.JustVolumeOn() {
		s.disc.SetVolumeFlag(true)
		s.titleCall.SetVolumeFlag(true)
		s.disc.Play()
	} else if s.vChecker.JustVolumeOff() {
		s.disc.SetVolumeFlag(false)
		s.titleCall.SetVolumeFlag(false)
	}
}

// Draw draws background and characters.
func (s *TitleScene) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(s.bg, op)
	text.Draw(screen, versionInfo, fonts.GamerFontS, int(s.verPos.X), int(s.verPos.Y), color.White)
	text.Draw(screen, messages.TitleStart, fonts.GamerFontL, int(s.msgPos.X), int(s.msgPos.Y), color.Black)
	s.volumeBtn.Draw(screen)
}

// StartMusic starts playing music
func (s *TitleScene) StartMusic(isVolumeOn bool) {
	s.volumeBtn.SetSelectState(isVolumeOn)
	s.updateVolume()
	s.disc.SetVolume(0.3)
	s.disc.Play()
	s.titleCall.Play()
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
