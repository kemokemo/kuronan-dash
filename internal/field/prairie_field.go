package field

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// PrairieField is the field of prairie.
type PrairieField struct {
	bg               *ebiten.Image
	bgOp             *ebiten.DrawImageOptions
	tile             *ebiten.Image
	fartherParts     []ScrollableObject
	closerParts      []ScrollableObject
	obstacles        []Obstacle
	foods            []Food
	lanes            *Lanes
	goals            []Goal
	collisionCounter int
	brokenCounter    int
}

// Initialize initializes all resources to draw.
func (p *PrairieField) Initialize(lanes *Lanes, goalX float64) {
	p.bg = images.SkyBackground
	p.bgOp = &ebiten.DrawImageOptions{}
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
		{images.MountainFar, genPosField, genPosSet{3, 1280, 500}, 0.1},
		{images.CloudFar, genPosAir, genPosSet{10, 2000, 500}, 0.15},
		{images.MountainNear, genPosField, genPosSet{3, 518, 500}, 0.2},
		{images.CloudNear, genPosAir, genPosSet{10, 3000, 400}, 0.5},
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
	oniArray := genOnigiri(images.OnigiriAnimation[0], p.lanes.GetLaneHeights(), genPosField, genPosSet{30, 1000, 550}, 1.0)
	for i := range oniArray {
		p.closerParts = append(p.closerParts, oniArray[i])
		p.foods = append(p.foods, oniArray[i])
	}
	manjuuArray := genYakiManjuu(images.IkariYakiAnimation[0], p.lanes.GetLaneHeights(), genPosField, genPosSet{10, 1500, 500}, 1.0)
	for i := range manjuuArray {
		p.closerParts = append(p.closerParts, manjuuArray[i])
		p.foods = append(p.foods, manjuuArray[i])
	}
	ikariArray := genIkariYaki(images.IkariYakiAnimation[0], p.lanes.GetLaneHeights(), genPosField, genPosSet{3, 2400, 600}, 1.0)
	for i := range ikariArray {
		p.closerParts = append(p.closerParts, ikariArray[i])
		p.foods = append(p.foods, ikariArray[i])
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
		{images.Grass3, genPosField, genPosSet{10, 900, 3000}, 1.0},
	}
	for _, asset := range assets {
		array := genParts(asset.img, p.lanes.GetLaneHeights(), asset.gpf, asset.gps, asset.kv)
		for i := range array {
			p.closerParts = append(p.closerParts, array[i])
		}
	}

	// Goal
	w := images.Goal_back.Bounds().Dx()
	h := images.Goal_back.Bounds().Dy()
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
	screen.DrawImage(p.bg, p.bgOp)

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
// If it hits, it returns the stamina and tension gained.
func (p *PrairieField) EatFoods(hr *view.HitRectangle) (int, int) {
	var stamina, tension int
	var s, t int
	for i := range p.foods {
		if p.foods[i].IsCollided(hr) {
			s, t = p.foods[i].Eat()
			stamina += s
			tension += t
		}
	}

	return stamina, tension
}

func (p *PrairieField) AttackObstacles(hr *view.HitRectangle, power float64) (int, int) {
	p.collisionCounter = 0
	p.brokenCounter = 0

	for i := range p.obstacles {
		if p.obstacles[i].IsCollided(hr) {
			p.collisionCounter++
			p.obstacles[i].Attack(power)
			if p.obstacles[i].IsBroken() {
				p.brokenCounter++
			}
		}
	}

	return p.collisionCounter, p.brokenCounter
}
