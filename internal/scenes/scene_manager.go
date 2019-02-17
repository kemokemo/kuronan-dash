// Copy from github.com/hajimehoshi/ebiten/example/blocks

package scenes

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/kuronan-dash/internal/character"
	"github.com/kemokemo/kuronan-dash/internal/music"
	"github.com/kemokemo/kuronan-dash/internal/util"
)

const (
	// ScreenWidth is the width of scenes.
	ScreenWidth = 1280
	// ScreenHeight is the heigt of scenes.
	ScreenHeight = 720

	transitionMaxCount = 20
)

// SceneManager manages all scenes.
type SceneManager struct {
	current         Scene
	next            Scene
	transitionFrom  *ebiten.Image
	transitionTo    *ebiten.Image
	op              *ebiten.DrawImageOptions
	transitionCount int
	charaManager    *character.CharacterManager
	jukeBox         *music.JukeBox
}

// GameState describe the state of this game.
type GameState struct {
	SceneManager *SceneManager
	Input        *util.Input
}

// NewSceneManager returns a new SceneManager.
func NewSceneManager() (*SceneManager, error) {
	sm := &SceneManager{}
	var err error
	sm.transitionFrom, err = ebiten.NewImage(ScreenWidth, ScreenHeight, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	sm.transitionTo, err = ebiten.NewImage(ScreenWidth, ScreenHeight, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	sm.op = &ebiten.DrawImageOptions{}
	return sm, nil
}

// SetResources sets the resources like music, character images and so on.
func (s *SceneManager) SetResources(j *music.JukeBox, cm *character.CharacterManager) {
	s.jukeBox = j
	s.charaManager = cm
}

// Update updates the status of this scene.
func (s *SceneManager) Update(input *util.Input) error {
	if s.transitionCount == 0 {
		return s.current.Update(&GameState{
			SceneManager: s,
			Input:        input,
		})
	}
	s.transitionCount--
	if s.transitionCount > 0 {
		return nil
	}
	s.current = s.next
	s.next = nil
	return nil
}

// Draw draws background and characters. This function play music too.
func (s *SceneManager) Draw(r *ebiten.Image) {
	if s.transitionCount == 0 {
		s.current.Draw(r)
		return
	}

	s.transitionFrom.Clear()
	s.current.Draw(s.transitionFrom)
	s.transitionTo.Clear()
	s.next.Draw(s.transitionTo)
	r.DrawImage(s.transitionFrom, nil)

	alpha := 1 - float64(s.transitionCount)/float64(transitionMaxCount)
	s.op.ColorM.Scale(1, 1, 1, alpha)
	r.DrawImage(s.transitionTo, s.op)
}

// GoTo sets resources to the new scene and change the current scene
// to the new scene.
func (s *SceneManager) GoTo(scene Scene) {
	scene.SetResources(s.jukeBox, s.charaManager)
	if s.current == nil {
		s.current = scene
	} else {
		s.next = scene
		s.transitionCount = transitionMaxCount
	}
}
