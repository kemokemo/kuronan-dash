package assets

type GameSoundControl int

const (
	PauseGameSound GameSoundControl = iota
	StartGameSound
	StopGameSound
	SoundOn
	SoundOff
)
