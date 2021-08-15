package move

// State describes the state of a character.
type State int

// States
const (
	Dash State = iota
	Walk
	Ascending
	Descending
	Pause
)
