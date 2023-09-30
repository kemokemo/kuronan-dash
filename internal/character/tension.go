package character

import "github.com/kemokemo/kuronan-dash/internal/move"

// Tension is the character's tension.
// If this reaches the max value, the character can shot a skill skill.
type Tension struct {
	max     int
	val     int
	languor int
	valRate int
}

func NewTension(max, border int) *Tension {
	return &Tension{max: max, languor: border}
}

func (t *Tension) Initialize() {
	t.val = 0
	t.valRate = 0
}

// AddByState adds val to tension's val.
func (t *Tension) AddByState(state move.State) {
	switch state {
	case move.Dash:
		t.Add(2)
	case move.Walk:
		t.Add(1)
	case move.SkillDash, move.SkillWalk:
		t.remove(1)
	default:
		// not add tension.
	}
}

func (t *Tension) AddByAttack(brokenNum int) {
	t.Add(brokenNum * 5)
}

func (t *Tension) Add(val int) {
	for i := 0; i < val; i++ {
		t.valRate++
		if t.valRate >= t.languor {
			if t.val < t.max {
				t.val++
			}
			t.valRate = 0
		}
	}
}

func (t *Tension) remove(val int) {
	if t.val < 0 {
		return
	}
	t.val -= val
	if t.val < 0 {
		t.val = 0
	}
}

func (t *Tension) Get() int {
	return t.val
}

func (t *Tension) GetMax() float64 {
	return float64(t.max)
}

func (t *Tension) IsMax() bool {
	return t.val >= t.max
}
