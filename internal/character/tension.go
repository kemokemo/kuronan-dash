package character

import "github.com/kemokemo/kuronan-dash/internal/move"

// Tension is the character's tension.
// If this reaches the max value, the character can shot a special skill.
type Tension struct {
	max     int
	val     int
	languor int
	valRate int
}

func NewTension(max, border int) *Tension {
	return &Tension{max: max, languor: border}
}

// AddByState adds val to tension's val.
func (t *Tension) AddByState(state move.State) {
	switch state {
	case move.Dash:
		t.add(2)
	case move.Walk:
		t.add(1)
	default:
		// not add tension.
	}
}

func (t *Tension) AddByAttack(brokenNum int) {
	t.add(brokenNum * 5)
}

func (t *Tension) add(val int) {
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

func (t *Tension) Get() int {
	return t.val
}
