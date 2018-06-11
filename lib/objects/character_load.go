package objects

import (
	"bytes"
	"fmt"
	"image"

	_ "image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/kuronan-dash/assets/images"
)

func getMainImage(ct CharacterType) (*ebiten.Image, error) {
	b, err := getByteOfMainImages(ct)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

func getByteOfMainImages(ct CharacterType) ([]byte, error) {
	switch ct {
	case Kurona:
		return images.Kurona_taiki_png, nil
	case Koma:
		return images.Koma_taiki_png, nil
	case Shishimaru:
		return images.Shishimaru_taiki_png, nil
	default:
		return nil, fmt.Errorf("CharacterType %v is unknown", ct)
	}
}

func getAnimationFrames(ct CharacterType) ([]*ebiten.Image, error) {
	frames := []*ebiten.Image{}

	readers := getBytesReaderOfAnimation(ct)
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

func getBytesReaderOfAnimation(ct CharacterType) []*bytes.Reader {
	switch ct {
	case Kurona:
		return []*bytes.Reader{
			bytes.NewReader(images.Kurona_00_png),
			bytes.NewReader(images.Kurona_01_png),
			bytes.NewReader(images.Kurona_02_png),
		}
	case Koma:
		return []*bytes.Reader{
			bytes.NewReader(images.Koma_00_png),
			bytes.NewReader(images.Koma_01_png),
			bytes.NewReader(images.Koma_02_png),
			bytes.NewReader(images.Koma_03_png),
		}
	case Shishimaru:
		return []*bytes.Reader{
			bytes.NewReader(images.Shishimaru_00_png),
			bytes.NewReader(images.Shishimaru_01_png),
			bytes.NewReader(images.Shishimaru_02_png),
			bytes.NewReader(images.Shishimaru_03_png),
			bytes.NewReader(images.Shishimaru_04_png),
			bytes.NewReader(images.Shishimaru_05_png),
		}
	default:
		return nil
	}
}
