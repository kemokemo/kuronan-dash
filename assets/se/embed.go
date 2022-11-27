package se

import _ "embed"

var (
	//go:embed click.wav
	click_wav []byte

	//go:embed jump.wav
	jump_wav []byte
	//go:embed drop.wav
	drop_wav []byte

	//go:embed attack-scratch.wav
	attack_scratch_wav []byte
	//go:embed attack-swipe.wav
	attack_swipe_wav []byte

	//go:embed title-call.wav
	title_call_wav []byte
	//go:embed character-select-voice.wav
	character_select_voice_wav []byte
	//go:embed ready.wav
	ready_wav []byte
	//go:embed go.wav
	go_wav []byte
	//go:embed stage-clear-voice.wav
	stage_clear_voice_wav []byte
	//go:embed sp-voice-kurona.wav
	sp_voice_kurona_wav []byte
	//go:embed sp-voice-koma.wav
	sp_voice_koma_wav []byte
	//go:embed sp-voice-shishimaru.wav
	sp_voice_shishimaru_wav []byte
)
