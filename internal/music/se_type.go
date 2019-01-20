package music

type SeType int

const (
	// SE_Jump is the sound for jumping
	Se_Jump SeType = iota
)

func getSeName(st SeType) string {
	switch st {
	case Se_Jump:
		return "ジャンプ"
	default:
		return "-"
	}
}
