package input

// InputChecker is the checker for user input.
type InputChecker interface {
	Update()
	TriggeredUp() bool
	TriggeredDown() bool
	TriggeredLeft() bool
	TriggeredRight() bool
	TriggeredStart() bool
	TriggeredPause() bool
	TriggeredAttack() bool
	TriggeredSpecial() bool
}

type VolumeChecker interface {
	Update()
	IsVolumeOn() bool
	JustVolumeOn() bool
	JustVolumeOff() bool
}
