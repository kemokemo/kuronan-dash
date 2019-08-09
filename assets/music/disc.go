package music

import (
	"github.com/hajimehoshi/ebiten/audio"
)

// Disc is a music player
type Disc struct {
	Name   string
	player *audio.Player
}

// Close closes inner resources.
func (d *Disc) Close() error {
	return d.player.Close()
}

// Play plays a preselected disc.
func (d *Disc) Play() error {
	if d.player.IsPlaying() {
		return nil
	}
	err := d.player.Rewind()
	if err != nil {
		return err
	}
	err = AudioContext.Update()
	if err != nil {
		return err
	}
	return d.player.Play()
}

// Pause pauses music.
func (d *Disc) Pause() error {
	if !d.player.IsPlaying() {
		return nil
	}
	return d.player.Pause()
}

// Stop stops music. (pause and rewind)
func (d *Disc) Stop() error {
	if !d.player.IsPlaying() {
		return nil
	}
	err := d.player.Pause()
	if err != nil {
		return err
	}
	return d.player.Rewind()
}
