package kuronandash

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Animation describes an animation
type Animation struct {
	ImagesPaths  []string
	DurationStep int
	DurationTime float64

	frames          []*ebiten.Image
	maxFrameNum     int
	currentFrameNum int

	deltaTotalSteps int
}

// Init loads asset images.
func (a *Animation) Init() error {
	if a.ImagesPaths == nil || len(a.ImagesPaths) == 0 {
		err := fmt.Errorf("paths is empty, please set valid path info of images")
		return err
	}
	a.frames = []*ebiten.Image{}
	for _, path := range a.ImagesPaths {
		image, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterNearest)
		if err != nil {
			return err
		}
		a.frames = append(a.frames, image)
	}
	a.maxFrameNum = len(a.ImagesPaths)
	return nil
}

// GetCurrentFrame returns a current frame image. This function determines
// the current frame based on the information on how far a character moved.
// The deltaStepCount is the delta value of the step count number.
// If deltaStepCount is grater than the DurationStep, this function will
// return the next frame.
func (a *Animation) GetCurrentFrame(deltaStepCount int) *ebiten.Image {
	a.deltaTotalSteps += deltaStepCount
	if a.deltaTotalSteps > a.DurationStep {
		a.currentFrameNum++
		a.deltaTotalSteps = 0
	}
	if a.currentFrameNum < 0 || a.maxFrameNum-1 < a.currentFrameNum {
		a.currentFrameNum = 0
	}
	return a.frames[a.currentFrameNum]
}
