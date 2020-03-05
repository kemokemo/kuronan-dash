package music

import (
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
)

func loadPlayer(b []byte) (*audio.Player, error) {
	m, err := mp3.Decode(AudioContext, audio.BytesReadSeekCloser(b))
	if err != nil {
		return nil, err
	}
	s := audio.NewInfiniteLoop(m, m.Length())
	return audio.NewPlayer(AudioContext, s)
}
