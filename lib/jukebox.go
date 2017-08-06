package kuronandash

import (
	"os"

	"github.com/hajimehoshi/ebiten/audio"
	mp3 "github.com/hajimehoshi/go-mp3"
)

// JukeBox loads all music files and play any time you want.
type JukeBox struct {
	player *audio.Player
}

// NewJukeBox creates a new JukeBox instance.
func NewJukeBox(musicPath string) (*JukeBox, error) {
	// TODO: save music file name
	f, err := os.Open(musicPath)
	if err != nil {
		return nil, err
	}
	s, err := mp3.NewDecoder(f)
	if err != nil {
		return nil, err
	}
	c, err := audio.NewContext(s.SampleRate())
	if err != nil {
		return nil, err
	}
	p, err := audio.NewPlayer(c, s)
	if err != nil {
		return nil, err
	}

	return &JukeBox{
		player: p,
	}, nil
}

// Play plays the music file specified by the name arg.
func (j *JukeBox) Play(name string) error {
	// TODO: find the music file name
	if j.player.IsPlaying() {
		return nil
	}
	return j.player.Play()
}

// Pause pauses music.
func (j *JukeBox) Pause() error {
	if !j.player.IsPlaying() {
		return nil
	}
	return j.player.Pause()
}

// Stop stops music.
func (j *JukeBox) Stop() error {
	if !j.player.IsPlaying() {
		return nil
	}
	err := j.player.Pause()
	if err != nil {
		return err
	}
	return j.player.Rewind()
}

// Close closes all resources that JukeBox used.
func (j *JukeBox) Close() error {
	return j.player.Close()
}
