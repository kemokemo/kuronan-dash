package character

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

// Consumes encourages a decrease in stamina by the value specified in the argument.
// The actual stamina is reduced when the consumption is greater than the "endurance".
// The stamina value does not go below zero.
func (s *Stamina) Consumes(val int) {
	s.valRate -= val
	if s.valRate <= 0 {
		if s.val > 0 {
			s.val--
		}
		s.valRate = s.endurance
	}
}

// Restore restores the stamina to the value specified in the argument.
// However, it will not restore more than the maximum value.
func (s *Stamina) Restore(val int) {
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
