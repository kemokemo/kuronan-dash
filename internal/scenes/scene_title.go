// Copy from github.com/hajimehoshi/ebiten/v2/example/blocks

package scenes

import (
	"fmt"
	"image/color"
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	vpad "github.com/kemokemo/ebiten-virtualpad"
	"github.com/kemokemo/kuronan-dash/assets"
	"github.com/kemokemo/kuronan-dash/assets/fonts"
	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/assets/music"
	"github.com/kemokemo/kuronan-dash/assets/se"
	"github.com/kemokemo/kuronan-dash/internal/input"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// TitleScene is the scene for title.
type TitleScene struct {
	bg                      *ebiten.Image
	disc                    *music.Disc
	clickSe                 *se.Player
	titleCall               *se.Player
	verPos                  view.Vector
	titlePos                view.Vector
	msgPos                  view.Vector
	iChecker                input.InputChecker
	startTitleBtn           vpad.TriggerButton
	curtain                 *Curtain
	isStarting              bool
	isClosing               bool
	once                    sync.Once
	secondOrLater           bool
	gameSoundControlCh      <-chan assets.GameSoundControl
	gameSoundCancellationCh chan struct{}
}

// Initialize initializes all resources.
func (s *TitleScene) Initialize(gameSoundControlCh <-chan assets.GameSoundControl) error {
	s.gameSoundControlCh = gameSoundControlCh
	s.gameSoundCancellationCh = make(chan struct{})

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
	s.startTitleBtn.SetLocation(view.ScreenWidth/2-220, view.ScreenHeight/2+20)
	s.startTitleBtn.SetTriggerButton([]ebiten.Key{ebiten.KeySpace})
	s.iChecker = &input.TitleInputChecker{StartBtn: s.startTitleBtn}

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

// Draw draws background and characters.
func (s *TitleScene) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(s.bg, op)

	fpsTextOp := &text.DrawOptions{}
	fpsTextOp.GeoM.Translate(10, view.ScreenHeight-15)
	fpsTextOp.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, fmt.Sprintf("FPS: %3.1f", ebiten.ActualFPS()), fonts.GamerFontSS, fpsTextOp)

	versionTextOp := &text.DrawOptions{}
	versionTextOp.GeoM.Translate(view.ScreenWidth-180, view.ScreenHeight-15)
	versionTextOp.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, versionInfo, fonts.GamerFontSS, versionTextOp)

	s.startTitleBtn.Draw(screen)

	if s.isStarting || s.isClosing {
		s.curtain.Draw(screen)
	}
}

func (s *TitleScene) Start(gameSoundState bool) {
	s.setVolume(gameSoundState)

	if s.secondOrLater {
		s.isStarting = true
		s.curtain.Start(false)
	} else {
		s.isStarting = false
	}

	s.once.Do(func() {
		s.secondOrLater = true
	})

	go s.playSounds()
}

func (s *TitleScene) setVolume(flag bool) {
	s.disc.SetVolumeFlag(flag)
	s.clickSe.SetVolumeFlag(flag)
	s.titleCall.SetVolumeFlag(flag)
}

func (s *TitleScene) playSounds() {
	for {
		select {
		case sControl := <-s.gameSoundControlCh:
			switch sControl {
			case assets.PauseGameSound:
				s.disc.Pause()
			case assets.StartGameSound:
				s.disc.SetVolume(0.3)
				s.disc.Play()
				s.titleCall.Play()
			case assets.StopGameSound:
				s.disc.Stop()
			case assets.SoundOn:
				s.setVolume(true)
				// discだけはBGMなので、ここでPlayも実行して再生始める
				s.disc.Play()
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

func (s *TitleScene) Close() error {
	// SEなどassetsは、SceneManagerでまとめてCloseするのでここでは閉じない。

	s.disc.Stop()

	// playSounds ゴルーチンを止める
	close(s.gameSoundCancellationCh)
	return nil
}
