package images

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

func loadBackground() error {
	var err error

	TitleBackground, err = loadSingleImage(bg_title_png)
	if err != nil {
		return err
	}
	SelectBackground, err = loadSingleImage(bg_select_png)
	if err != nil {
		return err
	}
	SkyBackground, err = loadSingleImage(bg_prairie_png)
	if err != nil {
		return err
	}
	Curtain, err = loadSingleImage(curtain_png)
	if err != nil {
		return err
	}

	PauseLayer = ebiten.NewImage(view.ScreenWidth, view.ScreenHeight)
	PauseLayer.Fill(color.RGBA{R: 40, G: 40, B: 40, A: 200})

	return nil
}
