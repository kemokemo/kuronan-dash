// Copy from github.com/hajimehoshi/ebiten/example/blocks

package kuronandash

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var imageBackground *ebiten.Image

func init() {
	var err error
	imageBackground, _, err = ebitenutil.NewImageFromFile("assets/images/title/background.png", ebiten.FilterNearest)
	if err != nil {
		log.Printf("Failed to load the background image:%v", err)
		return
	}
}

type TitleScene struct {
	jukeBox *JukeBox
}

func anyGamepadAbstractButtonPressed(i *Input) bool {
	for _, b := range virtualGamepadButtons {
		if i.gamepadConfig.IsButtonPressed(b) {
			return true
		}
	}
	return false
}

func (s *TitleScene) SetResources(j *JukeBox, c *Character) {
	s.jukeBox = j
	err := s.jukeBox.SelectDisc("shibugaki_no_kuroneko")
	if err != nil {
		log.Printf("Failed to select disc:%v", err)
	}
}

func (s *TitleScene) Update(state *GameState) error {
	if state.Input.StateForKey(ebiten.KeySpace) == 1 {
		state.SceneManager.GoTo(NewGameScene())
		return nil
	}
	if anyGamepadAbstractButtonPressed(state.Input) {
		state.SceneManager.GoTo(NewGameScene())
		return nil
	}
	return nil
}

func (s *TitleScene) Draw(r *ebiten.Image) {
	err := s.jukeBox.Play()
	if err != nil {
		log.Printf("Failed to play with JukeBox:%v", err)
		return
	}

	s.drawTitleBackground(r)
	drawLogo(r, "Kuronan Dash!")
	drawMessage(r, "Press SPACE to start")
}

func (s *TitleScene) drawTitleBackground(r *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	r.DrawImage(imageBackground, op)
}

func drawLogo(r *ebiten.Image, str string) {
	const scale = 4
	x := 0
	y := 32
	drawTextWithShadowCenter(r, str, x, y, scale, color.NRGBA{0x00, 0x00, 0x80, 0xff}, ScreenWidth)
}

func drawMessage(r *ebiten.Image, message string) {
	x := 0
	y := ScreenHeight - 48
	drawTextWithShadowCenter(r, message, x, y, 1, color.NRGBA{0x80, 0, 0, 0xff}, ScreenWidth)
}
