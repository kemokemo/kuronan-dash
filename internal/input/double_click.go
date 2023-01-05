package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type DoubleClick struct {
	mBtn       ebiten.MouseButton
	counter    int
	waitPeriod int
	firstClick bool
}

func NewDoubleClick(mBtn ebiten.MouseButton) *DoubleClick {
	return &DoubleClick{
		mBtn:       mBtn,
		counter:    0,
		waitPeriod: 10,
	}
}

func (d *DoubleClick) Triggered() bool {
	if !d.firstClick {
		d.firstClick = inpututil.IsMouseButtonJustPressed(d.mBtn)
	} else {
		d.counter++
		if d.counter <= d.waitPeriod {
			if inpututil.IsMouseButtonJustPressed(d.mBtn) {
				d.counter = 0
				d.firstClick = false
				return true
			}
		} else {
			d.counter = 0
			d.firstClick = false
		}
	}

	return false
}
