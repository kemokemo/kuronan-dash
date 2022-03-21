package se

import _ "embed"

var (
	//go:embed jump.wav
	jump_wav []byte
	//go:embed drop.wav
	drop_wav []byte

	//go:embed attack-scratch.wav
	attack_scratch_wav []byte
	//go:embed attack-swipe.wav
	attack_swipe_wav []byte
)
