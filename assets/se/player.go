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
	player  *audio.Player
	disable bool
}

// Play plays a sound effect
// If you execute this feature before playing finished, you can get the new sound from start.
func (p *Player) Play() {
	if p.disable {
		return
	}

	err := p.player.Rewind()
	if err != nil {
		log.Println("failed to rewind, ", err)
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

func (p *Player) SetVolumeFlag(isVolumeOn bool) {
	p.disable = !isVolumeOn
	if p.disable {
		p.player.SetVolume(0)
	} else {
		p.player.SetVolume(1.0)
	}
}

func (p *Player) SetVolumeValue(v float64) {
	if p.disable {
		return
	} else if v < 0 {
		p.player.SetVolume(0)
	} else if 1 < v {
		p.player.SetVolume(1)
	} else {
		p.player.SetVolume(v)
	}
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

	return &Player{player: p}, nil
}
