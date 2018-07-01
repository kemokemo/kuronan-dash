package objects

// CharacterState describes the state of a character.
type CharacterState int

const (
	// Stop describes a character is stopping
	Stop CharacterState = iota
	// Walk describes a character is walking
	Walk
	// Dash is Dash!
	Dash
	// Ascending describes a character is jumping
	Ascending
	// Descending describes a character is descending
	Descending
)
