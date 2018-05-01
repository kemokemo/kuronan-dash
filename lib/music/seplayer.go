package music

import (
	"os"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
)

// SePlayer is a player to play a sound effect.
type SePlayer struct {
	file   *os.File
	player *audio.Player
}

// NewSePlayer returns a SePlayer.
func NewSePlayer(soundPath string) (*SePlayer, error) {
	f, err := os.Open(soundPath)
	if err != nil {
		return nil, err
	}
	s, err := wav.Decode(audioContext, f)
	if err != nil {
		return nil, err
	}
	p, err := audio.NewPlayer(audioContext, s)
	if err != nil {
		return nil, err
	}

	return &SePlayer{
		file:   f,
		player: p,
	}, nil
}

// Play plays a sound effect
func (s *SePlayer) Play() error {
	if !s.player.IsPlaying() {
		err := s.player.Rewind()
		if err != nil {
			return err
		}
		return s.player.Play()
	}
	return nil
}
