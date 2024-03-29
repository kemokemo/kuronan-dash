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
	AttackKomaFist, err = loadSingleImage(koma_fist_png)
	if err != nil {
		return err
	}
	AttackShishimaruFist, err = loadSingleImage(shishimaru_fist_png)
	if err != nil {
		return err
	}
	KuronaSpEffect, err = loadSingleImage(kurona_sp_effect_png)
	if err != nil {
		return err
	}
	KomaSpEffect, err = loadSingleImage(koma_sp_effect_png)
	if err != nil {
		return err
	}
	ShishimaruSpEffect, err = loadSingleImage(shishimaru_sp_effect_png)
	if err != nil {
		return err
	}
	KuronaSpBack, err = loadSingleImage(kurona_sp_back_png)
	if err != nil {
		return err
	}
	KomaSpBack, err = loadSingleImage(koma_sp_back_png)
	if err != nil {
		return err
	}
	ShishimaruSpBack, err = loadSingleImage(shishimaru_sp_back_png)
	if err != nil {
		return err
	}
	KuronaMapIcon, err = loadSingleImage(kurona_map_icon_png)
	if err != nil {
		return err
	}
	KomaMapIcon, err = loadSingleImage(koma_map_icon_png)
	if err != nil {
		return err
	}
	ShishimaruMapIcon, err = loadSingleImage(shishimaru_map_icon_png)
	if err != nil {
		return err
	}
	SpecialReadyIcon, err = loadSingleImage(sp_charge_icon_png)
	if err != nil {
		return err
	}
	WalkStateIcon, err = loadSingleImage(walk_icon_png)
	if err != nil {
		return err
	}
	StaminaEmptyIcon, err = loadSingleImage(stamina_empty_png)
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
