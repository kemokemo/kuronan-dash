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
	"github.com/kemokemo/kuronan-dash/internal/character"
	"github.com/kemokemo/kuronan-dash/internal/music"
	"github.com/kemokemo/kuronan-dash/internal/util"
)

var titleBG *ebiten.Image

func init() {
	var err error
	img, _, err := image.Decode(bytes.NewReader(images.Title_bg_png))
	if err != nil {
		log.Printf("Failed to load the 'Title_bg_png':%v", err)
	}
	titleBG, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Printf("Failed to create a new image from 'Title_bg_png':%v", err)
		return
	}
}

// TitleScene is the scene for title.
type TitleScene struct {
	jb *music.JukeBox
	cm *character.CharacterManager
}

// SetResources sets the resources like music, character images and so on.
func (s *TitleScene) SetResources(j *music.JukeBox, cm *character.CharacterManager) {
	s.jb = j
	s.cm = cm
	err := s.jb.SelectDisc(music.Title)
	if err != nil {
		log.Printf("Failed to select disc:%v", err)
	}
}

// Update updates the status of this scene.
func (s *TitleScene) Update(state *GameState) error {
	if state.Input.StateForKey(ebiten.KeySpace) == 1 {
		state.SceneManager.GoTo(NewSelectScene())
		return nil
	}
	if util.AnyGamepadAbstractButtonPressed(state.Input) {
		state.SceneManager.GoTo(NewSelectScene())
		return nil
	}
	return nil
}

// Draw draws background and characters. This function play music too.
func (s *TitleScene) Draw(r *ebiten.Image) {
	err := s.jb.Play()
	if err != nil {
		log.Printf("Failed to play with JukeBox:%v", err)
		return
	}

	op := &ebiten.DrawImageOptions{}
	r.DrawImage(titleBG, op)
	text.Draw(r, "黒菜んダッシュ", mplus.Gothic12r, 10, 32, color.Black)
	text.Draw(r, "Spaceを押して始めよう!", mplus.Gothic12r, 10, ScreenHeight-48, color.Black)
}
