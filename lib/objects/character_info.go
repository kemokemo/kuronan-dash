package objects

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
)

// CharacterInfo describes the information of a character.
// This will be displayed the selecting characters scene.
type CharacterInfo struct {
	Type        CharacterType
	MainImage   *ebiten.Image
	Description string
}

// NewCharacterInfo returns a new CharacterInfo regarding args.
func NewCharacterInfo(cType CharacterType) (CharacterInfo, error) {
	ci := CharacterInfo{}
	ci.Type = cType
	// TODO: create the MainImage and Description regarding the cType
	img, err := getMainImage(cType)
	if err != nil {
		return ci, err
	}
	ci.MainImage = img
	ci.Description = fmt.Sprintf("キャラクター番号: %v", cType)
	return ci, nil
}
