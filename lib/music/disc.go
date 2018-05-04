package music

import (
	"github.com/hajimehoshi/ebiten/audio"
)

type disc struct {
	name   string
	player *audio.Player
}

func (d *disc) close() error {
	return d.player.Close()
}

func newDisc(dt DiscType) (*disc, error) {
	d := disc{
		name: getMusicName(dt),
	}
	var err error
	d.player, err = getMusicPlayer(dt)
	return &d, err
}
