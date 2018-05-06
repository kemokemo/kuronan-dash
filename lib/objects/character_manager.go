package objects

import (
	"fmt"
)

// CharacterManager manages the characters.
type CharacterManager struct {
	charaMap map[CharacterType]*Character
	infoMap  map[CharacterType]*CharacterInfo
	selected CharacterType
}

// NewCharacterManager returns the new created CharacterManager.
// Please call the Close method when you no longer use this instance.
func NewCharacterManager() (*CharacterManager, error) {
	cm := CharacterManager{}
	cm.charaMap = make(map[CharacterType]*Character)
	cm.infoMap = make(map[CharacterType]*CharacterInfo)
	cm.selected = Kurona

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
		cm.infoMap[cType] = &ci
	}

	return &cm, nil
}

// GetCharacterInfoMap returns the character info map of the available characters.
func (cm *CharacterManager) GetCharacterInfoMap() map[CharacterType]*CharacterInfo {
	infoMap := make(map[CharacterType]*CharacterInfo, len(cm.infoMap))
	for cType := range cm.infoMap {
		infoMap[cType] = cm.infoMap[cType]
	}
	return infoMap
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

// Close closes inner resources.
func (cm *CharacterManager) Close() error {
	var err, e error
	for i := range cm.charaMap {
		e = cm.charaMap[i].Close()
		if e != nil {
			err = fmt.Errorf("%v %v", err, e)
		}
	}
	return err
}
