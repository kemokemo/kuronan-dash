package scenes

import "github.com/kemokemo/kuronan-dash/internal/input"

// GameState describe the state of this game.
type GameState struct {
	SceneManager *SceneManager
	Input        *input.Input
}

// State is the state of the game scene.
type gameState int

const (
	wait gameState = iota
	run
	pause
	gameover
)
