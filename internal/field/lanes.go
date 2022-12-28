package field

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

type LaneType string

const (
	PrairieLane LaneType = "PrairieLane"
)

// Lanes not only updates and draws multiple SingleLane objects, but also manages
// the height of the lane that the character is currently aiming for.
type Lanes struct {
	lanes       []*SingleLane
	laneHeights []float64
	laneIndex   int
}

func NewLanes(lType LaneType) *Lanes {
	var infos []ScrollInfo
	// add new field's lanes here
	switch lType {
	case PrairieLane:
		infos = newPrairieLanesInfo()
	default:
		infos = newPrairieLanesInfo()
	}

	var lanes []*SingleLane
	for i := range infos {
		l := &SingleLane{}
		l.Initialize(infos[i].img, infos[i].pos, infos[i].vel)
		lanes = append(lanes, l)
	}

	return &Lanes{
		lanes:       lanes,
		laneHeights: []float64{pLaneHeight1, pLaneHeight2, pLaneHeight3},
		laneIndex:   0,
	}
}

func (l *Lanes) Update(scrollV *view.Vector) {
	for i := range l.lanes {
		l.lanes[i].Update(scrollV)
	}
}

func (l *Lanes) Draw(screen *ebiten.Image) {
	for i := range l.lanes {
		l.lanes[i].Draw(screen)
	}
}

// GoToUpperLane sets the upper lane as the destination and returns true
// if the destination lane exists, false if not.
func (l *Lanes) GoToUpperLane() bool {
	if l.laneIndex <= 0 {
		l.laneIndex = 0
		return false
	}

	l.laneIndex--
	return true
}

// GoToLowerLane sets the lower lane as the destination and returns true
// if the destination lane exists, false if not.
func (l *Lanes) GoToLowerLane() bool {
	if l.laneIndex >= len(l.laneHeights)-1 {
		l.laneIndex = len(l.laneHeights) - 1
		return false
	}

	l.laneIndex++
	return true
}

func (l *Lanes) GetLaneHeights() []float64 {
	return l.laneHeights
}

func (l *Lanes) GetTargetLaneHeight() float64 {
	return l.laneHeights[l.laneIndex]
}
