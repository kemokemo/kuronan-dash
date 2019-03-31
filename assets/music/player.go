package music

import (
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
)

func loadPlayer(b []byte) (*audio.Player, error) {
	s, err := mp3.Decode(AudioContext, audio.BytesReadSeekCloser(b))
	if err != nil {
		return nil, err
	}
	return audio.NewPlayer(AudioContext, s)
}
