package music

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
	"github.com/kemokemo/kuronan-dash/assets/se"
)

func getSePlayer(st SeType) (*audio.Player, error) {
	b, err := getSeByteData(st)
	if err != nil {
		return nil, err
	}
	s, err := wav.Decode(audioContext, audio.BytesReadSeekCloser(b))
	if err != nil {
		return nil, err
	}
	return audio.NewPlayer(audioContext, s)
}

func getSeByteData(st SeType) ([]byte, error) {
	switch st {
	case Se_Jump:
		return se.Jump_wav, nil
	default:
		return nil, fmt.Errorf("unknown se type %v", st)
	}
}
