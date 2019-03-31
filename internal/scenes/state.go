package scenes

// State is the state of the game scene.
type state int

const (
	beforeRun state = iota
	running
	pause
	gameover
)
