package objects

import (
	"fmt"
)

// CharacterManager manages the characters.
type CharacterManager struct {
	charaMap map[CharacterType]*Character
	infoMap  map[CharacterType]CharacterInfo
	selected CharacterType
}

// NewCharacterManager returns the new created CharacterManager.
func NewCharacterManager() (*CharacterManager, error) {
	cm := CharacterManager{}
	cm.charaMap = make(map[CharacterType]*Character)
	cm.infoMap = make(map[CharacterType]CharacterInfo)
	cm.selected = kurona

	for _, cType := range CharacterTypeList {
		c, err := NewCharacter(cType)
		if err != nil {
			return &cm, err
		}
		cm.charaMap[cType] = c
	}

	for _, cType := range CharacterTypeList {
		ci, err := NewCharacterInfo(cType)
		if err != nil {
			return &cm, err
		}
		cm.infoMap[cType] = ci
	}

	return &cm, nil
}

// GetCharacterInfo returns the info of the available characters.
func (cm *CharacterManager) GetCharacterInfo() []CharacterInfo {
	infos := make([]CharacterInfo, 0, len(cm.infoMap))
	for cType := range cm.infoMap {
		infos = append(infos, cm.infoMap[cType])
	}
	return infos
}

// SelectCharacter selects a character regarding the argument.
func (cm *CharacterManager) SelectCharacter(ct CharacterType) error {
	if cm.charaMap[ct] == nil {
		return fmt.Errorf("ordered CharacterInfo is not included:%v", ct)
	}
	cm.selected = ct
	return nil
}

// GetSelectedCharacter returns the selected character.
func (cm *CharacterManager) GetSelectedCharacter() *Character {
	return cm.charaMap[cm.selected]
}
