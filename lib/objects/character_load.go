package objects

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func getMainImage(cType CharacterType) (*ebiten.Image, error) {
	// TODO: return unique *ebiten.Image regarding the cType
	image, _, err := ebitenutil.NewImageFromFile("assets/images/character/koma_taiki.png", ebiten.FilterNearest)
	if err != nil {
		return nil, err
	}
	return image, nil
}

func getAnimationImages(cType CharacterType) []string {
	// TODO: return unique []string regarding the cType
	return []string{
		"assets/images/character/koma_00.png",
		"assets/images/character/koma_01.png",
		"assets/images/character/koma_02.png",
		"assets/images/character/koma_03.png",
	}
}
