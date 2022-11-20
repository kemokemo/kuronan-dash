package input

import (
	vpad "github.com/kemokemo/ebiten-virtualpad"
)

func NewVolumeChecker(btn vpad.SelectButton, defaultVal bool) VolumeChecker {
	btn.SetSelectState(defaultVal)
	return &VolumeInputChecker{
		VolumeBtn: btn,
		current:   defaultVal,
		previous:  defaultVal,
	}
}

type VolumeInputChecker struct {
	VolumeBtn       vpad.SelectButton
	current         bool
	previous        bool
	isJustVolumeOn  bool
	isJustVolumeOff bool
}

func (i *VolumeInputChecker) Update() {
	i.VolumeBtn.Update()

	i.previous = i.current
	i.current = i.VolumeBtn.IsSelected()

	if !i.previous && i.current {
		i.isJustVolumeOn = true
	} else {
		i.isJustVolumeOn = false
	}

	if i.previous && !i.current {
		i.isJustVolumeOff = true
	} else {
		i.isJustVolumeOff = false
	}
}

func (i *VolumeInputChecker) IsVolumeOn() bool {
	return i.VolumeBtn.IsSelected()
}

func (i *VolumeInputChecker) JustVolumeOn() bool {
	return i.isJustVolumeOn
}

func (i *VolumeInputChecker) JustVolumeOff() bool {
	return i.isJustVolumeOff
}
