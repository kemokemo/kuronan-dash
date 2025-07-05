package music

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

// Disc is a music player
type Disc struct {
	Name    string
	player  *audio.Player
	disable bool
}

// Close closes inner resources.
func (d *Disc) Close() error {
	return d.player.Close()
}

func (d *Disc) SetVolume(volume float64) {
	d.player.SetVolume(volume)
}

// Play plays a preselected disc.
func (d *Disc) Play() {
	if d.disable {
		return
	}
	if d.player.IsPlaying() {
		return
	}
	d.player.Play()
}

// Pause pauses music.
func (d *Disc) Pause() {
	d.player.Pause()
}

// Stop stops music. (pause and rewind)
func (d *Disc) Stop() {
	d.player.Pause()
	err := d.player.Rewind()
	if err != nil {
		log.Println("failed to stop disc, disc name is ", d.Name)
	}
}

func (d *Disc) SetVolumeFlag(isVolumeOn bool) {
	d.disable = !isVolumeOn
	if d.disable {
		d.player.SetVolume(0)
		d.Pause()
	} else {
		d.player.SetVolume(0.5)
	}
}
