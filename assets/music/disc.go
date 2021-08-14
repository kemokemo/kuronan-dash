package music

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
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
func (d *Disc) Play() {
	d.player.Play()
}

// Pause pauses music.
func (d *Disc) Pause() {
	d.player.Pause()
}

// Stop stops music. (pause and rewind)
func (d *Disc) Stop() error {
	d.player.Pause()
	return d.player.Rewind()
}
