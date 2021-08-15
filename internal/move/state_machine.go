package move

import (
	"fmt"
	"log"

	"github.com/kemokemo/kuronan-dash/assets/se"
	"github.com/kemokemo/kuronan-dash/internal/input"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// StateMachine manages the player's state.
type StateMachine struct {
	pos            *view.Vector
	laneHeights    []float64
	maxLaneNum     int
	currentLaneNum int
	targetLaneNum  int
	current        State
	previous       State
	isBlocked      bool
	jumpSe         *se.Player
	dropSe         *se.Player
}

func NewStateMachine() StateMachine {
	return StateMachine{
		pos:            &view.Vector{X: view.DrawPosition, Y: 0.0},
		currentLaneNum: 0,
		current:        Dash,
		previous:       Pause,
		isBlocked:      false,
		jumpSe:         se.Jump,
		dropSe:         se.Drop,
	}
}

// SetHeights sets the heights of the lanes.
func (sm *StateMachine) SetHeights(heights []float64) error {
	sm.maxLaneNum = len(heights)
	if sm.maxLaneNum == 0 {
		return fmt.Errorf("heights is empty")
	}

	sm.laneHeights = heights
	sm.pos.Y = heights[0]
	return nil
}

// Update updates the state.
func (sm *StateMachine) Update(stamina int, charaPosV *view.Vector) State {
	sm.pos.Add(charaPosV)

	switch sm.current {
	case Pause:
		if input.TriggeredOne() {
			sm.current = sm.previous
		}
	case Ascending:
		if sm.IsTop(charaPosV.Y) {
			sm.current = sm.previous
		}
	case Descending:
		if sm.IsBottom(charaPosV.Y) {
			sm.current = sm.previous
		}
	case Dash:
		if stamina <= 0 || sm.isBlocked {
			sm.previous = Dash
			sm.current = Walk
		}
		sm.keyCheck(charaPosV.Y)
	case Walk:
		if stamina > 0 {
			sm.previous = Walk
			sm.current = Dash
		}
		sm.keyCheck(charaPosV.Y)
	default:
		log.Println("unknown state: ", sm.current)
	}
	return sm.current
}

func (sm *StateMachine) keyCheck(vY float64) {
	if input.TriggeredOne() {
		sm.previous = sm.current
		sm.current = Pause
	} else if input.TriggeredUp() {
		if !sm.IsTop(vY) {
			sm.targetLaneNum = sm.currentLaneNum - 1
			sm.previous = sm.current
			sm.current = Ascending
			sm.jumpSe.Play() // TODO: error handling
		}
	} else if input.TriggeredDown() {
		if !sm.IsBottom(vY) {
			sm.targetLaneNum = sm.currentLaneNum + 1
			sm.previous = sm.current
			sm.current = Descending
			sm.dropSe.Play() // TODO: error handling
		}
	}
}

// IsTop returns which the player's lane is top one.
func (sm *StateMachine) IsTop(vY float64) bool {
	if sm.currentLaneNum == 0 {
		return true
	}
	if sm.IsReachedTarget(vY) {
		return true
	}

	return false
}

// IsBottom returns which the player's lane is bottom one.
func (sm *StateMachine) IsBottom(vY float64) bool {
	if sm.currentLaneNum == (sm.maxLaneNum - 1) {
		return true
	}
	if sm.IsReachedTarget(vY) {
		return true
	}

	return false
}

// IsReachedTarget returns which the player reached the target.
//
// TODO: In fact, I think the distance to the landing should be feedback to the 'VelocityController'
// to handle it so that it doesn't sink into the ground.
func (sm *StateMachine) IsReachedTarget(vY float64) bool {
	var res bool
	switch sm.current {
	case Ascending:
		// In order to make a fluffy landing, the landing during the ascent process was changed to only when the speed was downward.
		if vY > 0 {
			// A way to make it look like you are grounded on the lane. I think it can be improved a little more.
			if sm.laneHeights[sm.targetLaneNum]-2.0 <= sm.pos.Y && sm.pos.Y <= sm.laneHeights[sm.targetLaneNum]+2.0 {
				sm.current = sm.previous
				sm.currentLaneNum = sm.targetLaneNum
				res = true
			}
		}
	case Descending:
		if sm.laneHeights[sm.targetLaneNum] <= sm.pos.Y {
			sm.current = sm.previous
			sm.currentLaneNum = sm.targetLaneNum
			res = true
		}
	default:
		res = false
	}
	return res
}

// SetBlockState sets the blocked state of player.
func (sm *StateMachine) SetBlockState(isBlocked bool) {
	sm.isBlocked = isBlocked
}

func (sm *StateMachine) GetPosition() *view.Vector {
	return sm.pos
}
