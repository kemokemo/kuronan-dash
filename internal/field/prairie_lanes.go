package field

import (
	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

const (
	pLaneHeight1 = 200.0
	pLaneHeight2 = pLaneHeight1 + 170.0
	pLaneHeight3 = pLaneHeight2 + 170.0
)

func newPrairieLanesInfo() []ScrollInfo {
	w, _ := images.TilePrairie.Size()
	width := float64(w)

	infos := []ScrollInfo{
		{images.TilePrairie, &view.Vector{X: 0.0, Y: pLaneHeight1}, &view.Vector{X: 0.0, Y: 0.0}},
		{images.TilePrairie, &view.Vector{X: width, Y: pLaneHeight1}, &view.Vector{X: 0.0, Y: 0.0}},
		{images.TilePrairie, &view.Vector{X: 2.0 * width, Y: pLaneHeight1}, &view.Vector{X: 0.0, Y: 0.0}},

		{images.TilePrairie, &view.Vector{X: 0.0, Y: pLaneHeight2}, &view.Vector{X: 0.0, Y: 0.0}},
		{images.TilePrairie, &view.Vector{X: width, Y: pLaneHeight2}, &view.Vector{X: 0.0, Y: 0.0}},
		{images.TilePrairie, &view.Vector{X: 2.0 * width, Y: pLaneHeight2}, &view.Vector{X: 0.0, Y: 0.0}},

		{images.TilePrairie, &view.Vector{X: 0.0, Y: pLaneHeight3}, &view.Vector{X: 0.0, Y: 0.0}},
		{images.TilePrairie, &view.Vector{X: width, Y: pLaneHeight3}, &view.Vector{X: 0.0, Y: 0.0}},
		{images.TilePrairie, &view.Vector{X: 2.0 * width, Y: pLaneHeight3}, &view.Vector{X: 0.0, Y: 0.0}},
	}

	return infos
}
