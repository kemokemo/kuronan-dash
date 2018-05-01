package objects

import (
	"bytes"
	"image"

	_ "image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/kuronan-dash/assets/images"
)

func getMainImage(cType CharacterType) (*ebiten.Image, error) {
	// TODO: return unique *ebiten.Image regarding the cType
	img, _, err := image.Decode(bytes.NewReader(images.Koma_taiki_png))
	if err != nil {
		return nil, err
	}
	mainImage, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	return mainImage, nil
}

func getAnimationFrames(cType CharacterType) ([]*ebiten.Image, error) {
	// TODO: return unique []*ebiten.Image regarding the cType
	frames := []*ebiten.Image{}
	readers := []*bytes.Reader{
		bytes.NewReader(images.Koma_00_png),
		bytes.NewReader(images.Koma_01_png),
		bytes.NewReader(images.Koma_02_png),
		bytes.NewReader(images.Koma_03_png),
	}

	for index := range readers {
		img, _, err := image.Decode(readers[index])
		if err != nil {
			return nil, err
		}
		frame, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
		if err != nil {
			return nil, err
		}
		frames = append(frames, frame)
	}

	return frames, nil
}
