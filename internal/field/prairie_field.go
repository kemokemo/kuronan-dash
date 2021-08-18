package field

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// gameSpeed is the scroll speed for the lane to move.
const gameSpeed = 2.0

// PrairieField is the field of prairie.
type PrairieField struct {
	bg           *ebiten.Image
	tile         *ebiten.Image
	fartherParts []ScrollableObject
	closerParts  []ScrollableObject
	obstacles    []Obstacle
	foods        []Food
	lanes        *Lanes
}

// Initialize initializes all resources to draw.
func (p *PrairieField) Initialize(lanes *Lanes) {
	p.bg = images.SkyBackground
	p.tile = images.TilePrairie
	p.lanes = lanes
	p.createParts()
}

// create all field parts to draw.
func (p *PrairieField) createParts() {
	// Farther parts
	type ast struct {
		img *ebiten.Image
		gpf genPosFunc
		gps genPosSet
		gvf genVelFunc
		gvs genVelSet
	}

	assets := []ast{
		{images.MountainFar, genPosField, genPosSet{3, 1280, 500}, genVel, genVelSet{-0.5, 0.0, false}},
		{images.CloudFar, genPosAir, genPosSet{10, 2000, 2000}, genVel, genVelSet{-16.0, 0.0, true}},
		{images.MountainNear, genPosField, genPosSet{3, 518, 500}, genVel, genVelSet{-1.0, 0.0, false}},
		{images.CloudNear, genPosAir, genPosSet{10, 5000, 3000}, genVel, genVelSet{-16.0, 0.0, true}},
		{images.Grass1, genPosField, genPosSet{10, 600, 2000}, genVel, genVelSet{-1.8, 0.0, false}},
		{images.Grass3, genPosField, genPosSet{10, 900, 3000}, genVel, genVelSet{-1.1, 0.0, false}},
	}
	for _, asset := range assets {
		array := genParts(asset.img, p.lanes.GetLaneHeights(), asset.gpf, asset.gps, asset.gvf, asset.gvs)
		for i := range array {
			p.fartherParts = append(p.fartherParts, array[i])
		}
	}

	// Foods
	assets = []ast{
		{images.Onigiri, genPosField, genPosSet{30, 1000, 550}, genVel, genVelSet{0.0, 0.0, false}},
	}
	for _, asset := range assets {
		array := genOnigiri(asset.img, p.lanes.GetLaneHeights(), asset.gpf, asset.gps, asset.gvf, asset.gvs)
		for i := range array {
			p.closerParts = append(p.closerParts, array[i])
			p.foods = append(p.foods, array[i])
		}
	}

	// Obstacles
	assets = []ast{
		{images.RockNormal, genPosField, genPosSet{30, 300, 1000}, genVel, genVelSet{0.0, 0.0, false}},
	}
	for _, asset := range assets {
		array := genRocks(asset.img, p.lanes.GetLaneHeights(), asset.gpf, asset.gps, asset.gvf, asset.gvs)
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
		{images.Grass3, genPosField, genPosSet{10, 900, 3000}, genVel, genVelSet{-1 * gameSpeed, 0.0, false}},
	}
	for _, asset := range assets {
		array := genParts(asset.img, p.lanes.GetLaneHeights(), asset.gpf, asset.gps, asset.gvf, asset.gvs)
		for i := range array {
			p.closerParts = append(p.closerParts, array[i])
		}
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
}

// DrawCloser draws the closer field part.
func (p *PrairieField) DrawCloser(screen *ebiten.Image) {
	// レーンよりも手前のパーツを描画

	/// 近くの草むら
	for i := range p.closerParts {
		p.closerParts[i].Draw(screen)
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
