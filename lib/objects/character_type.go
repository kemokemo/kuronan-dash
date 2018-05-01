package objects

// CharacterType describes the type of the user character.
type CharacterType int

const (
	kurona CharacterType = iota
	koma
	shishimaru
)

// CharacterTypeList is the list of all CharacterType items.
var CharacterTypeList = []CharacterType{kurona, koma, shishimaru}
