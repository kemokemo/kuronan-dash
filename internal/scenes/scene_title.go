// Copy from github.com/hajimehoshi/ebiten/v2/example/blocks

package scenes

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/kemokemo/kuronan-dash/assets/fonts"
	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/assets/music"
	"github.com/kemokemo/kuronan-dash/internal/input"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// TitleScene is the scene for title.
type TitleScene struct {
	bg       *ebiten.Image
	disc     *music.Disc
	verPos   view.Vector
	titlePos view.Vector
	msgPos   view.Vector
}

// Initialize initializes all resources.
func (s *TitleScene) Initialize() error {
	s.bg = images.TitleBackground
	s.disc = music.Title
	s.verPos = view.Vector{X: 10, Y: float64(view.ScreenHeight) - 15}
	s.titlePos = view.Vector{
		X: float64(view.ScreenWidth/2) - 200,
		Y: 80}
	s.msgPos = view.Vector{
		X: float64(view.ScreenWidth/2) - 170,
		Y: float64(view.ScreenHeight/2) - 48}
	return nil
}

// Update updates the status of this scene.
func (s *TitleScene) Update(state *GameState) error {
	if input.TriggeredOne() {
		state.SceneManager.GoTo(&SelectScene{})
		return nil
	}
	return nil
}

// Draw draws background and characters.
func (s *TitleScene) Draw(r *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	r.DrawImage(s.bg, op)
	text.Draw(r, versionInfo, fonts.GamerFontS, int(s.verPos.X), int(s.verPos.Y), color.White)
	text.Draw(r, "くろなんダッシュ", fonts.GamerFontLL, int(s.titlePos.X), int(s.titlePos.Y), color.Black)
	text.Draw(r, "Space をおして はじめよう!", fonts.GamerFontM, int(s.msgPos.X), int(s.msgPos.Y), color.Black)
}

// StartMusic starts playing music
func (s *TitleScene) StartMusic() {
	s.disc.Play()
}

// StopMusic stops playing music
func (s *TitleScene) StopMusic() error {
	return s.disc.Stop()
}
