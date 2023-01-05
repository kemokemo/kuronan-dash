package character

import "github.com/kemokemo/kuronan-dash/internal/move"

// Stamina manages the consumption and recovery of stamina.
type Stamina struct {
	max       int
	val       int
	endurance int
	valRate   int
}

// NewStamina returns a new stamina engine.
func NewStamina(max, endurance int) *Stamina {
	return &Stamina{
		max:       max,
		val:       max,
		endurance: endurance,
		valRate:   endurance,
	}
}

// Initialize initializes stamina engine.
func (s *Stamina) Initialize() {
	s.val = s.max
	s.valRate = s.endurance
}

// ConsumesByState encourages a decrease in stamina by the value specified in the argument.
// The actual stamina is reduced when the consumption is greater than the "endurance".
// The stamina value does not go below zero.
func (s *Stamina) ConsumesByState(state move.State) {
	switch state {
	case move.Dash, move.SkillDash:
		s.consumes(2)
	case move.Walk, move.SkillWalk:
		s.consumes(1)
	default:
		// not consume stamina.
	}
}

func (s *Stamina) ConsumeByAttack(brokenNum int) {
	s.consumes(brokenNum * 5)
}

func (s *Stamina) consumes(val int) {
	s.valRate -= val
	if s.valRate <= 0 {
		if s.val > 0 {
			s.val--
		}
		s.valRate = s.endurance
	}
}

// Add restores the stamina to the value specified in the argument.
// However, it will not restore more than the maximum value.
func (s *Stamina) Add(val int) {
	v := s.val + val
	if v > s.max {
		s.val = s.max
	} else {
		s.val = v
	}
}

// GetStamina returns the current stamina value.
func (s *Stamina) GetStamina() int {
	return s.val
}

// GetMaxStamina returns the max value of this stamina.
func (s *Stamina) GetMaxStamina() float64 {
	return float64(s.max)
}
