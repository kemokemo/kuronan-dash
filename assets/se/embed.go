package se

import _ "embed"

var (
	//go:embed menu-select.wav
	menu_select_wav []byte

	//go:embed jump.wav
	jump_wav []byte
	//go:embed drop.wav
	drop_wav []byte
	//go:embed collision.wav
	collision_wav []byte
	//go:embed break-rock.wav
	break_rock_wav []byte
	//go:embed pickup-item.wav
	pickup_item_wav []byte

	//go:embed attack-scratch.wav
	attack_scratch_wav []byte
	//go:embed attack-punch.wav
	attack_punch_wav []byte

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
