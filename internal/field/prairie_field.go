package field

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// PrairieField is the field of prairie.
type PrairieField struct {
	bg      *ebiten.Image
	prairie *ebiten.Image
	mtNear  *ebiten.Image
	mtFar   *ebiten.Image
	cloud   *ebiten.Image
	cloud2  *ebiten.Image

	speed    ScrollSpeed
	viewFast view.Viewport
	viewSlow view.Viewport
}

// Initialize initializes all resources to draw.
func (p *PrairieField) Initialize() {
	p.bg = images.SkyBackground
	p.prairie = images.TilePrairie
	p.mtNear = images.MountainNear
	p.mtFar = images.MountainFar
	p.cloud = images.Cloud
	p.cloud2 = images.Cloud2

	p.viewFast = view.Viewport{}
	p.viewFast.SetSize(p.prairie.Size())
	p.viewFast.SetVelocity(2.0)

	p.viewSlow = view.Viewport{}
	p.viewSlow.SetSize(p.prairie.Size())
	p.viewSlow.SetVelocity(1.0)
}

// SetScrollSpeed sets the speed to scroll.
func (p *PrairieField) SetScrollSpeed(speed ScrollSpeed) {
	p.speed = speed
}

// Update moves viewport for the all field parts.
func (p *PrairieField) Update() {
	switch p.speed {
	case Normal:
		p.viewFast.SetVelocity(2.0)
		p.viewSlow.SetVelocity(1.0)
	case Slow:
		p.viewFast.SetVelocity(1.0)
		p.viewSlow.SetVelocity(0.5)
	}

	p.viewFast.Move(view.Left)
	p.viewSlow.Move(view.Left)
}

// Draw draws the all field parts.
func (p *PrairieField) Draw(screen *ebiten.Image) error {
	err := screen.DrawImage(p.bg, &ebiten.DrawImageOptions{})
	if err != nil {
		return fmt.Errorf("failed to draw a background,%v", err)
	}

	x16, y16 := p.viewSlow.Position()
	offsetX, offsetY := float64(x16)/16, float64(y16)/16

	// まず遠くの風景を描画
	wP, hP := p.prairie.Size()
	wC, hC := p.cloud.Size()
	wMF, hMF := p.mtFar.Size()
	for _, h := range LaneHeights {
		for i := 0; i < repeat; i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(wMF*i), float64(h-hMF+hP))
			op.GeoM.Translate(offsetX, offsetY)
			screen.DrawImage(p.mtFar, op)

			op.GeoM.Translate(float64(wC), float64(-hC))
			screen.DrawImage(p.cloud, op)
		}
	}

	// 異なる速度のViewPort情報に切り替え
	x16, y16 = p.viewFast.Position()
	offsetX, offsetY = float64(x16)/16, float64(y16)/16

	// つぎに近くの風景を描画
	wC, hC = p.cloud2.Size()
	wMN, hMN := p.mtNear.Size()
	for _, h := range LaneHeights {
		for i := 0; i < repeat; i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(wMN*i), float64(h-hMN+hP))
			op.GeoM.Translate(offsetX, offsetY)
			screen.DrawImage(p.mtNear, op)

			op.GeoM.Translate(float64(wC), float64(hC/2))
			screen.DrawImage(p.cloud2, op)
		}
	}

	// さいごのレーンを描画
	for _, h := range LaneHeights {
		for i := 0; i < repeat; i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(wP*i), float64(h))
			op.GeoM.Translate(offsetX, offsetY)
			screen.DrawImage(p.prairie, op)
		}
	}
	return nil
}
