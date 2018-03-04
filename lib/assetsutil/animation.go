package assetsutil

import (
	"fmt"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// StepAnimation is an animation.
// This animates according to the number of steps.
type StepAnimation struct {
	ImagesPaths     []string
	DurationSteps   int
	once            sync.Once
	frames          []*ebiten.Image
	maxFrameNum     int
	currentFrameNum int
	walkedSteps     int
}

// Init loads images and initializes private parameters.
// If you call this function multiple times, it is only the
// first time to load the images.
func (s *StepAnimation) Init() (err error) {
	s.once.Do(func() {
		s.frames, err = loadImages(s.ImagesPaths)
		s.maxFrameNum = len(s.ImagesPaths)
	})
	if err != nil {
		return err
	}
	s.currentFrameNum = 0
	s.walkedSteps = 0
	return nil
}

func loadImages(paths []string) ([]*ebiten.Image, error) {
	if paths == nil || len(paths) == 0 {
		err := fmt.Errorf("paths is empty, please set valid path info of images")
		return nil, err
	}
	frames := []*ebiten.Image{}
	for _, path := range paths {
		image, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterNearest)
		if err != nil {
			return nil, err
		}
		frames = append(frames, image)
	}
	return frames, nil
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
	if s.walkedSteps > s.DurationSteps {
		s.currentFrameNum++
		s.walkedSteps = 0
	}
	if s.currentFrameNum < 0 || s.maxFrameNum-1 < s.currentFrameNum {
		s.currentFrameNum = 0
	}
	return s.frames[s.currentFrameNum]
}

// TimeAnimation is an animation.
// This animates according to elapsed time.
type TimeAnimation struct {
	ImagesPaths     []string
	DurationSecond  float64
	once            sync.Once
	frames          []*ebiten.Image
	maxFrameNum     int
	currentFrameNum int
	switchedTime    time.Time
	elapsed         float64
}

// Init loads asset images and initializes private parameters.
func (t *TimeAnimation) Init() (err error) {
	t.once.Do(func() {
		t.frames, err = loadImages(t.ImagesPaths)
		t.maxFrameNum = len(t.ImagesPaths)
	})
	if err != nil {
		return err
	}
	t.currentFrameNum = 0
	t.switchedTime = time.Now()
	t.elapsed = 0.0
	return nil
}

// GetCurrentFrame returns a current frame image. This function determines
// the current frame according to elapsed time.
// If the elapsed time is grater than the DurationSecond, this function
// will return the next frame.
func (t *TimeAnimation) GetCurrentFrame() *ebiten.Image {
	t.elapsed = time.Now().Sub(t.switchedTime).Seconds()
	if t.elapsed >= t.DurationSecond {
		t.currentFrameNum++
		t.switchedTime = time.Now()
	}
	if t.currentFrameNum < 0 || t.maxFrameNum-1 < t.currentFrameNum {
		t.currentFrameNum = 0
	}
	return t.frames[t.currentFrameNum]
}
