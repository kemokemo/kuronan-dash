package field

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

type Goal struct {
	bkImg *ebiten.Image
	frImg *ebiten.Image
	bkOp  *ebiten.DrawImageOptions
	frOp  *ebiten.DrawImageOptions
	v0    *view.Vector
}

// Initialize initializes the object.
//  args:
//   bkimg: the image to be drawn at the back of the character.
//   frimg: the image to be drawn at the front of the character.
//   pos: the initial position
//   vel: the velocity to move this object
func (g *Goal) Initialize(bkImg, frImg *ebiten.Image, pos *view.Vector, vel *view.Vector) {
	g.bkImg = bkImg
	g.frImg = frImg
	g.v0 = vel

	g.bkOp = &ebiten.DrawImageOptions{}
	g.bkOp.GeoM.Translate(pos.X, pos.Y)
	w := frImg.Bounds().Dx()
	g.frOp = &ebiten.DrawImageOptions{}
	g.frOp.GeoM.Translate(pos.X+float64(w), pos.Y)
}

// Update updates the position and velocity of this object.
//  args:
//   scrollV: the velocity to scroll field parts.
func (g *Goal) Update(scrollV *view.Vector) {
	g.bkOp.GeoM.Translate(g.v0.X+scrollV.X, g.v0.Y+scrollV.Y)
	g.frOp.GeoM.Translate(g.v0.X+scrollV.X, g.v0.Y+scrollV.Y)
}

func (g *Goal) DrawBack(screen *ebiten.Image) {
	screen.DrawImage(g.bkImg, g.bkOp)
}

func (g *Goal) DrawFront(screen *ebiten.Image) {
	screen.DrawImage(g.frImg, g.frOp)
}
