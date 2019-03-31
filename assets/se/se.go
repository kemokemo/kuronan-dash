package se

// sound effect
var (
	Jump *Player
)

// LoadSE loads all sound effects.
func LoadSE() error {
	var err error

	Jump, err = loadPlayer(jump_wav)
	if err != nil {
		return err
	}

	return nil
}

// CloseSE closes all sound effects.
func CloseSE() error {
	return Jump.Close()
}
