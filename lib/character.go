package kuronandash

import "github.com/hajimehoshi/ebiten"

// Position describes the position by x and y.
type Position struct {
	X int
	Y int
}

// NewCharacter creates a new character instance.
func NewCharacter(imagePaths []string) (*Character, error) {
	c := &Character{
		animation: StepAnimation{
			ImagesPaths:   imagePaths,
			DurationSteps: 5,
		},
	}
	err := c.animation.Init()
	if err != nil {
		return nil, err
	}
	return c, nil
}

// Character describes a character.
type Character struct {
	animation StepAnimation
	position  Position
	moved     bool
}

// SetInitialPosition sets the initial position for this character.
func (c *Character) SetInitialPosition(pos Position) {
	c.position = pos
}

// Move moves the character regarding the user input.
func (c *Character) Move() {
	c.moved = false
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.GamepadAxis(0, 0) <= -0.5 {
		c.position.X--
		c.moved = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.GamepadAxis(0, 0) >= 0.5 {
		c.position.X++
		c.moved = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.GamepadAxis(0, 1) <= -0.5 {
		c.position.Y--
		c.moved = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.GamepadAxis(0, 1) >= 0.5 {
		c.position.Y++
		c.moved = true
	}

	if c.moved {
		c.animation.AddStep(1)
	}
}

// Draw draws the character image.
func (c *Character) Draw(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.position.X), float64(c.position.Y))
	return screen.DrawImage(c.animation.GetCurrentFrame(), op)
}
