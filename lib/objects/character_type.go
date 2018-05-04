package objects

// CharacterType describes the type of the user character.
type CharacterType int

const (
	// Kurona is 黒菜さん
	Kurona CharacterType = iota
	// Koma is 独楽ちゃん
	Koma
	// Shishimaru is 獅子丸きゅん
	Shishimaru
)

// CharacterTypeList is the list of all CharacterType items.
var CharacterTypeList = []CharacterType{Kurona, Koma, Shishimaru}
