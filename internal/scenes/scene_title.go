// Copy from github.com/hajimehoshi/ebiten/example/blocks

package scenes

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"

	"github.com/kemokemo/kuronan-dash/assets/fonts"
	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/assets/music"
	"github.com/kemokemo/kuronan-dash/internal/input"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// TitleScene is the scene for title.
type TitleScene struct {
	bg       *ebiten.Image
	Version  string
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
	if state.Input.StateForKey(ebiten.KeySpace) == 1 {
		state.SceneManager.GoTo(&SelectScene{})
		return nil
	}
	if input.AnyGamepadAbstractButtonPressed(state.Input) {
		state.SceneManager.GoTo(&SelectScene{})
		return nil
	}
	return nil
}

// Draw draws background and characters.
func (s *TitleScene) Draw(r *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	r.DrawImage(s.bg, op)
	text.Draw(r, s.Version, fonts.GamerFontS, int(s.verPos.X), int(s.verPos.Y), color.White)
	text.Draw(r, "くろなんダッシュ", fonts.GamerFontLL, int(s.titlePos.X), int(s.titlePos.Y), color.Black)
	text.Draw(r, "Space をおして はじめよう!", fonts.GamerFontM, int(s.msgPos.X), int(s.msgPos.Y), color.Black)
	return nil
}

// StartMusic starts playing music
func (s *TitleScene) StartMusic() error {
	return s.disc.Play()
}

// StopMusic stops playing music
func (s *TitleScene) StopMusic() error {
	return s.disc.Stop()
}
