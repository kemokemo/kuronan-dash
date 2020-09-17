package se

import "fmt"

// sound effect
var (
	Jump *Player
	Drop *Player
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
	return err
}
