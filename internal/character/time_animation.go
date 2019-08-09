package character

import (
	"time"

	"github.com/hajimehoshi/ebiten"
)

// TimeAnimation is an animation.
// This animates according to elapsed time.
type TimeAnimation struct {
	frames          []*ebiten.Image
	durationSecond  float64
	maxFrameNum     int
	currentFrameNum int
	switchedTime    time.Time
	elapsed         float64
}

// NewTimeAnimation returns a new TimeAnimation generated with args.
func NewTimeAnimation(frames []*ebiten.Image, duration float64) *TimeAnimation {
	return &TimeAnimation{
		frames:         frames,
		durationSecond: duration,
		maxFrameNum:    len(frames),
		switchedTime:   time.Now(),
		elapsed:        0.0,
	}
}

// GetCurrentFrame returns a current frame image. This function determines
// the current frame according to elapsed time.
// If the elapsed time is grater than the DurationSecond, this function
// will return the next frame.
func (t *TimeAnimation) GetCurrentFrame() *ebiten.Image {
	t.elapsed = time.Now().Sub(t.switchedTime).Seconds()
	if t.elapsed >= t.durationSecond {
		t.currentFrameNum++
		t.switchedTime = time.Now()
	}
	if t.currentFrameNum < 0 || t.maxFrameNum-1 < t.currentFrameNum {
		t.currentFrameNum = 0
	}
	return t.frames[t.currentFrameNum]
}
