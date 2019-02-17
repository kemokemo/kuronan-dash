package character

// State describes the state of a character.
type State int

const (
	// Stop describes a character is stopping
	Stop State = iota
	// Walk describes a character is walking
	Walk
	// Dash is Dash!
	Dash
	// Ascending describes a character is jumping
	Ascending
	// Descending describes a character is descending
	Descending
)
