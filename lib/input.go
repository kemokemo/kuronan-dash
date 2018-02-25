// Copy from github.com/hajimehoshi/ebiten/example/blocks

package kuronandash

import (
	"github.com/hajimehoshi/ebiten"
)

// Input manages the input state including gamepads and keyboards.
type Input struct {
	keyStates                  map[ebiten.Key]int
	anyGamepadButtonPressed    bool
	virtualGamepadButtonStates map[virtualGamepadButton]int
	gamepadConfig              gamepadConfig
}

// StateForKey returns time length indicating how long the key is pressed.
func (i *Input) StateForKey(key ebiten.Key) int {
	if i.keyStates == nil {
		return 0
	}
	return i.keyStates[key]
}

// IsAnyGamepadButtonPressed returns a boolean value indicating
// whether any gamepad button is pressed.
func (i *Input) IsAnyGamepadButtonPressed() bool {
	return i.anyGamepadButtonPressed
}

func (i *Input) stateForVirtualGamepadButton(b virtualGamepadButton) int {
	if i.virtualGamepadButtonStates == nil {
		return 0
	}
	return i.virtualGamepadButtonStates[b]
}

func (i *Input) Update() {
	if i.keyStates == nil {
		i.keyStates = map[ebiten.Key]int{}
	}
	for key := ebiten.Key(0); key <= ebiten.KeyMax; key++ {
		if !ebiten.IsKeyPressed(ebiten.Key(key)) {
			i.keyStates[key] = 0
			continue
		}
		i.keyStates[key]++
	}

	const gamepadID = 0
	i.anyGamepadButtonPressed = false
	for b := ebiten.GamepadButton(0); b <= ebiten.GamepadButtonMax; b++ {
		if ebiten.IsGamepadButtonPressed(gamepadID, b) {
			i.anyGamepadButtonPressed = true
			break
		}
	}

	if i.virtualGamepadButtonStates == nil {
		i.virtualGamepadButtonStates = map[virtualGamepadButton]int{}
	}
	for _, b := range virtualGamepadButtons {
		if !i.gamepadConfig.IsButtonPressed(b) {
			i.virtualGamepadButtonStates[b] = 0
			continue
		}
		i.virtualGamepadButtonStates[b]++
	}
}

func (i *Input) IsRotateRightJustPressed() bool {
	if i.StateForKey(ebiten.KeySpace) == 1 || i.StateForKey(ebiten.KeyX) == 1 {
		return true
	}
	return i.stateForVirtualGamepadButton(virtualGamepadButtonButtonB) == 1
}

func (i *Input) IsRotateLeftJustPressed() bool {
	if i.StateForKey(ebiten.KeyZ) == 1 {
		return true
	}
	return i.stateForVirtualGamepadButton(virtualGamepadButtonButtonA) == 1
}

func (i *Input) StateForLeft() int {
	v := i.StateForKey(ebiten.KeyLeft)
	if 0 < v {
		return v
	}
	return i.stateForVirtualGamepadButton(virtualGamepadButtonLeft)
}

func (i *Input) StateForRight() int {
	v := i.StateForKey(ebiten.KeyRight)
	if 0 < v {
		return v
	}
	return i.stateForVirtualGamepadButton(virtualGamepadButtonRight)
}

func (i *Input) StateForDown() int {
	v := i.StateForKey(ebiten.KeyDown)
	if 0 < v {
		return v
	}
	return i.stateForVirtualGamepadButton(virtualGamepadButtonDown)
}
