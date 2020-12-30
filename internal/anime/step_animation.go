package anime

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// StepAnimation is an animation.
// This animates according to the number of steps.
type StepAnimation struct {
	frames          []*ebiten.Image
	durationSteps   int
	maxFrameNum     int
	currentFrameNum int
	walkedSteps     int
}

// NewStepAnimation returns a new StepAnimation generated with args.
func NewStepAnimation(frames []*ebiten.Image, duration int) *StepAnimation {
	return &StepAnimation{
		frames:        frames,
		durationSteps: duration,
		maxFrameNum:   len(frames),
	}
}

// AddStep adds steps information. If your character moved, please
// call this function with steps information.
func (s *StepAnimation) AddStep(steps int) {
	s.walkedSteps += steps
}

// GetCurrentFrame returns a current frame image. This function determines
// the current frame based on the information on how far a character moved.
// If the sum of steps is grater than the DurationSteps, this function will
// return the next frame.
func (s *StepAnimation) GetCurrentFrame() *ebiten.Image {
	if s.walkedSteps > s.durationSteps {
		s.currentFrameNum++
		s.walkedSteps = 0
	}
	if s.currentFrameNum < 0 || s.maxFrameNum-1 < s.currentFrameNum {
		s.currentFrameNum = 0
	}
	return s.frames[s.currentFrameNum]
}
