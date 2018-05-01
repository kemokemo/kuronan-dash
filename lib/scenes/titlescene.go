// Copy from github.com/hajimehoshi/ebiten/example/blocks

package scenes

import (
	"bytes"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	mplus "github.com/hajimehoshi/go-mplusbitmap"
	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/lib/music"
	"github.com/kemokemo/kuronan-dash/lib/objects"
	"github.com/kemokemo/kuronan-dash/lib/util"
)

var imageBackground *ebiten.Image

func init() {
	var err error
	img, _, err := image.Decode(bytes.NewReader(images.Title_bg_png))
	if err != nil {
		log.Printf("Failed to load the 'Title_bg_png':%v", err)
	}
	imageBackground, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Printf("Failed to create a new image from 'Title_bg_png':%v", err)
		return
	}
}

// TitleScene is the scene for title.
type TitleScene struct {
	jukeBox *music.JukeBox
}

// SetResources sets the resources like music, character images and so on.
func (s *TitleScene) SetResources(j *music.JukeBox, c *objects.Character) {
	s.jukeBox = j
	err := s.jukeBox.SelectDisc("shibugaki_no_kuroneko")
	if err != nil {
		log.Printf("Failed to select disc:%v", err)
	}
}

// Update updates the status of this scene.
func (s *TitleScene) Update(state *GameState) error {
	if state.Input.StateForKey(ebiten.KeySpace) == 1 {
		state.SceneManager.GoTo(NewGameScene())
		return nil
	}
	if util.AnyGamepadAbstractButtonPressed(state.Input) {
		state.SceneManager.GoTo(NewGameScene())
		return nil
	}
	return nil
}

// Draw draws background and characters. This function play music too.
func (s *TitleScene) Draw(r *ebiten.Image) {
	err := s.jukeBox.Play()
	if err != nil {
		log.Printf("Failed to play with JukeBox:%v", err)
		return
	}

	s.drawTitleBackground(r)
	text.Draw(r, "黒菜んダッシュ", mplus.Gothic12r, 10, 32, color.Black)
	text.Draw(r, "Spaceを押して始めよう!", mplus.Gothic12r, 10, ScreenHeight-48, color.Black)
}

func (s *TitleScene) drawTitleBackground(r *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	r.DrawImage(imageBackground, op)
}
