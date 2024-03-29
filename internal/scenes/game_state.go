package scenes

// GameState describe the state of this game.
type GameState struct {
	SceneManager *SceneManager
}

// State is the state of the game scene.
type gameState int

const (
	wait gameState = iota
	readyCall
	goCall
	run
	pause
	skillEffect
	stageClear
	gameOver
)
