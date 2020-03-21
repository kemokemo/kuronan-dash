package field

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

func generatePoints(hT, hP, num, param1, param2 int) []view.Vector {
	var points []view.Vector

	rand.Seed(time.Now().UnixNano())
	for _, h := range LaneHeights {
		for index := 0; index < num; index++ {
			r := rand.Float64()
			pos := view.Vector{
				X: float64((index+1)*param1) + float64(param2)*r,
				Y: h - float64(hP-3),
			}
			points = append(points, pos)
		}
	}
	return points
}

func generateCloudPoints(hT, hP, num, param1, param2 int) []view.Vector {
	var points []view.Vector

	rand.Seed(time.Now().UnixNano())
	for _, h := range LaneHeights {
		for index := 0; index < num; index++ {
			r := rand.Float64()
			pos := view.Vector{
				X: float64(param1) + float64(param2)*r,
				Y: h - 40.0 - 100.0*r - float64(hP/2),
			}
			points = append(points, pos)
		}
	}
	return points
}

func create(img *ebiten.Image, num, param1, param2 int, vel view.Vector, moreRandom bool) []ScrollableObject {
	var array []ScrollableObject
	_, hT := images.TilePrairie.Size()

	_, hP := img.Size()
	var points []view.Vector
	if moreRandom {
		points = generateCloudPoints(hT, hP, num, param1, param2)
	} else {
		points = generatePoints(hT, hP, num, param1, param2)
	}
	for _, point := range points {
		fp := &Parts{}
		if moreRandom {
			rand.Seed(time.Now().UnixNano())
			r := rand.Float64()
			fp.Initialize(img, point, view.Vector{X: vel.X * r, Y: vel.Y})
		} else {
			fp.Initialize(img, point, vel)
		}
		array = append(array, fp)
	}
	return array
}
