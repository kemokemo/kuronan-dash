package move

import (
	"fmt"
	"image"
	"log"

	vpad "github.com/kemokemo/ebiten-virtualpad"
	"github.com/kemokemo/kuronan-dash/assets/se"
	"github.com/kemokemo/kuronan-dash/internal/field"
	"github.com/kemokemo/kuronan-dash/internal/input"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// StateMachine manages the player's state.
type StateMachine struct {
	pos         *view.Vector
	current     State
	previous    State
	isBlocked   bool
	jumpSe      *se.Player
	dropSe      *se.Player
	attackSe    *se.Player
	lanes       *field.Lanes
	offset      *view.Vector
	iChecker    input.InputChecker
	attacked    bool
	drawing     bool
	duration    int
	maxDuration int
}

func NewStateMachine(lanes *field.Lanes, typeSe se.SoundType, maxDuration int) (*StateMachine, error) {
	heights := lanes.GetLaneHeights()
	if len(heights) == 0 {
		return nil, fmt.Errorf("heights is empty")
	}

	sm := StateMachine{
		pos:         &view.Vector{X: view.DrawPosition, Y: lanes.GetTargetLaneHeight()},
		current:     Dash,
		previous:    Pause,
		isBlocked:   false,
		jumpSe:      se.Jump,
		dropSe:      se.Drop,
		attackSe:    se.GetAttackSe(typeSe),
		lanes:       lanes,
		offset:      &view.Vector{X: 0.0, Y: 0.0},
		maxDuration: maxDuration,
		duration:    maxDuration,
	}

	return &sm, nil
}

func (sm *StateMachine) SetInputChecker(laneRectArray []image.Rectangle, upBtn, downBtn, atkBtn vpad.TriggerButton) {
	sm.iChecker = &input.PlayerInputChecker{
		RectArray: laneRectArray,
		UpBtn:     upBtn,
		DownBtn:   downBtn,
		AttackBtn: atkBtn,
	}
}

// Update updates the state.
func (sm *StateMachine) Update(stamina int, charaPosV *view.Vector) State {
	sm.pos.Add(charaPosV)
	sm.offset.Y = 0.0

	sm.updateWithStaminaAndMove(stamina, charaPosV)
	sm.updateWithKey(charaPosV.Y)

	return sm.current
}

func (sm *StateMachine) updateWithStaminaAndMove(stamina int, charaPosV *view.Vector) {
	switch sm.current {
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
	case Walk:
		if stamina > 0 {
			sm.previous = Walk
			sm.current = Dash
		}
	default:
		log.Println("unknown state: ", sm.current)
	}
}

func (sm *StateMachine) updateWithKey(vY float64) {
	if !(sm.current == Dash) && !(sm.current == Walk) {
		return
	}

	sm.iChecker.Update()

	if sm.iChecker.TriggeredUp() {
		if !sm.lanes.GoToUpperLane() {
			return
		}

		sm.previous = sm.current
		sm.current = Ascending
		err := sm.jumpSe.Play()
		if err != nil {
			log.Println("failed to play jump SE: ", err)
		}
	} else if sm.iChecker.TriggeredDown() {
		if !sm.lanes.GoToLowerLane() {
			return
		}

		sm.previous = sm.current
		sm.current = Descending
		err := sm.dropSe.Play()
		if err != nil {
			log.Println("failed to play drop SE: ", err)
		}
	}

	if sm.duration < sm.maxDuration {
		if sm.drawing {
			sm.duration++
		}
		sm.attacked = false
	} else {
		if sm.iChecker.TriggeredAttack() {
			sm.attacked = true
			sm.drawing = true
			sm.duration = 0
			err := sm.attackSe.Play()
			if err != nil {
				log.Println("failed to play attack SE: ", err)
			}
		} else {
			sm.attacked = false
			sm.drawing = false
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

func (sm *StateMachine) Attacked() bool {
	return sm.attacked
}

func (sm *StateMachine) DrawAttack() bool {
	return sm.drawing
}
