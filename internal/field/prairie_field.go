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
	bg            *ebiten.Image
	prairie       *ebiten.Image
	mountainsNear []Mountain
	mountainsFar  []Mountain
	cloudsNear    []Cloud
	cloudsFar     []Cloud

	speed       ScrollSpeed
	viewPrairie view.RotateViewport
}

// Initialize initializes all resources to draw.
func (p *PrairieField) Initialize() {
	p.bg = images.SkyBackground
	p.prairie = images.TilePrairie
	p.createMountains()
	p.createClouds()

	p.viewPrairie = view.RotateViewport{}
	p.viewPrairie.SetSize(p.prairie.Size())
	p.viewPrairie.SetVelocity(2.0)
}

const mountNum = 3

func (p *PrairieField) createMountains() {
	rand.Seed(time.Now().UnixNano())

	wMt, hMt := images.MountainNear.Size()
	_, hP := images.TilePrairie.Size()
	for _, h := range LaneHeights {
		for index := 0; index < cloudNum; index++ {
			m := Mountain{}
			r := rand.Float32()
			m.Initialize(images.MountainNear,
				view.Position{
					X: wMt*index + int(500*r),
					Y: h - hMt + hP,
				},
				1.0,
			)
			p.mountainsNear = append(p.mountainsNear, m)
		}
	}

	wMt, hMt = images.MountainFar.Size()
	for _, h := range LaneHeights {
		for index := 0; index < cloudNum; index++ {
			m := Mountain{}
			r := rand.Float32()
			m.Initialize(images.MountainFar,
				view.Position{
					X: wMt*index + int(500*r),
					Y: h - hMt + hP,
				},
				0.5,
			)
			p.mountainsFar = append(p.mountainsFar, m)
		}
	}
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
	for i := range p.mountainsNear {
		p.mountainsNear[i].SetSpeed(speed)
	}
	for i := range p.mountainsFar {
		p.mountainsFar[i].SetSpeed(speed)
	}
	for i := range p.cloudsNear {
		p.cloudsNear[i].SetSpeed(speed)
	}
	for i := range p.cloudsFar {
		p.cloudsFar[i].SetSpeed(speed)
	}
}

// Update moves viewport for the all field parts.
func (p *PrairieField) Update() {
	switch p.speed {
	case Normal:
		p.viewPrairie.SetVelocity(2.0)
	case Slow:
		p.viewPrairie.SetVelocity(1.0)
	}

	p.viewPrairie.Move(view.Left)
	for i := range p.mountainsNear {
		p.mountainsNear[i].Update()
	}
	for i := range p.mountainsFar {
		p.mountainsFar[i].Update()
	}
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
	for i := range p.mountainsFar {
		err := p.mountainsFar[i].Draw(screen)
		if err != nil {
			return fmt.Errorf("failed to draw mountainsFar,%v", err)
		}
	}
	/// 遠くの雲
	for i := range p.cloudsFar {
		err := p.cloudsFar[i].Draw(screen)
		if err != nil {
			return fmt.Errorf("failed to draw cloudsFar,%v", err)
		}
	}

	// つぎに近くの風景を描画
	/// 近くの山。異なる速度のViewPort情報に切り替え
	for i := range p.mountainsNear {
		err := p.mountainsNear[i].Draw(screen)
		if err != nil {
			return fmt.Errorf("failed to draw mountainsNear,%v", err)
		}
	}
	/// 近くの雲
	for i := range p.cloudsNear {
		err := p.cloudsNear[i].Draw(screen)
		if err != nil {
			return fmt.Errorf("failed to draw cloudsNear,%v", err)
		}
	}

	// さいごのレーンを描画
	wP, _ := images.TilePrairie.Size()
	x16, y16 := p.viewPrairie.Position()
	offsetX, offsetY := float64(x16)/16, float64(y16)/16
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
