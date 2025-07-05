package input

import (
	vpad "github.com/kemokemo/ebiten-virtualpad"
	"github.com/kemokemo/kuronan-dash/assets"
)

func NewVolumeChecker(btn vpad.SelectButton, defaultVal bool, ch chan<- assets.GameSoundControl) VolumeChecker {
	btn.SetSelectState(defaultVal)
	return &VolumeInputChecker{
		VolumeBtn:          btn,
		current:            defaultVal,
		previous:           defaultVal,
		gameSoundControlCh: ch,
	}
}

type VolumeInputChecker struct {
	VolumeBtn          vpad.SelectButton
	current            bool
	previous           bool
	gameSoundControlCh chan<- assets.GameSoundControl
}

func (i *VolumeInputChecker) Update() {
	i.VolumeBtn.Update()

	i.previous = i.current
	i.current = i.VolumeBtn.IsSelected()

	if !i.previous && i.current {
		i.gameSoundControlCh <- assets.SoundOn
	} else if i.previous && !i.current {
		i.gameSoundControlCh <- assets.SoundOff
	}
}

func (i *VolumeInputChecker) IsVolumeOn() bool {
	return i.VolumeBtn.IsSelected()
}
