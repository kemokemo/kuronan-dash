package scenes

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	mplus "github.com/hajimehoshi/go-mplusbitmap"
	"golang.org/x/image/font"

	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/assets/music"
	chara "github.com/kemokemo/kuronan-dash/internal/character"
	"github.com/kemokemo/kuronan-dash/internal/input"
	"github.com/kemokemo/kuronan-dash/internal/ui"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

const (
	frameWidth    = 5
	margin        = 20
	scale         = 2
	windowSpacing = 15
	windowMargin  = 20
	fontSize      = 12
	lineSpacing   = 2
)

var (
	windowWidth  int
	windowHeight int
)

// SelectScene is the scene to select the player character.
type SelectScene struct {
	bg         *ebiten.Image
	bgViewPort *view.Viewport
	disc       *music.Disc
	charaList  []*chara.Player
	windowList []*ui.FrameWindow
	selector   int
	fontNormal font.Face
}

// Initialize initializes all resources.
func (s *SelectScene) Initialize() error {
	s.bg = images.SelectBackground
	s.bgViewPort = &view.Viewport{}
	s.bgViewPort.SetSize(s.bg.Size())
	s.disc = music.Title
	s.charaList = []*chara.Player{chara.Kurona, chara.Koma, chara.Shishimaru}

	windowWidth = (ScreenWidth - windowSpacing*2 - windowMargin*2) / len(s.charaList)
	windowHeight = ScreenHeight - windowMargin*2 - 100

	s.windowList = make([]*ui.FrameWindow, len(s.charaList))
	for i := range s.charaList {
		win, err := ui.NewFrameWindow(
			windowMargin+(windowWidth+windowSpacing)*int(i),
			windowMargin*2, windowWidth, windowHeight, frameWidth)
		if err != nil {
			log.Println("failed to create a new frame window:", err)
		}
		win.SetColors(
			color.RGBA{64, 64, 64, 255},
			color.RGBA{192, 192, 192, 255},
			color.RGBA{33, 228, 68, 255})
		if i == 0 {
			s.selector = i
			win.SetBlink(true)
		}
		s.windowList[i] = win
	}
	s.fontNormal = mplus.Gothic12r

	return nil
}

// Update updates the status of this scene.
func (s *SelectScene) Update(state *GameState) error {
	s.bgViewPort.Move(view.UpperRight)

	if ebiten.IsRunningSlowly() {
		return nil
	}

	s.checkSelectorChanged()
	if state.Input.StateForKey(ebiten.KeySpace) == 1 ||
		input.AnyGamepadAbstractButtonPressed(state.Input) {
		chara.Selected = s.charaList[s.selector]
		state.SceneManager.GoTo(&Stage01Scene{})
	}

	return nil
}

func (s *SelectScene) checkSelectorChanged() {
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		if int(s.selector) < len(s.windowList)-1 {
			s.selector++
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		if int(s.selector) > 0 {
			s.selector--
		}
	}
}

// Draw draws background and characters.
func (s *SelectScene) Draw(screen *ebiten.Image) {
	s.drawBackground(screen)
	text.Draw(screen, "← → のカーソルキーでキャラクターを選んでSpaceキーを押してね！",
		mplus.Gothic12r, windowMargin, windowMargin, color.Black)
	s.drawWindows(screen)
	s.drawCharacters(screen)
}

func (s *SelectScene) drawBackground(screen *ebiten.Image) {
	x16, y16 := s.bgViewPort.Position()
	offsetX, offsetY := float64(x16)/16, float64(y16)/16

	// Draw bgImage on the screen repeatedly.
	const repeat = 3
	w, h := s.bg.Size()
	for j := 0; j < repeat; j++ {
		for i := 0; i < repeat; i++ {
			op := &ebiten.DrawImageOptions{}
			screenWidth, _ := screen.Size()
			op.GeoM.Translate(float64(screenWidth)-float64(w*(i+1)), float64(h*j))
			op.GeoM.Translate(offsetX, offsetY)
			screen.DrawImage(s.bg, op)
		}
	}
}

func (s *SelectScene) drawWindows(screen *ebiten.Image) {
	for i := range s.windowList {
		if i == s.selector {
			s.windowList[i].SetBlink(true)
		} else {
			s.windowList[i].SetBlink(false)
		}
		s.windowList[i].DrawWindow(screen)
	}
}

func (s *SelectScene) drawCharacters(screen *ebiten.Image) {
	for i := range s.charaList {
		err := s.drawChara(screen, i)
		if err != nil {
			log.Println("failed to draw a character:", err)
		}
		s.drawMessage(screen, i)
	}
}

func (s *SelectScene) drawChara(screen *ebiten.Image, i int) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale) // important: you have to scale before translating.
	op.GeoM.Translate(s.takeHorizontalCenterPosition(i))
	return screen.DrawImage(s.charaList[i].StandingImage, op)
}

func (s *SelectScene) takeHorizontalCenterPosition(i int) (x, y float64) {
	rect := s.windowList[i].GetWindowRect()
	width, _ := s.charaList[i].StandingImage.Size()
	x = float64((rect.Max.X-rect.Min.X)/2 + rect.Min.X - (width*scale)/2)
	y = float64(rect.Min.Y + margin)
	return x, y
}

func (s *SelectScene) drawMessage(screen *ebiten.Image, i int) {
	runes := []rune(s.charaList[i].Description)
	rect := s.windowList[i].GetWindowRect()
	splitlen := (rect.Max.X - rect.Min.X - margin) / fontSize
	startPoint := s.takeTextPosition(i)

	lineNum := 1
	for i := 0; i < len(runes); i += splitlen {
		x := startPoint.X
		y := startPoint.Y + (fontSize+lineSpacing)*lineNum
		if i+splitlen < len(runes) {
			text.Draw(screen, string(runes[i:(i+splitlen)]), s.fontNormal, x, y, color.White)
		} else {
			text.Draw(screen, string(runes[i:]), s.fontNormal, x, y, color.White)
		}
		lineNum++
	}
}

func (s *SelectScene) takeTextPosition(i int) image.Point {
	rect := s.windowList[i].GetWindowRect()
	x := rect.Min.X + margin
	_, height := s.charaList[i].StandingImage.Size()
	y := rect.Min.Y + margin*2 + height*scale
	return image.Point{X: x, Y: y}
}

// StartMusic starts playing music
func (s *SelectScene) StartMusic() error {
	return s.disc.Play()
}

// StopMusic stops playing music
func (s *SelectScene) StopMusic() error {
	return s.disc.Stop()
}
