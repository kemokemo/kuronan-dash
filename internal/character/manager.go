package character

import (
	"fmt"
)

// Manager manages the characters.
type Manager struct {
	charaMap map[ID]*Character
	infoMap  map[ID]*Info
	selected ID
}

// NewManager returns the new created CharacterManager.
// Please call the Close method when you no longer use this instance.
func NewManager() (*Manager, error) {
	cm := Manager{}
	cm.charaMap = make(map[ID]*Character)
	cm.infoMap = make(map[ID]*Info)
	cm.selected = Kurona

	for _, cType := range CharaIDs {
		c, err := NewCharacter(cType)
		if err != nil {
			return &cm, err
		}
		cm.charaMap[cType] = c
	}

	for _, cType := range CharaIDs {
		ci, err := NewCharacterInfo(cType)
		if err != nil {
			return &cm, err
		}
		cm.infoMap[cType] = &ci
	}

	return &cm, nil
}

// GetCharacterInfoMap returns the character info map of the available characters.
func (cm *Manager) GetCharacterInfoMap() map[ID]*Info {
	infoMap := make(map[ID]*Info, len(cm.infoMap))
	for cType := range cm.infoMap {
		infoMap[cType] = cm.infoMap[cType]
	}
	return infoMap
}

// SelectCharacter selects a character regarding the argument.
func (cm *Manager) SelectCharacter(id ID) error {
	if cm.charaMap[id] == nil {
		return fmt.Errorf("ordered CharacterInfo is not included:%v", id)
	}
	cm.selected = id
	return nil
}

// GetSelectedCharacter returns the selected character.
func (cm *Manager) GetSelectedCharacter() *Character {
	return cm.charaMap[cm.selected]
}

// Close closes inner resources.
func (cm *Manager) Close() error {
	var err, e error
	for i := range cm.charaMap {
		e = cm.charaMap[i].Close()
		if e != nil {
			err = fmt.Errorf("%v %v", err, e)
		}
	}
	return err
}
