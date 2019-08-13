package field

import "github.com/hajimehoshi/ebiten"

// Field is the interface to draw the field.
type Field interface {
	Initialize()
	SetScrollSpeed(speed ScrollSpeed)
	Update()
	Draw(screen *ebiten.Image) error
}

// lane information to draw
const repeat = 3

const (
	firstLaneHeight  = 200
	secondLaneHeight = firstLaneHeight + 170
	thirdLaneHeight  = secondLaneHeight + 170
)

// LaneHeights is the height array to draw lanes.
var LaneHeights = []int{firstLaneHeight, secondLaneHeight, thirdLaneHeight}
