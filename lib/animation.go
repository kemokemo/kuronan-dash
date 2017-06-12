package kuronandash

import (
	"fmt"
	"sync"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// StepAnimation is an animation.
// This animates according to step number information.
type StepAnimation struct {
	ImagesPaths     []string
	DurationSteps   int
	once            sync.Once
	frames          []*ebiten.Image
	maxFrameNum     int
	currentFrameNum int
	totalSteps      int
}

// Init loads asset images and initializes private parameters.
func (a *StepAnimation) Init() (err error) {
	a.once.Do(func() {
		if a.ImagesPaths == nil || len(a.ImagesPaths) == 0 {
			err = fmt.Errorf("paths is empty, please set valid path info of images")
			return
		}
		a.frames = []*ebiten.Image{}
		var image *ebiten.Image
		for _, path := range a.ImagesPaths {
			image, _, err = ebitenutil.NewImageFromFile(path, ebiten.FilterNearest)
			if err != nil {
				return
			}
			a.frames = append(a.frames, image)
		}
		a.maxFrameNum = len(a.ImagesPaths)
	})
	if err != nil {
		return err
	}

	a.currentFrameNum = 0
	a.totalSteps = 0

	return nil
}

// GetCurrentFrame returns a current frame image. This function determines
// the current frame based on the information on how far a character moved.
// As steps variables, please set the number of steps since calling this
// function last time. If the sum of steps is grater than the DurationSteps,
// this function will return the next frame.
func (a *StepAnimation) GetCurrentFrame(steps int) *ebiten.Image {
	a.totalSteps += steps
	if a.totalSteps > a.DurationSteps {
		a.currentFrameNum++
		a.totalSteps = 0
	}
	if a.currentFrameNum < 0 || a.maxFrameNum-1 < a.currentFrameNum {
		a.currentFrameNum = 0
	}
	return a.frames[a.currentFrameNum]
}
