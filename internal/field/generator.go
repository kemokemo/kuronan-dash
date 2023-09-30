package field

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// genPosFunc generates the positions to place objects.
type genPosFunc func(height int, laneHeights []float64, g genPosSet) []*view.Vector

var fieldRand *rand.Rand

// genPosField generates the positions of objects to be placed on the field.
func genPosField(height int, laneHeights []float64, g genPosSet) []*view.Vector {
	var points []*view.Vector

	fieldRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, h := range laneHeights {
		for index := 0; index < g.amount; index++ {
			r := fieldRand.Float64()
			pos := &view.Vector{
				X: float64((index+1)*g.randomRough) + float64(g.randomFine)*(1-r),
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

	fieldRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	var upperH float64
	polarity := -1.0
	for _, h := range laneHeights {
		for index := 0; index < g.amount; index++ {
			r := rand.Float64()
			pos := &view.Vector{
				X: float64(g.randomRough)*(1.0-r) + float64(g.randomFine)*r,
				Y: h - (h-upperH)/2 + 45.0*r*polarity,
			}
			points = append(points, pos)
			polarity *= -1.0
		}
		upperH = h
	}
	return points
}

// genParts generates scrollable objects.
func genParts(img *ebiten.Image, laneHeights []float64, gpf genPosFunc, gps genPosSet, kv float64) []*Parts {
	var array []*Parts

	hP := img.Bounds().Dy()
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

	hP := img.Bounds().Dy()
	points := gpf(hP, laneHeights, gps)
	for _, point := range points {
		oni := &Onigiri{}
		oni.Initialize(img, point, kv)
		array = append(array, oni)
	}
	return array
}

func genYakiManjuu(img *ebiten.Image, laneHeights []float64, gpf genPosFunc, gps genPosSet, kv float64) []*YakiManjuu {
	var array []*YakiManjuu

	hP := img.Bounds().Dy()
	points := gpf(hP, laneHeights, gps)
	for _, point := range points {
		manjuu := &YakiManjuu{}
		manjuu.Initialize(img, point, kv)
		array = append(array, manjuu)
	}
	return array
}

func genIkariYaki(img *ebiten.Image, laneHeights []float64, gpf genPosFunc, gps genPosSet, kv float64) []*IkariYaki {
	var array []*IkariYaki

	hP := img.Bounds().Dy()
	points := gpf(hP, laneHeights, gps)
	for _, point := range points {
		ikari := &IkariYaki{}
		ikari.Initialize(img, point, kv)
		array = append(array, ikari)
	}
	return array
}

// genRocks generates scrollable objects.
func genRocks(img *ebiten.Image, laneHeights []float64, gpf genPosFunc, gps genPosSet, kv float64) []*Rock {
	var array []*Rock

	hP := img.Bounds().Dy()
	points := gpf(hP, laneHeights, gps)
	for _, point := range points {
		r := &Rock{}
		r.Initialize(img, point, kv)
		array = append(array, r)
	}
	return array
}

func randBool() bool {
	fieldRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(2) == 0
}
