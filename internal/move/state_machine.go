package move

import (
	"fmt"
	"log"

	"github.com/kemokemo/kuronan-dash/assets/se"
	"github.com/kemokemo/kuronan-dash/internal/field"
	"github.com/kemokemo/kuronan-dash/internal/input"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// StateMachine manages the player's state.
type StateMachine struct {
	pos       *view.Vector
	current   State
	previous  State
	isBlocked bool
	jumpSe    *se.Player
	dropSe    *se.Player
	lanes     *field.Lanes
	offset    *view.Vector
}

func NewStateMachine(lanes *field.Lanes) (*StateMachine, error) {
	heights := lanes.GetLaneHeights()
	if len(heights) == 0 {
		return nil, fmt.Errorf("heights is empty")
	}

	sm := StateMachine{
		pos:       &view.Vector{X: view.DrawPosition, Y: lanes.GetTargetLaneHeight()},
		current:   Dash,
		previous:  Pause,
		isBlocked: false,
		jumpSe:    se.Jump,
		dropSe:    se.Drop,
		lanes:     lanes,
		offset:    &view.Vector{X: 0.0, Y: 0.0},
	}

	return &sm, nil
}

// Update updates the state.
func (sm *StateMachine) Update(stamina int, charaPosV *view.Vector) State {
	sm.pos.Add(charaPosV)
	sm.offset.Y = 0.0

	switch sm.current {
	case Pause:
		if input.TriggeredOne() {
			sm.current = sm.previous
		}
	case Ascending:
		if sm.IsReachedUpperLane(charaPosV.Y) {
			sm.current = sm.previous
		}
	case Descending:
		if sm.IsReachedLowerLane(charaPosV.Y) {
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
		if sm.lanes.GoToUpperLane() {
			sm.previous = sm.current
			sm.current = Ascending

			err := sm.jumpSe.Play()
			if err != nil {
				log.Println("failed to play jump SE: ", err)
			}
		}
	} else if input.TriggeredDown() {
		if sm.lanes.GoToLowerLane() {
			sm.previous = sm.current
			sm.current = Descending

			err := sm.dropSe.Play()
			if err != nil {
				log.Println("failed to play drop SE: ", err)
			}
		}
	}
}

// IsReachedUpperLane returns which the player reached to the target upper lane.
// If reached to the target lane, sm.offset is set.
func (sm *StateMachine) IsReachedUpperLane(vY float64) bool {
	if vY > 0 {
		return sm.IsReachedLowerLane(vY)
	}

	return false
}

// IsReachedLowerLane returns which the player reached to the target lower lane.
// If reached to the target lane, sm.offset is set.
func (sm *StateMachine) IsReachedLowerLane(vY float64) bool {
	th := sm.lanes.GetTargetLaneHeight()

	nextPosY := sm.pos.Y + vY
	if nextPosY >= th {
		sm.offset.Y = th - nextPosY
		sm.current = sm.previous
		return true
	}

	return false
}

// SetBlockState sets the blocked state of player.
func (sm *StateMachine) SetBlockState(isBlocked bool) {
	sm.isBlocked = isBlocked
}

func (sm *StateMachine) GetPosition() *view.Vector {
	return sm.pos
}

func (sm *StateMachine) GetOffsetV() *view.Vector {
	return sm.offset
}
