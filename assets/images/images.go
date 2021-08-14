package images

import (
	_ "image/png" // to load png images

	"github.com/hajimehoshi/ebiten/v2"
)

// background
var (
	TitleBackground  *ebiten.Image
	SelectBackground *ebiten.Image
	SkyBackground    *ebiten.Image
)

// field parts
//  - prairie  field
var (
	TilePrairie  *ebiten.Image
	Grass1       *ebiten.Image
	Grass2       *ebiten.Image
	Grass3       *ebiten.Image
	MountainNear *ebiten.Image
	MountainFar  *ebiten.Image
	CloudNear    *ebiten.Image
	CloudFar     *ebiten.Image
	RockNormal   *ebiten.Image
)

// Foods
var (
	Onigiri *ebiten.Image
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
