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
	pos            *view.Vector
	current        State
	previous       State
	isBlocked      bool
	lanes          *field.Lanes
	offset         *view.Vector
	iChecker       input.InputChecker
	attacked       bool
	drawing        bool
	atkDuration    int
	atkMaxDuration int
	startSpEffect  bool
	finishSpEffect bool
	spDuration     int
	spMaxDuration  int
	soundTypeCh    chan<- se.SoundType
	xOffset        displayOffset
}

func NewStateMachine(lanes *field.Lanes, atkMaxDuration int, spMaxDuration int) (*StateMachine, error) {
	heights := lanes.GetLaneHeights()
	if len(heights) == 0 {
		return nil, fmt.Errorf("heights is empty")
	}

	sm := StateMachine{
		pos:            &view.Vector{X: view.DrawPosition, Y: lanes.GetTargetLaneHeight()},
		current:        Dash,
		previous:       Pause,
		isBlocked:      false,
		lanes:          lanes,
		offset:         &view.Vector{X: 0.0, Y: 0.0},
		atkMaxDuration: atkMaxDuration,
		atkDuration:    atkMaxDuration,
		spDuration:     0,
		spMaxDuration:  spMaxDuration,
	}

	return &sm, nil
}

func (sm *StateMachine) SetInputChecker(laneRectArray []image.Rectangle, upBtn, downBtn, atkBtn, spBtn vpad.TriggerButton) {
	sm.iChecker = &input.PlayerInputChecker{
		RectArray: laneRectArray,
		UpBtn:     upBtn,
		DownBtn:   downBtn,
		AttackBtn: atkBtn,
		SkillBtn:  spBtn,
	}
}

// Update updates the state.
func (sm *StateMachine) Update(stamina int, tension int, isMaxTension bool, charaPosV *view.Vector) State {
	sm.pos.Add(charaPosV)
	sm.offset.Y = 0.0

	sm.updateWithStaminaAndMove(stamina, tension, charaPosV)
	sm.updateWithKey(isMaxTension, charaPosV.Y)
	sm.updateXAxisOffset()

	return sm.current
}

func (sm *StateMachine) updateWithStaminaAndMove(stamina int, tension int, charaPosV *view.Vector) {
	switch sm.current {
	case Ascending:
		if sm.IsReachedUpperLane(charaPosV.Y) {
			sm.current = sm.previous
		}
	case Descending:
		if sm.isReachedLowerLane(charaPosV.Y) {
			sm.current = sm.previous
		}
	case Dash:
		if stamina <= 0 || sm.isBlocked {
			sm.previous = Dash
			sm.current = Walk
		}
	case SkillDash:
		sm.finishSpEffect = false
		if stamina <= 0 || sm.isBlocked {
			sm.previous = SkillDash
			sm.current = Walk
		} else if tension <= 0 {
			sm.previous = SkillDash
			sm.current = Dash
		}
	case Walk:
		if stamina > 0 && !sm.isBlocked {
			if sm.previous == SkillDash {
				sm.current = SkillDash
			} else {
				sm.current = Dash
			}
			sm.previous = Walk
		}
	default:
		log.Println("unknown state: ", sm.current)
	}
}

func (sm *StateMachine) updateWithKey(isMaxTension bool, vY float64) {
	if !(sm.current == Dash) && !(sm.current == Walk) && !(sm.current == SkillDash) && !(sm.current == SkillEffect) {
		return
	}

	sm.iChecker.Update()

	if sm.iChecker.TriggeredUp() {
		if !sm.lanes.GoToUpperLane() {
			return
		}

		sm.previous = sm.current
		sm.current = Ascending
		sm.soundTypeCh <- se.Jump
	} else if sm.iChecker.TriggeredDown() {
		if !sm.lanes.GoToLowerLane() {
			return
		}

		sm.previous = sm.current
		sm.current = Descending
		sm.soundTypeCh <- se.Drop
	}

	if sm.atkDuration < sm.atkMaxDuration {
		if sm.drawing {
			sm.atkDuration++
		}
		sm.attacked = false
	} else {
		if sm.iChecker.TriggeredAttack() {
			sm.attacked = true
			sm.drawing = true
			sm.atkDuration = 0
			sm.soundTypeCh <- se.Attack
		} else {
			sm.attacked = false
			sm.drawing = false
		}
	}

	if sm.iChecker.TriggeredSkill() && isMaxTension {
		sm.previous = sm.current
		sm.current = SkillEffect
		sm.startSpEffect = true
	}
}

func (sm *StateMachine) updateXAxisOffset() {
	sm.xOffset.Update(sm.current)
	sm.offset.X = sm.xOffset.GetXAxisOffset()
}

func (sm *StateMachine) UpdateSkillEffect(playingSound bool) {
	sm.startSpEffect = false
	sm.spDuration++
	if sm.spDuration >= sm.spMaxDuration && !playingSound {
		sm.spDuration = 0
		sm.finishSpEffect = true
		sm.current = SkillDash
	}
}

// IsReachedUpperLane returns which the player reached to the target upper lane.
// If reached to the target lane, sm.offset is set.
func (sm *StateMachine) IsReachedUpperLane(vY float64) bool {
	if vY > 0 {
		return sm.isReachedLowerLane(vY)
	}

	return false
}

// isReachedLowerLane returns which the player reached to the target lower lane.
// If reached to the target lane, sm.offset is set.
func (sm *StateMachine) isReachedLowerLane(vY float64) bool {
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

func (sm *StateMachine) StartSpEffect() bool {
	return sm.startSpEffect
}

func (sm *StateMachine) FinishSpEffect() bool {
	return sm.finishSpEffect
}

func (sm *StateMachine) SetSeChan(ch chan<- se.SoundType) {
	sm.soundTypeCh = ch
}
