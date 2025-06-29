package fonts

import (
	"bytes"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	// GamerFontSS is the very small size font of the strong gamer font.
	GamerFontSS *text.GoTextFace

	// GamerFontS is the small size font of the strong gamer font.
	GamerFontS *text.GoTextFace

	// GamerFontM is the middle size font of the strong gamer font.
	GamerFontM *text.GoTextFace

	// GamerFontL is the large size font of the strong gamer font.
	GamerFontL *text.GoTextFace

	// GamerFontLL is the extra large size font of the strong gamer font.
	GamerFontLL *text.GoTextFace
)

// LoadFonts loads the fonts data.
func LoadFonts() error {
	fontSource, err := text.NewGoTextFaceSource(bytes.NewReader(the_strong_gamer_ttf))
	if err != nil {
		return fmt.Errorf("failed to load font, %v", err)
	}

	GamerFontSS = &text.GoTextFace{
		Source: fontSource,
		Size:   12,
	}
	GamerFontS = &text.GoTextFace{
		Source: fontSource,
		Size:   16,
	}
	GamerFontM = &text.GoTextFace{
		Source: fontSource,
		Size:   24,
	}
	GamerFontL = &text.GoTextFace{
		Source: fontSource,
		Size:   36,
	}
	GamerFontLL = &text.GoTextFace{
		Source: fontSource,
		Size:   60,
	}
	return nil
}
