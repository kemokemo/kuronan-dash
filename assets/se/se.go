package se

import "fmt"

// sound effect
var (
	Jump *Player
	Drop *Player

	attackScratch *Player
	attackSwipe   *Player

	TitleCall            *Player
	CharacterSelectVoice *Player
)

// LoadSE loads all sound effects.
func LoadSE() error {
	var err error

	Jump, err = loadPlayer(jump_wav)
	if err != nil {
		return err
	}
	Drop, err = loadPlayer(drop_wav)
	if err != nil {
		return err
	}
	attackScratch, err = loadPlayer(attack_scratch_wav)
	if err != nil {
		return err
	}
	attackSwipe, err = loadPlayer(attack_swipe_wav)
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

	return nil
}

// CloseSE closes all sound effects.
func CloseSE() error {
	var err, e error
	e = Jump.Close()
	if err != nil {
		err = fmt.Errorf("%v:%v", err, e)
	}
	e = Drop.Close()
	if err != nil {
		err = fmt.Errorf("%v:%v", err, e)
	}
	e = attackScratch.Close()
	if err != nil {
		err = fmt.Errorf("%v:%v", err, e)
	}
	e = attackSwipe.Close()
	if err != nil {
		err = fmt.Errorf("%v:%v", err, e)
	}
	return err
}

// sound type
type SoundType int

const (
	KuronaSe SoundType = iota
	KomaSe
	ShishimaruSe
)

func GetAttackSe(st SoundType) *Player {
	switch st {
	case KuronaSe:
		return attackScratch
	case KomaSe, ShishimaruSe:
		return attackSwipe
	default:
		return attackScratch
	}
}
