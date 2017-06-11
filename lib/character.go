package kuronandash

import "github.com/hajimehoshi/ebiten"

// Position describes the position by x and y.
type Position struct {
	X int
	Y int
}

// Character describes a charactor.
type Character struct {
	ImagesPaths []string

	animation Animation
	stepCount int
	position  Position
}

// Init loads asset files.
func (c *Character) Init() error {
	c.animation = Animation{
		ImagesPaths:  c.ImagesPaths,
		DurationStep: 5,
	}
	err := c.animation.Init()
	if err != nil {
		return err
	}
	return nil
}

// SetInitialPosition sets the initial position for this character.
func (c *Character) SetInitialPosition(pos Position) {
	c.position = pos
}

// Move moves the charactor regarding the user input.
func (c *Character) Move() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.GamepadAxis(0, 0) <= -0.5 {
		c.position.X--
		c.stepCount++
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.GamepadAxis(0, 0) >= 0.5 {
		c.position.X++
		c.stepCount++
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.GamepadAxis(0, 1) <= -0.5 {
		c.position.Y--
		c.stepCount++
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.GamepadAxis(0, 1) >= 0.5 {
		c.position.Y++
		c.stepCount++
	}
}

// Draw draws the character image.
func (c *Character) Draw(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.position.X), float64(c.position.Y))
	return screen.DrawImage(c.animation.GetCurrentFrame(c.getDeltaStepCount()), op)
}

func (c *Character) getDeltaStepCount() int {
	d := c.stepCount
	c.stepCount = 0
	return d
}
