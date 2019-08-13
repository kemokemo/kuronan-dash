package field

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// PrairieField is the field of prairie.
type PrairieField struct {
	bg         *ebiten.Image
	prairie    *ebiten.Image
	mtNear     *ebiten.Image
	mtFar      *ebiten.Image
	cloudsNear []Cloud
	cloudsFar  []Cloud

	speed       ScrollSpeed
	viewPrairie view.Viewport
	viewMtNear  view.Viewport
	viewMtFar   view.Viewport
}

// Initialize initializes all resources to draw.
func (p *PrairieField) Initialize() {
	p.bg = images.SkyBackground
	p.prairie = images.TilePrairie
	p.mtNear = images.MountainNear
	p.mtFar = images.MountainFar
	p.createClouds()

	p.viewPrairie = view.Viewport{}
	p.viewPrairie.SetSize(p.prairie.Size())
	p.viewPrairie.SetVelocity(2.0)

	p.viewMtNear = view.Viewport{}
	p.viewMtNear.SetSize(p.prairie.Size())
	p.viewMtNear.SetVelocity(1.0)

	p.viewMtFar = view.Viewport{}
	p.viewMtFar.SetSize(p.prairie.Size())
	p.viewMtFar.SetVelocity(0.5)
}

const cloudNum = 10

func (p *PrairieField) createClouds() {
	rand.Seed(time.Now().UnixNano())

	_, hC := images.CloudNear.Size()
	for _, h := range LaneHeights {
		for index := 0; index < cloudNum; index++ {
			c := Cloud{}
			r := rand.Float32()
			c.Initialize(images.CloudNear,
				view.Position{
					X: int(200*cloudNum + 2000*r),
					Y: h - 50 - int(100*r) - hC/2,
				})
			c.SetSpeed(Normal)

			r = rand.Float32()
			c.SetMagnification(r)
			p.cloudsNear = append(p.cloudsNear, c)
		}
	}

	_, hC = images.CloudFar.Size()
	for _, h := range LaneHeights {
		for index := 0; index < cloudNum; index++ {
			c := Cloud{}
			r := rand.Float32()
			c.Initialize(images.CloudFar,
				view.Position{
					X: int(500*cloudNum + 3000*r),
					Y: h - 40 - int(100*r) - hC/2,
				})
			c.SetSpeed(Normal)

			r = rand.Float32()
			c.SetMagnification(r)
			p.cloudsFar = append(p.cloudsFar, c)
		}
	}
}

// SetScrollSpeed sets the speed to scroll.
func (p *PrairieField) SetScrollSpeed(speed ScrollSpeed) {
	p.speed = speed
}

// Update moves viewport for the all field parts.
func (p *PrairieField) Update() {
	switch p.speed {
	case Normal:
		p.viewPrairie.SetVelocity(2.0)
		p.viewMtNear.SetVelocity(1.0)
		p.viewMtFar.SetVelocity(0.5)
		for i := range p.cloudsNear {
			p.cloudsNear[i].SetSpeed(Normal)
		}
		for i := range p.cloudsFar {
			p.cloudsFar[i].SetSpeed(Normal)
		}
	case Slow:
		p.viewPrairie.SetVelocity(1.0)
		p.viewMtNear.SetVelocity(0.5)
		p.viewMtFar.SetVelocity(0.25)
		for i := range p.cloudsNear {
			p.cloudsNear[i].SetSpeed(Slow)
		}
		for i := range p.cloudsFar {
			p.cloudsFar[i].SetSpeed(Slow)
		}
	}

	p.viewPrairie.Move(view.Left)
	p.viewMtNear.Move(view.Left)
	p.viewMtFar.Move(view.Left)
	for i := range p.cloudsNear {
		p.cloudsNear[i].Update()
	}
	for i := range p.cloudsFar {
		p.cloudsFar[i].Update()
	}
}

// Draw draws the all field parts.
func (p *PrairieField) Draw(screen *ebiten.Image) error {
	err := screen.DrawImage(p.bg, &ebiten.DrawImageOptions{})
	if err != nil {
		return fmt.Errorf("failed to draw a prairie background,%v", err)
	}

	// まず遠くの風景を描画
	/// 遠くの山
	x16, y16 := p.viewMtFar.Position()
	offsetX, offsetY := float64(x16)/16, float64(y16)/16
	wP, hP := p.prairie.Size()
	wMF, hMF := p.mtFar.Size()
	for _, h := range LaneHeights {
		for i := 0; i < repeat; i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(wMF*i), float64(h-hMF+hP))
			op.GeoM.Translate(offsetX, offsetY)
			err := screen.DrawImage(p.mtFar, op)
			if err != nil {
				return fmt.Errorf("failed to draw a mtFar,%v", err)
			}
		}
	}
	/// 遠くの雲
	for i := range p.cloudsFar {
		err := p.cloudsFar[i].Draw(screen)
		if err != nil {
			return fmt.Errorf("failed to draw a cloudFar,%v", err)
		}
	}

	// つぎに近くの風景を描画
	/// 近くの山。異なる速度のViewPort情報に切り替え
	x16, y16 = p.viewMtNear.Position()
	offsetX, offsetY = float64(x16)/16, float64(y16)/16
	wMN, hMN := p.mtNear.Size()
	for _, h := range LaneHeights {
		for i := 0; i < repeat; i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(wMN*i), float64(h-hMN+hP))
			op.GeoM.Translate(offsetX, offsetY)
			err := screen.DrawImage(p.mtNear, op)
			if err != nil {
				return fmt.Errorf("failed to draw a mtNear,%v", err)
			}
		}
	}
	/// 近くの雲
	for i := range p.cloudsNear {
		err := p.cloudsNear[i].Draw(screen)
		if err != nil {
			return fmt.Errorf("failed to draw a cloudNear,%v", err)
		}
	}

	// さいごのレーンを描画
	x16, y16 = p.viewPrairie.Position()
	offsetX, offsetY = float64(x16)/16, float64(y16)/16
	for _, h := range LaneHeights {
		for i := 0; i < repeat; i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(wP*i), float64(h))
			op.GeoM.Translate(offsetX, offsetY)
			err := screen.DrawImage(p.prairie, op)
			if err != nil {
				return fmt.Errorf("failed to draw the prairie field,%v", err)
			}
		}
	}
	return nil
}
