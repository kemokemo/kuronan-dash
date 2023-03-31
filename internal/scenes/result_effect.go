package scenes

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/internal/anime"
)

const effectNum = 10

type ResultEffect struct {
	effects1 []*anime.TimeAnimation
	effects2 []*anime.TimeAnimation
	effects3 []*anime.TimeAnimation
	ops1     []*colorm.DrawImageOptions
	ops2     []*colorm.DrawImageOptions
	ops3     []*colorm.DrawImageOptions
	colors   []colorm.ColorM
	posRand  *rand.Rand
}

func (r *ResultEffect) Initialize() {
	colorRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < effectNum; i++ {
		var cm colorm.ColorM
		cm.Scale(0, 0, 0, 1)

		rc := colorRand.Float64()
		gc := colorRand.Float64()
		bc := colorRand.Float64()
		cm.Translate(rc, gc, bc, 0)

		r.colors = append(r.colors, cm)
	}

	r.posRand = rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < effectNum; i++ {
		ta := anime.NewTimeAnimation(images.ResultConfettiAnimation, 0.1)
		r.effects1 = append(r.effects1, ta)
		r.effects2 = append(r.effects2, ta)
		r.effects3 = append(r.effects3, ta)

		op := &colorm.DrawImageOptions{}
		op.GeoM.Translate(float64(r.posRand.Int31n(1000)-1000), float64(r.posRand.Int31n(1000)-1000))
		r.ops1 = append(r.ops1, op)

		op2 := &colorm.DrawImageOptions{}
		op2.GeoM.Translate(float64(r.posRand.Int31n(600)-800), float64(r.posRand.Int31n(1000)-1000))
		r.ops2 = append(r.ops2, op2)

		op3 := &colorm.DrawImageOptions{}
		op3.GeoM.Translate(float64(2000-r.posRand.Int31n(1000)), float64(r.posRand.Int31n(1000)-1000))
		r.ops3 = append(r.ops3, op3)
	}
}

func (r *ResultEffect) Update() {
	// todo 右下に落下
	for i := 0; i < effectNum; i++ {
		r.ops1[i].GeoM.Translate(r.posRand.Float64(), r.posRand.Float64())
		r.ops2[i].GeoM.Translate(0, r.posRand.Float64())
		r.ops3[i].GeoM.Translate(-r.posRand.Float64(), r.posRand.Float64())
	}
}

func (r *ResultEffect) Draw(screen *ebiten.Image) {
	for i := 0; i < effectNum; i++ {
		colorm.DrawImage(screen, r.effects1[i].GetCurrentFrame(), r.colors[i], r.ops1[i])
		colorm.DrawImage(screen, r.effects2[i].GetCurrentFrame(), r.colors[i], r.ops2[i])
		colorm.DrawImage(screen, r.effects3[i].GetCurrentFrame(), r.colors[i], r.ops3[i])
	}
}
