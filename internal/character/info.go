package character

import (
	"github.com/hajimehoshi/ebiten"
)

// Info describes the information of a character.
// This will be displayed the selecting characters scene.
type Info struct {
	ID          ID
	MainImage   *ebiten.Image
	Description string
}

// CharaInfoMap is the map of all CharacterInfo items.
var CharaInfoMap map[ID]*Info

// NewCharacterInfo returns a new CharacterInfo regarding args.
func NewCharacterInfo(id ID) (Info, error) {
	var err error
	ci := Info{}
	ci.ID = id
	ci.MainImage, err = getMainImage(id)
	if err != nil {
		return ci, err
	}
	ci.Description = getDescription(id)
	return ci, nil
}
