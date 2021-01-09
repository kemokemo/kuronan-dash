package field

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// genPosFunc generates the positions to place objects.
type genPosFunc func(height int, g genPosSet) []*view.Vector

// genPosField generates the positions of objects to be placed on the field.
func genPosField(height int, g genPosSet) []*view.Vector {
	var points []*view.Vector

	rand.Seed(time.Now().UnixNano())
	for _, h := range LaneHeights {
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
func genPosAir(h int, g genPosSet) []*view.Vector {
	var points []*view.Vector

	rand.Seed(time.Now().UnixNano())
	for _, h := range LaneHeights {
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

// genVelFunc generates the velocity of an object.
type genVelFunc func(g genVelSet) *view.Vector

func genVel(g genVelSet) *view.Vector {
	var vel *view.Vector
	if g.random {
		rand.Seed(time.Now().UnixNano())
		r := rand.Float64()
		vel = &view.Vector{X: g.x * r, Y: g.y * r}
	} else {
		vel = &view.Vector{X: g.x, Y: g.y}
	}
	return vel
}

// genParts generates scrollable objects.
func genParts(img *ebiten.Image, gpf genPosFunc, gps genPosSet, gvf genVelFunc, gvs genVelSet) []*Parts {
	var array []*Parts

	_, hP := img.Size()
	points := gpf(hP, gps)
	for _, point := range points {
		fp := &Parts{}
		vel := gvf(gvs)
		fp.Initialize(img, point, vel)
		array = append(array, fp)
	}
	return array
}

// genOnigiri generates onigiri.
func genOnigiri(img *ebiten.Image, gpf genPosFunc, gps genPosSet, gvf genVelFunc, gvs genVelSet) []*Onigiri {
	var array []*Onigiri

	_, hP := img.Size()
	points := gpf(hP, gps)
	for _, point := range points {
		oni := &Onigiri{}
		vel := gvf(gvs)
		oni.Initialize(img, point, vel)
		array = append(array, oni)
	}
	return array
}

// genRocks generates scrollable objects.
func genRocks(img *ebiten.Image, gpf genPosFunc, gps genPosSet, gvf genVelFunc, gvs genVelSet) []*Rock {
	var array []*Rock

	_, hP := img.Size()
	points := gpf(hP, gps)
	for _, point := range points {
		r := &Rock{}
		vel := gvf(gvs)
		r.Initialize(img, point, vel)
		array = append(array, r)
	}
	return array
}

func randBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 0
}
