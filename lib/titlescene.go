// Copy from github.com/hajimehoshi/ebiten/example/blocks

package kuronandash

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	mplusbitmap "github.com/hajimehoshi/go-mplusbitmap"
)

var imageBackground *ebiten.Image

func init() {
	var err error
	imageBackground, _, err = ebitenutil.NewImageFromFile("assets/images/title/background.png", ebiten.FilterNearest)
	if err != nil {
		log.Printf("Failed to load the background image:%v", err)
		return
	}
}

// TitleScene is the scene for title.
type TitleScene struct {
	jukeBox *JukeBox
}

func anyGamepadAbstractButtonPressed(i *Input) bool {
	for _, b := range virtualGamepadButtons {
		if i.gamepadConfig.IsButtonPressed(b) {
			return true
		}
	}
	return false
}

// SetResources sets the resources like music, character images and so on.
func (s *TitleScene) SetResources(j *JukeBox, c *Character) {
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
	if anyGamepadAbstractButtonPressed(state.Input) {
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
	text.Draw(r, "黒菜んダッシュ", mplusbitmap.Gothic12r, 10, 32, color.Black)
	text.Draw(r, "Spaceを押して始めよう!", mplusbitmap.Gothic12r, 10, ScreenHeight-48, color.Black)
}

func (s *TitleScene) drawTitleBackground(r *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	r.DrawImage(imageBackground, op)
}
