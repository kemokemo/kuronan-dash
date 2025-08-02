package se

import "fmt"

// sound effect
var (
	MenuSelect *Player

	JumpSe      *Player
	DropSe      *Player
	CollisionSe *Player
	BreakRock   *Player
	PickupItem  *Player

	AttackScratch *Player
	AttackPunch   *Player

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

	MenuSelect, err = loadPlayer(menu_select_wav)
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
	CollisionSe, err = loadPlayer(collision_wav)
	if err != nil {
		return err
	}
	BreakRock, err = loadPlayer(break_rock_wav)
	if err != nil {
		return err
	}
	PickupItem, err = loadPlayer(pickup_item_wav)
	if err != nil {
		return err
	}
	AttackScratch, err = loadPlayer(attack_scratch_wav)
	if err != nil {
		return err
	}
	AttackPunch, err = loadPlayer(attack_punch_wav)
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
	var err, e error

	e = MenuSelect.Close()
	if e != nil {
		err = fmt.Errorf("%v:%v", err, e)
	}
	e = JumpSe.Close()
	if e != nil {
		err = fmt.Errorf("%v:%v", err, e)
	}
	e = DropSe.Close()
	if e != nil {
		err = fmt.Errorf("%v:%v", err, e)
	}
	e = CollisionSe.Close()
	if e != nil {
		err = fmt.Errorf("%v:%v", err, e)
	}
	e = BreakRock.Close()
	if e != nil {
		err = fmt.Errorf("%v:%v", err, e)
	}
	e = PickupItem.Close()
	if e != nil {
		err = fmt.Errorf("%v:%v", err, e)
	}
	e = AttackScratch.Close()
	if e != nil {
		err = fmt.Errorf("%v:%v", err, e)
	}
	e = AttackPunch.Close()
	if e != nil {
		err = fmt.Errorf("%v:%v", err, e)
	}
	e = TitleCall.Close()
	if e != nil {
		err = fmt.Errorf("%v:%v", err, e)
	}
	e = CharacterSelectVoice.Close()
	if e != nil {
		err = fmt.Errorf("%v:%v", err, e)
	}
	e = ReadyVoice.Close()
	if e != nil {
		err = fmt.Errorf("%v:%v", err, e)
	}
	e = GoVoice.Close()
	if e != nil {
		err = fmt.Errorf("%v:%v", err, e)
	}
	e = StageClearVoice.Close()
	if e != nil {
		err = fmt.Errorf("%v:%v", err, e)
	}
	e = SpVoiceKurona.Close()
	if e != nil {
		err = fmt.Errorf("%v:%v", err, e)
	}
	e = SpVoiceKoma.Close()
	if e != nil {
		err = fmt.Errorf("%v:%v", err, e)
	}
	e = SpVoiceShishimaru.Close()
	if e != nil {
		err = fmt.Errorf("%v:%v", err, e)
	}

	return err
}
