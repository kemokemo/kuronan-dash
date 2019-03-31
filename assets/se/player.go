package se

import (
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
	"github.com/kemokemo/kuronan-dash/assets/music"
)

// Player is a player to play a sound effect.
type Player struct {
	player *audio.Player
}

// Play plays a sound effect
func (p *Player) Play() error {
	if !p.player.IsPlaying() {
		err := p.player.Rewind()
		if err != nil {
			return err
		}
		return p.player.Play()
	}
	return nil
}

// Close closes the inner resources.
func (p *Player) Close() error {
	return p.player.Close()
}

func loadPlayer(b []byte) (*Player, error) {
	s, err := wav.Decode(music.AudioContext, audio.BytesReadSeekCloser(b))
	if err != nil {
		return nil, err
	}

	p, err := audio.NewPlayer(music.AudioContext, s)
	if err != nil {
		return nil, err
	}

	return &Player{p}, nil
}
