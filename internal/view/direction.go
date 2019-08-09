package view

// Direction is the direction for the Viewport.
// If you want to move the image to the UpperRight, use the UpperRight direction.
type Direction int

// Directions
const (
	None Direction = iota
	Left
	Right
	Upper
	UpperLeft
	UpperRight
	Lower
	LowerLeft
	LowerRight
)

func getDirectionValue(d Direction) (x, y int) {
	switch d {
	case Left:
		return -1, 0
	case Right:
		return 1, 0
	case Upper:
		return 0, -1
	case UpperLeft:
		return -1, -1
	case UpperRight:
		return 1, -1
	case Lower:
		return 0, 1
	case LowerLeft:
		return -1, 1
	case LowerRight:
		return 1, 1
	default:
		return 0, 0
	}
}
