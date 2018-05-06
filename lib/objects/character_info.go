package objects

import (
	"github.com/hajimehoshi/ebiten"
)

// CharacterInfo describes the information of a character.
// This will be displayed the selecting characters scene.
type CharacterInfo struct {
	Type        CharacterType
	MainImage   *ebiten.Image
	Description string
}

// CharacterInfoMap is the map of all CharacterInfo items.
var CharacterInfoMap map[CharacterType]*CharacterInfo

// NewCharacterInfo returns a new CharacterInfo regarding args.
func NewCharacterInfo(ct CharacterType) (CharacterInfo, error) {
	var err error
	ci := CharacterInfo{}
	ci.Type = ct
	ci.MainImage, err = getMainImage(ct)
	if err != nil {
		return ci, err
	}
	ci.Description = getCharacterDescription(ct)
	return ci, nil
}
