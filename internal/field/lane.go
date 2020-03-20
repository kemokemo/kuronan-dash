package field

// lane information to draw
const repeat = 3

const (
	firstLaneHeight  = 200.0
	secondLaneHeight = firstLaneHeight + 170.0
	thirdLaneHeight  = secondLaneHeight + 170.0
)

// LaneHeights is the height array to draw lanes.
var LaneHeights = []float64{firstLaneHeight, secondLaneHeight, thirdLaneHeight}
