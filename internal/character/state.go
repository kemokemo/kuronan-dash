package character

// State describes the state of a character.
type state int

// States
const (
	stop state = iota
	walk
	dash
	ascending
	descending
)
