package move

// State describes the state of a character.
type State int

// States
const (
	Wait State = iota
	Dash
	Walk
	Ascending
	Descending
	Pause
	SpecialEffect
	Special
)

func (s State) String() string {
	var str = ""
	switch s {
	case Wait:
		str = "Wait"
	case Dash:
		str = "Dash"
	case Walk:
		str = "Walk"
	case Ascending:
		str = "Ascending"
	case Descending:
		str = "Descending"
	case Pause:
		str = "Pause"
	case SpecialEffect:
		str = "SpecialEffect"
	case Special:
		str = "Special"
	}
	return str
}
