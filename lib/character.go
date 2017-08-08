package kuronandash

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
)

// Position describes the position by x and y.
type Position struct {
	X int
	Y int
}

// Status describes the status of a character.
type Status int

const (
	// Stop describes a character is stopping
	Stop Status = iota
	// Walk describes a character is walking
	Walk
	// Dash is Dash!
	Dash
	// Ascending describes a character is jumping
	Ascending
	// Descending describes a character is descending
	Descending
)

// NewCharacter creates a new character instance.
func NewCharacter(context *audio.Context, imagePaths []string) (*Character, error) {
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
	c.jumpSe, err = NewSePlayer(context, "assets/se/jump.wav")
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
	status    Status
	jumpSe    *SePlayer
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

// SetState sets the status for this character.
func (c *Character) SetState(status Status) {
	c.status = status
}

// PlaySe plays a sound effect according to the status of this character.
func (c *Character) PlaySe() error {
	// TODO: ステータスに応じたSEを再生
	if c.status == Ascending {
		// TODO: このままだとジャンプ中SE再生し続けるので対策が必要
		return c.jumpSe.Play()
	}
	return nil
}
