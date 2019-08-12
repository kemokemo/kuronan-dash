package images

import (
	_ "image/png" // to load png images

	"github.com/hajimehoshi/ebiten"
)

// background
var (
	TitleBackground  *ebiten.Image
	SelectBackground *ebiten.Image
	SkyBackground    *ebiten.Image
)

// field parts
var (
	TilePrairie  *ebiten.Image
	MountainNear *ebiten.Image
	MountainFar  *ebiten.Image
	Cloud        *ebiten.Image
	Cloud2       *ebiten.Image
)

// character standing image
var (
	KuronaStanding     *ebiten.Image
	KomaStanding       *ebiten.Image
	ShishimaruStanding *ebiten.Image
)

// character animation
var (
	KuronaAnimation     []*ebiten.Image
	KomaAnimation       []*ebiten.Image
	ShishimaruAnimation []*ebiten.Image
)

// LoadImages loads all public images.
func LoadImages() error {
	err := loadBackground()
	if err != nil {
		return err
	}

	err = loadFieldparts()
	if err != nil {
		return err
	}

	err = loadStandingImages()
	if err != nil {
		return err
	}

	err = loadAnimation()
	if err != nil {
		return err
	}

	return nil
}
