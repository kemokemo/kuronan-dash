package music

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

// SampleRate is the sampling rate of music and se.
const SampleRate = 44100

// AudioContext is the context for the all audios.
var AudioContext *audio.Context

// Discs
var (
	Title   *Disc
	Select  *Disc
	Stage01 *Disc
)

// LoadAudioContext load audio context.
func LoadAudioContext() {
	AudioContext = audio.NewContext(SampleRate)
}

// LoadMusic loads all music.
func LoadMusic() error {
	var err error

	p, err := loadPlayer(shibugaki_no_kuroneko_mp3)
	if err != nil {
		return err
	}
	Title = &Disc{Name: "しぶがき の くろねこ", player: p}
	Title.SetVolume(0.6)

	p, err = loadPlayer(iza_tabidati_no_toki_mp3)
	if err != nil {
		return err
	}
	Select = &Disc{Name: "いざ たびだち の とき", player: p}
	Select.SetVolume(0.6)

	p, err = loadPlayer(hashire_kurona_mp3)
	if err != nil {
		return err
	}
	Stage01 = &Disc{Name: "はしれ! くろな!", player: p}
	Stage01.SetVolume(0.6)

	return nil
}

// CloseMusic closes all music.
func CloseMusic() error {
	var e error

	err := Title.Close()
	if err != nil {
		e = fmt.Errorf("%v", err)
	}
	err = Select.Close()
	if err != nil {
		e = fmt.Errorf("%v", err)
	}
	err = Stage01.Close()
	if err != nil {
		e = fmt.Errorf("%v:%v", e, err)
	}

	return e
}
