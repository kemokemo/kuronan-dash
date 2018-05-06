package scenes

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	mplus "github.com/hajimehoshi/go-mplusbitmap"
	"github.com/kemokemo/kuronan-dash/lib/music"
	"github.com/kemokemo/kuronan-dash/lib/objects"
	"github.com/kemokemo/kuronan-dash/lib/util"
)

// SelectScene is the scene to select the player character.
type SelectScene struct {
	jb *music.JukeBox
	cm *objects.CharacterManager
}

// NewSelectScene creates the new GameScene.
func NewSelectScene() *SelectScene {
	return &SelectScene{}
}

// SetResources sets the resources like music, character images and so on.
func (s *SelectScene) SetResources(j *music.JukeBox, cm *objects.CharacterManager) {
	s.jb = j
	s.cm = cm
	err := s.jb.SelectDisc(music.Title)
	if err != nil {
		log.Printf("Failed to select disc:%v", err)
	}
}

// Update updates the status of this scene.
func (s *SelectScene) Update(state *GameState) error {
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
func (s *SelectScene) Draw(r *ebiten.Image) {
	err := s.jb.Play()
	if err != nil {
		log.Printf("Failed to play with JukeBox:%v", err)
		return
	}

	text.Draw(r, "← → のカーソルキーでキャラクターを選んでSpaceキーを押してね！", mplus.Gothic12r, 10, 32, color.White)
}
