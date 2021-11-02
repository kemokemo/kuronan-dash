package field

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// PrairieField is the field of prairie.
type PrairieField struct {
	bg           *ebiten.Image
	tile         *ebiten.Image
	fartherParts []ScrollableObject
	closerParts  []ScrollableObject
	obstacles    []Obstacle
	foods        []Food
	lanes        *Lanes
	goals        []Goal
}

// Initialize initializes all resources to draw.
func (p *PrairieField) Initialize(lanes *Lanes, goalX float64) {
	p.bg = images.SkyBackground
	p.tile = images.TilePrairie
	p.lanes = lanes
	p.createParts(goalX)
}

// create all field parts to draw.
func (p *PrairieField) createParts(goalX float64) {
	// Farther parts
	type ast struct {
		img *ebiten.Image
		gpf genPosFunc
		gps genPosSet
		kv  float64
	}

	assets := []ast{
		{images.MountainFar, genPosField, genPosSet{3, 1280, 500}, 0.2},
		{images.CloudFar, genPosAir, genPosSet{10, 2000, 2000}, 0.3},
		{images.MountainNear, genPosField, genPosSet{3, 518, 500}, 0.4},
		{images.CloudNear, genPosAir, genPosSet{10, 5000, 3000}, 0.7},
		{images.Grass1, genPosField, genPosSet{10, 600, 2000}, 0.8},
		{images.Grass3, genPosField, genPosSet{10, 900, 3000}, 0.85},
	}
	for _, asset := range assets {
		array := genParts(asset.img, p.lanes.GetLaneHeights(), asset.gpf, asset.gps, asset.kv)
		for i := range array {
			p.fartherParts = append(p.fartherParts, array[i])
		}
	}

	// Foods
	assets = []ast{
		{images.Onigiri, genPosField, genPosSet{30, 1000, 550}, 1.0},
	}
	for _, asset := range assets {
		array := genOnigiri(asset.img, p.lanes.GetLaneHeights(), asset.gpf, asset.gps, asset.kv)
		for i := range array {
			p.closerParts = append(p.closerParts, array[i])
			p.foods = append(p.foods, array[i])
		}
	}

	// Obstacles
	assets = []ast{
		{images.RockNormal, genPosField, genPosSet{30, 300, 1000}, 1.0},
	}
	for _, asset := range assets {
		array := genRocks(asset.img, p.lanes.GetLaneHeights(), asset.gpf, asset.gps, asset.kv)
		for i := range array {
			if randBool() {
				p.fartherParts = append(p.fartherParts, array[i])
			} else {
				p.closerParts = append(p.closerParts, array[i])
			}
			p.obstacles = append(p.obstacles, array[i])
		}
	}

	// Closer parts
	assets = []ast{
		{images.Grass3, genPosField, genPosSet{10, 900, 3000}, 0.9},
	}
	for _, asset := range assets {
		array := genParts(asset.img, p.lanes.GetLaneHeights(), asset.gpf, asset.gps, asset.kv)
		for i := range array {
			p.closerParts = append(p.closerParts, array[i])
		}
	}

	// Goal
	w, h := images.Goal_back.Size()
	for _, lh := range p.lanes.laneHeights {
		goal := Goal{}
		goal.Initialize(
			images.Goal_back,
			images.Goal_front,
			&view.Vector{X: goalX + view.DrawPosition - float64(w)/4, Y: lh - float64(h) + 2},
			&view.Vector{X: 0, Y: 0})
		p.goals = append(p.goals, goal)
	}

}

// Update moves viewport for the all field parts.
func (p *PrairieField) Update(scrollV *view.Vector) {
	p.lanes.Update(scrollV)
	for i := range p.fartherParts {
		p.fartherParts[i].Update(scrollV)
	}
	for i := range p.closerParts {
		p.closerParts[i].Update(scrollV)
	}
	for i := range p.goals {
		p.goals[i].Update(scrollV)
	}
}

// DrawFarther draws the farther field parts.
func (p *PrairieField) DrawFarther(screen *ebiten.Image) {
	// 背景を描画
	screen.DrawImage(p.bg, &ebiten.DrawImageOptions{})

	// レーンよりも遠くのパーツを描画
	for i := range p.fartherParts {
		p.fartherParts[i].Draw(screen)
	}

	// レーンを描画
	p.lanes.Draw(screen)

	for i := range p.goals {
		p.goals[i].DrawBack(screen)
	}
}

// DrawCloser draws the closer field part.
func (p *PrairieField) DrawCloser(screen *ebiten.Image) {
	// レーンよりも手前のパーツを描画

	/// 近くの草むら
	for i := range p.closerParts {
		p.closerParts[i].Draw(screen)
	}

	for i := range p.goals {
		p.goals[i].DrawFront(screen)
	}
}

// IsCollidedWithObstacles returns whether the r is collided with this item.
func (p *PrairieField) IsCollidedWithObstacles(hr *view.HitRectangle) bool {
	for i := range p.obstacles {
		if p.obstacles[i].IsCollided(hr) {
			return true
		}
	}

	return false
}

// EatFoods determines if there is a conflict between the player and the food.
// If it hits, it returns the stamina gained.
func (p *PrairieField) EatFoods(hr *view.HitRectangle) int {
	var stamina int
	for i := range p.foods {
		if p.foods[i].IsCollided(hr) {
			stamina += p.foods[i].Eat()
		}
	}

	return stamina
}
