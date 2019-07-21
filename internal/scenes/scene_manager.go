// Copy from github.com/hajimehoshi/ebiten/example/blocks

package scenes

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/kuronan-dash/internal/input"
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

// Update updates the status of this scene.
func (s *SceneManager) Update(input *input.Input) error {
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
// to the new scene. This stops the music of the current and starts
// the music of the next.
func (s *SceneManager) GoTo(scene Scene) error {
	err := scene.Initialize()
	if err != nil {
		return err
	}

	if s.current == nil {
		s.current = scene
		err = s.current.StartMusic()
		if err != nil {
			return err
		}
	} else {
		err = s.current.StopMusic()
		if err != nil {
			return err
		}
		s.next = scene
		err = s.next.StartMusic()
		if err != nil {
			return err
		}
		s.transitionCount = transitionMaxCount
	}

	return nil
}
