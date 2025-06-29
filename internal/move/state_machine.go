package move

import (
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	vpad "github.com/kemokemo/ebiten-virtualpad"
	"github.com/kemokemo/kuronan-dash/assets/se"
	"github.com/kemokemo/kuronan-dash/internal/field"
	"github.com/kemokemo/kuronan-dash/internal/input"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// StateMachine manages the player's state.
type StateMachine struct {
	pos               *view.Vector
	current           State
	isBlocked         bool
	lanes             *field.Lanes
	offset            *view.Vector
	iChecker          input.InputChecker
	attacked          bool
	drawing           bool
	atkDuration       int
	atkMaxDuration    int
	spDuration        int
	spMaxDuration     int
	soundTypeCh       chan<- se.SoundType
	effectCompletedCh <-chan any
	xOffset           displayOffset
	collisionCounter  int
}

func NewStateMachine(lanes *field.Lanes, atkMaxDuration int, spMaxDuration int) (*StateMachine, error) {
	heights := lanes.GetLaneHeights()
	if len(heights) == 0 {
		return nil, fmt.Errorf("heights is empty")
	}

	sm := StateMachine{
		pos:              &view.Vector{X: view.DrawPosition, Y: lanes.GetTargetLaneHeight()},
		current:          Dash,
		isBlocked:        false,
		lanes:            lanes,
		offset:           &view.Vector{X: 0.0, Y: 0.0},
		atkMaxDuration:   atkMaxDuration,
		atkDuration:      atkMaxDuration,
		spDuration:       0,
		spMaxDuration:    spMaxDuration,
		collisionCounter: 100,
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
		DoubleClk: input.NewDoubleClick(ebiten.MouseButtonLeft),
	}
}

// Update updates the state.
func (sm *StateMachine) Update(stamina int, tension int, isMaxTension bool, charaPosV *view.Vector) State {
	sm.pos.Add(charaPosV)
	sm.offset.Y = 0.0

	sm.checkCollisionAction()
	sm.updateWithStaminaAndMove(stamina, tension, charaPosV)
	sm.updateWithKey(isMaxTension, charaPosV.Y)
	sm.updateXAxisOffset()

	return sm.current
}

func (sm *StateMachine) checkCollisionAction() {
	if !sm.isBlocked {
		return
	}

	sm.collisionCounter++
	if sm.collisionCounter < 10 {
		return
	}

	sm.soundTypeCh <- se.Blocked
	sm.collisionCounter = 0
}

func (sm *StateMachine) updateWithStaminaAndMove(stamina int, tension int, charaPosV *view.Vector) {
	switch sm.current {
	case Ascending:
		if sm.IsReachedUpperLane(charaPosV.Y) {
			sm.current = Dash
		}
	case SkillAscending:
		if sm.IsReachedUpperLane(charaPosV.Y) {
			sm.current = SkillDash
		}
	case Descending:
		if sm.isReachedLowerLane(charaPosV.Y) {
			sm.current = Dash
		}
	case SkillDescending:
		if sm.isReachedLowerLane(charaPosV.Y) {
			sm.current = SkillDash
		}
	case Dash:
		if stamina <= 0 || sm.isBlocked {
			sm.current = Walk
		}
	case Walk:
		if stamina > 0 && !sm.isBlocked {
			sm.current = Dash
		}
	case SkillEffect:
		// スキルのエフェクト完了待ち
	case SkillDash:
		if tension <= 0 {
			sm.current = Dash
		} else if stamina <= 0 || sm.isBlocked {
			sm.current = SkillWalk
		}
	case SkillWalk:
		if tension <= 0 {
			sm.current = Walk
		} else if stamina > 0 && !sm.isBlocked {
			sm.current = SkillDash
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
			// レーン移動中はキー入力による状態遷移は行わない
			return
		}

		switch sm.current {
		case SkillWalk, SkillDash:
			sm.current = SkillAscending
			sm.soundTypeCh <- se.Jump
		case Walk, Dash:
			sm.current = Ascending
			sm.soundTypeCh <- se.Jump
		default:
			// 状態遷移しない
		}
	} else if sm.iChecker.TriggeredDown() {
		if !sm.lanes.GoToLowerLane() {
			// レーン移動中はキー入力による状態遷移は行わない
			return
		}

		switch sm.current {
		case SkillWalk, SkillDash:
			sm.current = SkillDescending
			sm.soundTypeCh <- se.Drop
		case Walk, Dash:
			sm.current = Descending
			sm.soundTypeCh <- se.Drop
		default:
			// 状態遷移しない
		}
	}

	if sm.atkDuration < sm.atkMaxDuration {
		// 攻撃モーション中は攻撃ボタンの判定を行わない
		// todo: キャラによって攻撃の連射速度とか変えたいね
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
		sm.current = SkillEffect
		sm.soundTypeCh <- se.SpVoice
	}
}

func (sm *StateMachine) updateXAxisOffset() {
	sm.xOffset.Update(sm.current)
	sm.offset.X = sm.xOffset.GetXAxisOffset()
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
		return true
	}

	return false
}

// SetBlockState sets the blocked state of player.
func (sm *StateMachine) SetBlockState(isBlocked bool) {
	if !sm.isBlocked && isBlocked {
		// 障害物に当たり始めたタイミングで音を鳴らしたい
		sm.collisionCounter = 100
	}
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

func (sm *StateMachine) SetSeChan(ch chan<- se.SoundType) {
	sm.soundTypeCh = ch
}

func (sm *StateMachine) FinishSkillEffect() State {
	sm.current = SkillDash
	return sm.current
}
