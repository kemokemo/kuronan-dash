package objects

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/kuronan-dash/lib/music"
)

// Position describes the position by x and y.
type Position struct {
	X int
	Y int
}

// Character describes a character.
type Character struct {
	animation StepAnimation
	Position  Position
	moved     bool
	state     CharacterState
	jumpSe    *music.SePlayer
}

// NewCharacter creates a new character instance.
// Please call the Close method when you no longer use this instance.
func NewCharacter(ct CharacterType) (*Character, error) {
	frames, err := getAnimationFrames(ct)
	if err != nil {
		return nil, err
	}
	c := &Character{
		animation: StepAnimation{
			Frames:        frames,
			DurationSteps: 5,
		},
	}
	err = c.animation.Init()
	if err != nil {
		return nil, err
	}
	c.jumpSe, err = music.NewSePlayer("_assets/se/jump.wav")
	if err != nil {
		return nil, err
	}
	return c, nil
}

// SetInitialPosition sets the initial position for this character.
func (c *Character) SetInitialPosition(pos Position) {
	c.Position = pos
}

// Move moves the character regarding the user input.
func (c *Character) Move() {
	c.moved = false
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.GamepadAxis(0, 0) <= -0.5 {
		c.Position.X--
		c.moved = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.GamepadAxis(0, 0) >= 0.5 {
		c.Position.X++
		c.moved = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.GamepadAxis(0, 1) <= -0.5 {
		c.Position.Y--
		c.moved = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.GamepadAxis(0, 1) >= 0.5 {
		c.Position.Y++
		c.moved = true
	}

	if c.moved {
		c.animation.AddStep(1)
	}
}

// Draw draws the character image.
func (c *Character) Draw(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.Position.X), float64(c.Position.Y))
	return screen.DrawImage(c.animation.GetCurrentFrame(), op)
}

// SetState sets the status for this character.
func (c *Character) SetState(state CharacterState) {
	c.state = state
}

// PlaySe plays a sound effect according to the status of this character.
func (c *Character) PlaySe() error {
	// TODO: ステータスに応じたSEを再生
	if c.state == Ascending {
		// TODO: このままだとジャンプ中SE再生し続けるので対策が必要
		return c.jumpSe.Play()
	}
	return nil
}

// Close closes the inner resources.
func (c *Character) Close() error {
	return c.jumpSe.Close()
}
