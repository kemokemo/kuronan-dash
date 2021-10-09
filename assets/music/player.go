package music

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

func loadPlayer(b []byte) (*audio.Player, error) {
	m, err := mp3.DecodeWithSampleRate(SampleRate, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	s := audio.NewInfiniteLoop(m, m.Length())
	return AudioContext.NewPlayer(s)
}
