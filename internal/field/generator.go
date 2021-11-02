package field

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// genPosFunc generates the positions to place objects.
type genPosFunc func(height int, laneHeights []float64, g genPosSet) []*view.Vector

// genPosField generates the positions of objects to be placed on the field.
func genPosField(height int, laneHeights []float64, g genPosSet) []*view.Vector {
	var points []*view.Vector

	rand.Seed(time.Now().UnixNano())
	for _, h := range laneHeights {
		for index := 0; index < g.amount; index++ {
			r := rand.Float64()
			pos := &view.Vector{
				X: float64((index+1)*g.randomRough) + float64(g.randomFine)*r,
				Y: h - float64(height-1),
			}
			points = append(points, pos)
		}
	}
	return points
}

// genPosAir generates the positions of objects to be placed in the air.
func genPosAir(h int, laneHeights []float64, g genPosSet) []*view.Vector {
	var points []*view.Vector

	rand.Seed(time.Now().UnixNano())
	for _, h := range laneHeights {
		for index := 0; index < g.amount; index++ {
			r := rand.Float64()
			pos := &view.Vector{
				X: float64(g.randomRough) + float64(g.randomFine)*r,
				Y: h - 40.0 - 100.0*r - float64(h/2),
			}
			points = append(points, pos)
		}
	}
	return points
}

// genParts generates scrollable objects.
func genParts(img *ebiten.Image, laneHeights []float64, gpf genPosFunc, gps genPosSet, kv float64) []*Parts {
	var array []*Parts

	_, hP := img.Size()
	points := gpf(hP, laneHeights, gps)
	for _, point := range points {
		fp := &Parts{}
		fp.Initialize(img, point, kv)
		array = append(array, fp)
	}
	return array
}

// genOnigiri generates onigiri.
func genOnigiri(img *ebiten.Image, laneHeights []float64, gpf genPosFunc, gps genPosSet, kv float64) []*Onigiri {
	var array []*Onigiri

	_, hP := img.Size()
	points := gpf(hP, laneHeights, gps)
	for _, point := range points {
		oni := &Onigiri{}
		oni.Initialize(img, point, kv)
		array = append(array, oni)
	}
	return array
}

// genRocks generates scrollable objects.
func genRocks(img *ebiten.Image, laneHeights []float64, gpf genPosFunc, gps genPosSet, kv float64) []*Rock {
	var array []*Rock

	_, hP := img.Size()
	points := gpf(hP, laneHeights, gps)
	for _, point := range points {
		r := &Rock{}
		r.Initialize(img, point, kv)
		array = append(array, r)
	}
	return array
}

func randBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 0
}
