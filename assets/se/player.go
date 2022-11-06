package se

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/kemokemo/kuronan-dash/assets/music"
)

// Player is a player to play a sound effect.
type Player struct {
	player *audio.Player
}

// Play plays a sound effect
// If you execute this feature before playing finished, you can get the new sound from start.
func (p *Player) Play() {
	if !p.player.IsPlaying() {
		err := p.player.Rewind()
		if err != nil {
			log.Println("failed to rewind, ", err)
		}
	}
	p.player.Play()
}

func (p *Player) IsPlaying() bool {
	return p.player.IsPlaying()
}

// Close closes the inner resources.
func (p *Player) Close() error {
	return p.player.Close()
}

func loadPlayer(b []byte) (*Player, error) {
	s, err := wav.DecodeWithSampleRate(music.SampleRate, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	p, err := music.AudioContext.NewPlayer(s)
	if err != nil {
		return nil, err
	}

	return &Player{p}, nil
}
