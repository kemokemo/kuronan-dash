package fonts

import (
	"fmt"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

var (
	// GamerFontS is the small size font of the strong gamer font.
	GamerFontS font.Face

	// GamerFontM is the middle size font of the strong gamer font.
	GamerFontM font.Face

	// GamerFontL is the large size font of the strong gamer font.
	GamerFontL font.Face

	// GamerFontLL is the extra large size font of the strong gamer font.
	GamerFontLL font.Face
)

// LoadFonts loads the fonts data.
func LoadFonts() error {
	tt, err := truetype.Parse(the_strong_gamer_ttf)
	if err != nil {
		return fmt.Errorf("failed to load TheStrongGamer font,%s", err)
	}

	const dpi = 72
	GamerFontS = truetype.NewFace(tt, &truetype.Options{
		Size:    16,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	GamerFontM = truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	GamerFontL = truetype.NewFace(tt, &truetype.Options{
		Size:    36,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	GamerFontLL = truetype.NewFace(tt, &truetype.Options{
		Size:    60,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	return nil
}
