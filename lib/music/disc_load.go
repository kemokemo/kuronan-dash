package music

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"github.com/kemokemo/kuronan-dash/assets/audios"
)

func getPlayer(dt DiscType) (*audio.Player, error) {
	b, err := getByteData(dt)
	if err != nil {
		return nil, err
	}
	s, err := mp3.Decode(audioContext, audio.BytesReadSeekCloser(b))
	if err != nil {
		return nil, err
	}
	return audio.NewPlayer(audioContext, s)
}

func getByteData(dt DiscType) ([]byte, error) {
	switch dt {
	case Title:
		return audios.Hashire_kurona_mp3, nil
	case Stage01:
		return audios.Shibugaki_no_kuroneko_mp3, nil
	default:
		return nil, fmt.Errorf("unknown disc type %v", dt)
	}
}
