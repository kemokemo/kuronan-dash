package images

import (
	"bytes"
	"image"
	_ "image/png" // to load png images

	"github.com/hajimehoshi/ebiten"
)

// frames for the animation
var (
	kuronaFrames = []*bytes.Reader{
		bytes.NewReader(kurona_00_png),
		bytes.NewReader(kurona_01_png),
		bytes.NewReader(kurona_02_png),
		bytes.NewReader(kurona_03_png),
		bytes.NewReader(kurona_04_png),
		bytes.NewReader(kurona_05_png),
	}

	komaFrames = []*bytes.Reader{
		bytes.NewReader(koma_00_png),
		bytes.NewReader(koma_01_png),
		bytes.NewReader(koma_02_png),
		bytes.NewReader(koma_03_png),
		bytes.NewReader(koma_04_png),
		bytes.NewReader(koma_05_png),
	}

	shishimaruFrames = []*bytes.Reader{
		bytes.NewReader(shishimaru_00_png),
		bytes.NewReader(shishimaru_01_png),
		bytes.NewReader(shishimaru_02_png),
		bytes.NewReader(shishimaru_03_png),
		bytes.NewReader(shishimaru_04_png),
		bytes.NewReader(shishimaru_05_png),
	}
)

func loadAnimation() error {
	var err error

	KuronaAnimation, err = loadFrames(kuronaFrames)
	if err != nil {
		return err
	}
	KomaAnimation, err = loadFrames(komaFrames)
	if err != nil {
		return err
	}
	ShishimaruAnimation, err = loadFrames(shishimaruFrames)
	if err != nil {
		return err
	}

	return nil
}

func loadFrames(br []*bytes.Reader) ([]*ebiten.Image, error) {
	frames := []*ebiten.Image{}

	for index := range br {
		img, _, err := image.Decode(br[index])
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
