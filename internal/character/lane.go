package character

import "fmt"

// Lanes manages the player's lane to run.
type Lanes struct {
	heights []int
	max     int
	current int
	target  int
	state   state
}

// SetHeights sets the heights of the lanes.
func (l *Lanes) SetHeights(heights []int) error {
	length := len(heights)
	if length == 0 {
		return fmt.Errorf("heights is empty")
	}

	l.heights = heights
	l.max = length
	return nil
}

// IsTop returns which the player's lane is top one.
func (l *Lanes) IsTop() bool {
	return l.current == 0
}

// Ascend sets the target lane to ascend.
func (l *Lanes) Ascend() bool {
	if l.current == 0 {
		return false
	}
	if l.state == ascending || l.state == descending {
		return false
	}
	l.target = l.current - 1
	l.state = ascending
	return true
}

// IsBottom returns which the player's lane is bottom one.
func (l *Lanes) IsBottom() bool {
	return l.current == l.max
}

// Descend sets the target lane to descend.
// When this function is triggered successfully, true is returned.
func (l *Lanes) Descend() bool {
	if l.current == (l.max - 1) {
		return false
	}
	if l.state == ascending || l.state == descending {
		return false
	}

	l.target = l.current + 1
	l.state = descending
	return true
}

// IsReachedTarget returns wchich the player reached the target.
func (l *Lanes) IsReachedTarget(height int) bool {
	if (l.state == ascending && l.heights[l.target] >= height) ||
		(l.state == descending && l.heights[l.target] <= height) {
		l.state = dash
		l.current = l.target
		return true
	}
	return false
}
