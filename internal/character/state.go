package character

// State describes the state of a character.
type State int

// States
const (
	Stop State = iota
	Walk
	Dash
	Ascending
	Descending
	Pause
)
