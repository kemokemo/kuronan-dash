package assets

import (
	"fmt"

	"github.com/kemokemo/kuronan-dash/assets/fonts"
	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/assets/music"
	"github.com/kemokemo/kuronan-dash/assets/se"
)

// LoadAssets loads all assets.
// Please call this before using assets.
func LoadAssets() error {
	err := images.LoadImages()
	if err != nil {
		return err
	}

	// before loading music and sound effects, need to load audio context.
	music.LoadAudioContext()

	err = music.LoadMusic()
	if err != nil {
		return err
	}

	err = se.LoadSE()
	if err != nil {
		return err
	}

	err = fonts.LoadFonts()
	if err != nil {
		return err
	}

	return nil
}

// CloseAssets closes all assets.
func CloseAssets() error {
	var e error

	err := music.CloseMusic()
	if err != nil {
		e = fmt.Errorf("%v", err)
	}
	err = se.CloseSE()
	if err != nil {
		e = fmt.Errorf("%v:%v", e, err)
	}

	return e
}
