package music

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

const sampleRate = 44100

// AudioContext is the context for the all audios.
var AudioContext *audio.Context

// Discs
var (
	Title   *Disc
	Stage01 *Disc
)

// LoadAudioContext load audio context.
func LoadAudioContext() {
	AudioContext = audio.NewContext(sampleRate)
}

// LoadMusic loads all music.
func LoadMusic() error {
	var err error

	p, err := loadPlayer(shibugaki_no_kuroneko_mp3)
	if err != nil {
		return err
	}
	Title = &Disc{Name: "しぶがき の くろねこ", player: p}

	p, err = loadPlayer(hashire_kurona_mp3)
	if err != nil {
		return err
	}
	Stage01 = &Disc{Name: "はしれ! くろな!", player: p}

	return nil
}

// CloseMusic closes all music.
func CloseMusic() error {
	var e error

	err := Title.Close()
	if err != nil {
		e = fmt.Errorf("%v", err)
	}
	err = Stage01.Close()
	if err != nil {
		e = fmt.Errorf("%v:%v", e, err)
	}

	return e
}
