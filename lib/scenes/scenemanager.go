// Copy from github.com/hajimehoshi/ebiten/example/blocks

package scenes

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/kuronan-dash/lib/assetsutil"
	"github.com/kemokemo/kuronan-dash/lib/objects"
)

const (
	// ScreenWidth is the width of scenes.
	ScreenWidth = 800
	// ScreenHeight is the heigt of scenes.
	ScreenHeight = 480
)

var (
	transitionFrom *ebiten.Image
	transitionTo   *ebiten.Image
)

func init() {
	transitionFrom, _ = ebiten.NewImage(ScreenWidth, ScreenHeight, ebiten.FilterDefault)
	transitionTo, _ = ebiten.NewImage(ScreenWidth, ScreenHeight, ebiten.FilterDefault)
}

// Scene is interface for the all scenes.
type Scene interface {
	SetResources(j *assetsutil.JukeBox, c *objects.Character)
	Update(state *GameState) error
	Draw(screen *ebiten.Image)
}

const transitionMaxCount = 20

// SceneManager manages all scenes.
type SceneManager struct {
	current         Scene
	next            Scene
	transitionCount int
	character       *objects.Character
	jukeBox         *assetsutil.JukeBox
}

// GameState describe the state of this game.
type GameState struct {
	SceneManager *SceneManager
	Input        *assetsutil.Input
}

// SetResources sets the resources like music, character images and so on.
func (s *SceneManager) SetResources(j *assetsutil.JukeBox, c *objects.Character) {
	s.jukeBox = j
	s.character = c
}

// Update updates the status of this scene.
func (s *SceneManager) Update(input *assetsutil.Input) error {
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

	transitionFrom.Clear()
	s.current.Draw(transitionFrom)

	transitionTo.Clear()
	s.next.Draw(transitionTo)

	r.DrawImage(transitionFrom, nil)

	alpha := 1 - float64(s.transitionCount)/float64(transitionMaxCount)
	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, alpha)
	r.DrawImage(transitionTo, op)
}

// GoTo sets resources to the new scene and change the current scene
// to the new scene.
func (s *SceneManager) GoTo(scene Scene) {
	scene.SetResources(s.jukeBox, s.character)
	if s.current == nil {
		s.current = scene
	} else {
		s.next = scene
		s.transitionCount = transitionMaxCount
	}
}
