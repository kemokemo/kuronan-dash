package music

import (
	"github.com/hajimehoshi/ebiten/audio"
)

// SePlayer is a player to play a sound effect.
type SePlayer struct {
	name   string
	player *audio.Player
}

// NewSePlayer returns a SePlayer.
// Please call the Close method when you no longer use this instance.
func NewSePlayer(st SeType) (*SePlayer, error) {
	s := SePlayer{
		name: getSeName(st),
	}
	var err error
	s.player, err = getSePlayer(st)
	return &s, err
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

// Close closes the inner resources.
func (s *SePlayer) Close() error {
	return s.player.Close()
}
