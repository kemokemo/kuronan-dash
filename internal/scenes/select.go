package scenes

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	mplus "github.com/hajimehoshi/go-mplusbitmap"
	"github.com/kemokemo/kuronan-dash/internal/music"
	"github.com/kemokemo/kuronan-dash/internal/objects"
	"github.com/kemokemo/kuronan-dash/internal/ui"
	"github.com/kemokemo/kuronan-dash/internal/util"
	"golang.org/x/image/font"
)

const (
	frameWidth    = 2
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
	jb         *music.JukeBox
	cm         *objects.CharacterManager
	infoMap    map[objects.CharacterType]*objects.CharacterInfo
	winMap     map[objects.CharacterType]*ui.FrameWindow
	selector   objects.CharacterType
	fontNormal font.Face
}

// NewSelectScene creates the new GameScene.
func NewSelectScene() *SelectScene {
	return &SelectScene{}
}

// SetResources sets the resources like music, character images and so on.
func (s *SelectScene) SetResources(j *music.JukeBox, cm *objects.CharacterManager) {
	s.jb = j
	err := s.jb.SelectDisc(music.Title)
	if err != nil {
		log.Printf("Failed to select disc:%v", err)
	}

	s.cm = cm
	s.infoMap = s.cm.GetCharacterInfoMap()
	windowWidth = (ScreenWidth - windowSpacing*2 - windowMargin*2) / len(s.infoMap)
	windowHeight = ScreenHeight - windowMargin*2 - 100

	s.winMap = make(map[objects.CharacterType]*ui.FrameWindow, len(s.infoMap))
	for cType := range s.infoMap {
		win, err := ui.NewFrameWindow(
			windowMargin+(windowWidth+windowSpacing)*int(cType),
			windowMargin*2, windowWidth, windowHeight, frameWidth)
		if err != nil {
			log.Println("failed to create a new frame window", err)
		}
		win.SetColors(
			color.RGBA{64, 64, 64, 255},
			color.RGBA{192, 192, 192, 255},
			color.RGBA{0, 148, 255, 255})
		if cType == objects.Kurona {
			s.selector = cType
			win.SetBlink(true)
		}
		s.winMap[cType] = win
	}
	s.fontNormal = mplus.Gothic12r
}

// Update updates the status of this scene.
func (s *SelectScene) Update(state *GameState) error {
	s.checkSelectorChanged()

	if state.Input.StateForKey(ebiten.KeySpace) == 1 {
		err := s.cm.SelectCharacter(s.selector)
		if err != nil {
			return err
		}
		state.SceneManager.GoTo(NewGameScene())
		return nil
	}
	if util.AnyGamepadAbstractButtonPressed(state.Input) {
		state.SceneManager.GoTo(NewGameScene())
		return nil
	}
	return nil
}

// Draw draws background and characters. This function play music too.
func (s *SelectScene) Draw(r *ebiten.Image) {
	err := s.jb.Play()
	if err != nil {
		log.Printf("Failed to play with JukeBox:%v", err)
		return
	}

	for cType := range s.winMap {
		if cType == s.selector {
			s.winMap[cType].SetBlink(true)
		} else {
			s.winMap[cType].SetBlink(false)
		}
		s.winMap[cType].DrawWindow(r)

		s.drawMainImage(r, cType)

		s.drawMessage(r, cType)
	}
	text.Draw(r, "← → のカーソルキーでキャラクターを選んでSpaceキーを押してね！",
		mplus.Gothic12r, windowMargin, windowMargin, color.White)
}

func (s *SelectScene) takeHorizontalCenterPosition(cType objects.CharacterType) (x, y float64) {
	rect := s.winMap[cType].GetWindowRect()
	width, _ := s.infoMap[cType].MainImage.Size()
	x = float64((rect.Max.X-rect.Min.X)/2 + rect.Min.X - (width*scale)/2)
	y = float64(rect.Min.Y + margin)
	return x, y
}

func (s *SelectScene) takeTextPosition(cType objects.CharacterType) image.Point {
	rect := s.winMap[cType].GetWindowRect()
	x := rect.Min.X + margin
	_, height := s.infoMap[cType].MainImage.Size()
	y := rect.Min.Y + margin*2 + height*scale
	return image.Point{X: x, Y: y}
}

func (s *SelectScene) checkSelectorChanged() {
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		if int(s.selector) < len(s.winMap)-1 {
			s.selector++
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		if int(s.selector) > 0 {
			s.selector--
		}
	}
}

func (s *SelectScene) drawMainImage(screen *ebiten.Image, cType objects.CharacterType) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale) // important: you have to scale before translating.
	op.GeoM.Translate(s.takeHorizontalCenterPosition(cType))
	err := screen.DrawImage(s.infoMap[cType].MainImage, op)
	if err != nil {
		log.Println(err)
	}
}

func (s *SelectScene) drawMessage(screen *ebiten.Image, cType objects.CharacterType) {
	runes := []rune(s.infoMap[cType].Description)
	rect := s.winMap[cType].GetWindowRect()
	splitlen := (rect.Max.X - rect.Min.X - margin) / fontSize
	startPoint := s.takeTextPosition(cType)

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
