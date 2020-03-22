package field

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// PrairieField is the field of prairie.
type PrairieField struct {
	bg           *ebiten.Image
	tile         *ebiten.Image
	fartherParts []ScrollableObject
	closerParts  []ScrollableObject

	viewLane view.RotateViewport
}

// Initialize initializes all resources to draw.
func (p *PrairieField) Initialize() {
	p.bg = images.SkyBackground

	p.createFartherParts()
	p.createCloserParts()

	p.tile = images.TilePrairie
	p.viewLane = view.RotateViewport{}
	p.viewLane.SetSize(p.tile.Size())
	p.viewLane.SetVelocity(2.0)
}

func (p *PrairieField) createFartherParts() {
	assets := []struct {
		img        *ebiten.Image
		num        int
		param1     int
		param2     int
		vel        view.Vector
		moreRandom bool
	}{
		{images.MountainFar, 3, 1280, 500, view.Vector{X: -0.5, Y: 0.0}, false},
		{images.CloudFar, 10, 2000, 2000, view.Vector{X: -16.0, Y: 0.0}, true},
		{images.MountainNear, 3, 518, 500, view.Vector{X: -1.0, Y: 0.0}, false},
		{images.CloudNear, 10, 5000, 3000, view.Vector{X: -16.0, Y: 0.0}, true},
		{images.Grass1, 10, 600, 2000, view.Vector{X: -1.8, Y: 0.0}, false},
		{images.Grass3, 10, 900, 3000, view.Vector{X: -1.1, Y: 0.0}, false},
		{images.RockNormal, 30, 300, 1000, view.Vector{X: 0.0, Y: 0.0}, false},
	}

	for _, asset := range assets {
		array := create(asset.img, asset.num, asset.param1, asset.param2, asset.vel, asset.moreRandom)
		p.fartherParts = append(p.fartherParts, array...)
	}
}

func (p *PrairieField) createCloserParts() {
	assets := []struct {
		img        *ebiten.Image
		num        int
		param1     int
		param2     int
		vel        view.Vector
		moreRandom bool
	}{
		{images.Grass2, 10, 200, 1300, view.Vector{X: -1.4, Y: 0.0}, false},
	}

	for _, asset := range assets {
		array := create(asset.img, asset.num, asset.param1, asset.param2, asset.vel, asset.moreRandom)
		p.closerParts = append(p.closerParts, array...)
	}

}

// Update moves viewport for the all field parts.
func (p *PrairieField) Update(v view.Vector) {
	p.viewLane.SetVelocity(v.X)
	p.viewLane.Move(view.Left)

	for i := range p.fartherParts {
		p.fartherParts[i].Update(v)
	}
	for i := range p.closerParts {
		p.closerParts[i].Update(v)
	}
}

// DrawFarther draws the farther field parts.
func (p *PrairieField) DrawFarther(screen *ebiten.Image) error {
	// 背景を描画
	err := screen.DrawImage(p.bg, &ebiten.DrawImageOptions{})
	if err != nil {
		return fmt.Errorf("failed to draw a prairie background,%v", err)
	}

	// レーンよりも遠くかレーン上のパーツを描画
	for i := range p.fartherParts {
		err := p.fartherParts[i].Draw(screen)
		if err != nil {
			return fmt.Errorf("failed to draw fartherParts,%v", err)
		}
	}

	// レーンを描画
	wP, _ := images.TilePrairie.Size()
	x16, y16 := p.viewLane.Position()
	offsetX, offsetY := float64(x16)/16, float64(y16)/16
	for _, h := range LaneHeights {
		for i := 0; i < repeat; i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(wP*i), float64(h))
			op.GeoM.Translate(offsetX, offsetY)
			err := screen.DrawImage(p.tile, op)
			if err != nil {
				return fmt.Errorf("failed to draw the prairie field,%v", err)
			}
		}
	}

	return nil
}

// DrawCloser draws the closer field part.
func (p *PrairieField) DrawCloser(screen *ebiten.Image) error {
	// レーンよりも手前のパーツを描画

	/// 近くの草むら
	for i := range p.closerParts {
		err := p.closerParts[i].Draw(screen)
		if err != nil {
			return fmt.Errorf("failed to draw closerParts,%v", err)
		}
	}

	return nil
}
