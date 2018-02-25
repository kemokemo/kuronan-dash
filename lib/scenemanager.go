// Copy from github.com/hajimehoshi/ebiten/example/blocks

package kuronandash

import (
	"github.com/hajimehoshi/ebiten"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 480
)

var (
	transitionFrom *ebiten.Image
	transitionTo   *ebiten.Image
)

func init() {
	transitionFrom, _ = ebiten.NewImage(ScreenWidth, ScreenHeight, ebiten.FilterNearest)
	transitionTo, _ = ebiten.NewImage(ScreenWidth, ScreenHeight, ebiten.FilterNearest)
}

type Scene interface {
	SetResources(j *JukeBox, c *Character)
	Update(state *GameState) error
	Draw(screen *ebiten.Image)
}

const transitionMaxCount = 20

type SceneManager struct {
	current         Scene
	next            Scene
	transitionCount int
	character       *Character
	jukeBox         *JukeBox
}

type GameState struct {
	SceneManager *SceneManager
	Input        *Input
}

func (s *SceneManager) SetResources(j *JukeBox, c *Character) {
	s.jukeBox = j
	s.character = c
}

func (s *SceneManager) Update(input *Input) error {
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

func (s *SceneManager) GoTo(scene Scene) {
	scene.SetResources(s.jukeBox, s.character)
	if s.current == nil {
		s.current = scene
	} else {
		s.next = scene
		s.transitionCount = transitionMaxCount
	}
}
