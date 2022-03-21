package images

import (
	"bytes"
	"image"
	_ "image/png" // to load png images

	"github.com/hajimehoshi/ebiten/v2"
)

func loadCharacters() error {
	var err error

	KuronaStanding, err = loadSingleImage(kurona_taiki_png)
	if err != nil {
		return err
	}
	KomaStanding, err = loadSingleImage(koma_taiki_png)
	if err != nil {
		return err
	}
	ShishimaruStanding, err = loadSingleImage(shishimaru_taiki_png)
	if err != nil {
		return err
	}

	AttackScratch, err = loadSingleImage(scratch_png)
	if err != nil {
		return err
	}
	AttackFist, err = loadSingleImage(fist_png)
	if err != nil {
		return err
	}

	return nil
}

func loadSingleImage(b []byte) (*ebiten.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(img), nil
}
