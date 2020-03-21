package character

// State describes the state of a character.
type State int

// States
const (
	Pause State = iota
	Walk
	Dash
	Ascending
	Descending
)
