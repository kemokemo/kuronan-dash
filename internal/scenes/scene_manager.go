// Copy from github.com/hajimehoshi/ebiten/v2/example/blocks

package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	vpad "github.com/kemokemo/ebiten-virtualpad"
	"github.com/kemokemo/kuronan-dash/assets"
	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/internal/input"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

const transitionMaxCount = 20

var versionInfo = ""

// SceneManager manages all scenes.
type SceneManager struct {
	current            Scene
	next               Scene
	transitionFrom     *ebiten.Image
	transitionTo       *ebiten.Image
	op                 *ebiten.DrawImageOptions
	transitionCount    int
	alpha              float32
	volumeBtn          vpad.SelectButton
	gameSoundControlCh chan assets.GameSoundControl
	vChecker           input.VolumeChecker
}

// NewSceneManager returns a new SceneManager.
func NewSceneManager(ver string) *SceneManager {
	versionInfo = ver

	volumeBtn := vpad.NewSelectButton(images.VolumeOffButton, vpad.JustPressed, vpad.SelectColor)
	volumeBtn.SetLocation(view.ScreenWidth-58, 10)
	volumeBtn.SetSelectImage(images.VolumeOnButton)
	volumeBtn.SetSelectKeys([]ebiten.Key{ebiten.KeyV})
	gameSoundControlCh := make(chan assets.GameSoundControl)
	vChecker := input.NewVolumeChecker(volumeBtn, true, gameSoundControlCh)

	sm := &SceneManager{
		transitionFrom:     ebiten.NewImage(view.ScreenWidth, view.ScreenHeight),
		transitionTo:       ebiten.NewImage(view.ScreenWidth, view.ScreenHeight),
		op:                 &ebiten.DrawImageOptions{},
		volumeBtn:          volumeBtn,
		gameSoundControlCh: gameSoundControlCh,
		vChecker:           vChecker,
	}

	return sm
}

// Update updates the status of this scene.
func (s *SceneManager) Update() {
	// 現在のSceneに遷移完了したので、そちらのUpdateを実行。
	if s.transitionCount == 0 {
		s.current.Update(&GameState{
			SceneManager: s,
		})
		s.vChecker.Update()
		return
	}

	// Sceneの遷移中
	s.transitionCount--
	if s.transitionCount > 0 {
		return
	}

	// Sceneの遷移が完了
	s.current = s.next
	s.next = nil
	s.gameSoundControlCh <- assets.StartGameSound
}

// Draw draws background and characters. This function play music too.
func (s *SceneManager) Draw(r *ebiten.Image) {
	if s.transitionCount == 0 {
		s.current.Draw(r)
		s.volumeBtn.Draw(r)
		return
	}

	s.transitionFrom.Clear()
	s.current.Draw(s.transitionFrom)
	s.transitionTo.Clear()
	s.next.Draw(s.transitionTo)
	r.DrawImage(s.transitionFrom, nil)

	s.alpha = 1 - float32(s.transitionCount)/float32(transitionMaxCount)
	s.op.ColorScale.Scale(s.alpha, s.alpha, s.alpha, s.alpha)
	r.DrawImage(s.transitionTo, s.op)
}

// GoTo sets resources to the new scene and change the current scene
// to the new scene. This stops the music of the current and starts
// the music of the next.
func (s *SceneManager) GoTo(scene Scene) error {
	err := scene.Initialize(s.gameSoundControlCh)
	if err != nil {
		return err
	}

	// 最初のScene
	if s.current == nil {
		s.current = scene
		s.current.Start(s.vChecker.IsVolumeOn())
		s.gameSoundControlCh <- assets.StartGameSound
		return nil
	}

	// Sceneの遷移
	s.gameSoundControlCh <- assets.StopGameSound
	err = s.current.Close()
	if err != nil {
		return err
	}
	s.next = scene
	s.next.Start(s.vChecker.IsVolumeOn())
	s.transitionCount = transitionMaxCount

	return nil
}

func (s *SceneManager) Close() {
	close(s.gameSoundControlCh)
}
