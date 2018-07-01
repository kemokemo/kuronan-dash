package music

// DiscType describes the type of the music disc.
type DiscType int

const (
	// Title will be played when the title scene.
	Title DiscType = iota
	// Stage01 will be played when the stage 01 scene.
	Stage01
)

// DiscTypeList is the list of all DiscType items.
var DiscTypeList = []DiscType{
	Title,
	Stage01,
}

func getMusicName(dt DiscType) string {
	switch dt {
	case Title:
		return "渋垣の黒猫"
	case Stage01:
		return "走れ！黒菜！"
	default:
		return "-"
	}
}
