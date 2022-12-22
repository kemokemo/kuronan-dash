package se

import "fmt"

// sound effect
var (
	Click *Player

	JumpSe *Player
	DropSe *Player

	AttackScratch *Player
	AttackSwipe   *Player

	TitleCall            *Player
	CharacterSelectVoice *Player
	ReadyVoice           *Player
	GoVoice              *Player
	StageClearVoice      *Player
	SpVoiceKurona        *Player
	SpVoiceKoma          *Player
	SpVoiceShishimaru    *Player
)

// LoadSE loads all sound effects.
func LoadSE() error {
	var err error

	Click, err = loadPlayer(click_wav)
	if err != nil {
		return err
	}
	JumpSe, err = loadPlayer(jump_wav)
	if err != nil {
		return err
	}
	DropSe, err = loadPlayer(drop_wav)
	if err != nil {
		return err
	}
	AttackScratch, err = loadPlayer(attack_scratch_wav)
	if err != nil {
		return err
	}
	AttackSwipe, err = loadPlayer(attack_swipe_wav)
	if err != nil {
		return err
	}
	TitleCall, err = loadPlayer(title_call_wav)
	if err != nil {
		return err
	}
	CharacterSelectVoice, err = loadPlayer(character_select_voice_wav)
	if err != nil {
		return err
	}
	ReadyVoice, err = loadPlayer(ready_wav)
	if err != nil {
		return err
	}
	GoVoice, err = loadPlayer(go_wav)
	if err != nil {
		return err
	}
	StageClearVoice, err = loadPlayer(stage_clear_voice_wav)
	if err != nil {
		return err
	}
	SpVoiceKurona, err = loadPlayer(sp_voice_kurona_wav)
	if err != nil {
		return err
	}
	SpVoiceKoma, err = loadPlayer(sp_voice_koma_wav)
	if err != nil {
		return err
	}
	SpVoiceShishimaru, err = loadPlayer(sp_voice_shishimaru_wav)
	if err != nil {
		return err
	}

	return nil
}

// CloseSE closes all sound effects.
func CloseSE() error {
	// todo: Close対象が少なすぎる・・

	var err, e error
	e = JumpSe.Close()
	if err != nil {
		err = fmt.Errorf("%v:%v", err, e)
	}
	e = DropSe.Close()
	if err != nil {
		err = fmt.Errorf("%v:%v", err, e)
	}
	e = AttackScratch.Close()
	if err != nil {
		err = fmt.Errorf("%v:%v", err, e)
	}
	e = AttackSwipe.Close()
	if err != nil {
		err = fmt.Errorf("%v:%v", err, e)
	}
	return err
}
