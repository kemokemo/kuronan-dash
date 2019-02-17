package character

import "fmt"

// ID describes the type of the user character.
type ID int

const (
	// Kurona is 黒菜さん
	Kurona ID = iota
	// Koma is 独楽ちゃん
	Koma
	// Shishimaru is 獅子丸きゅん
	Shishimaru
)

// CharaIDs is the list of all CharacterType items.
var CharaIDs = []ID{Kurona, Koma, Shishimaru}

func getDescription(id ID) string {
	switch id {
	case Kurona:
		return "本作の主人公。いつも元気いっぱいな渋垣のアイドル。"
	case Koma:
		return "黒菜の親友でライバル。実直な性格で特技は鉄拳制裁。可愛い。"
	case Shishimaru:
		return "優しく穏やかな独楽の弟。女子力が非常に高く、一部では正ヒロインとの呼び声も高い。"
	default:
		return fmt.Sprintf("CharacterType %v is unknown", id)
	}
}
