// Copy from github.com/hajimehoshi/ebiten/v2/example/blocks

package scenes

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

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
	s.iChecker = &input.TitleInputChecker{}
	return nil
}

// Update updates the status of this scene.
func (s *TitleScene) Update(state *GameState) {
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

// Draw draws background and characters.
func (s *TitleScene) Draw(r *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	r.DrawImage(s.bg, op)
	text.Draw(r, versionInfo, fonts.GamerFontS, int(s.verPos.X), int(s.verPos.Y), color.White)
	text.Draw(r, messages.TitleStart, fonts.GamerFontL, int(s.msgPos.X), int(s.msgPos.Y), color.Black)
}

// StartMusic starts playing music
func (s *TitleScene) StartMusic() {
	s.titleCall.Play()
	s.disc.SetVolume(0.3)
	s.disc.Play()
}

// StopMusic stops playing music
func (s *TitleScene) StopMusic() error {
	s.titleCall.Close()
	return s.disc.Stop()
}
