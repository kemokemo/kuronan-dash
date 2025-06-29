package scenes

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	rating "github.com/kemokemo/ebiten-rating"
	vpad "github.com/kemokemo/ebiten-virtualpad"
	"github.com/kemokemo/kuronan-dash/assets/fonts"
	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/assets/messages"
	"github.com/kemokemo/kuronan-dash/assets/music"
	"github.com/kemokemo/kuronan-dash/assets/se"
	chara "github.com/kemokemo/kuronan-dash/internal/character"
	"github.com/kemokemo/kuronan-dash/internal/input"
	"github.com/kemokemo/kuronan-dash/internal/ui"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

const (
	frameWidth           = 5
	margin               = 30
	scale                = 2
	windowSpacing        = 20
	windowMargin         = 20
	fontSize             = 20
	lineSpacing          = 2
	ratingOffsetInitialX = 155
	ratingOffsetX        = 421
	ratingOffsetY        = 450
)

// SelectScene is the scene to select the player character.
type SelectScene struct {
	bg                *ebiten.Image
	bgViewPort        *view.Viewport
	disc              *music.Disc
	clickSe           *se.Player
	selectVoice       *se.Player
	msgWindow         *ui.MessageWindow
	fontNormal        *text.GoTextFace
	iChecker          input.InputChecker
	vChecker          input.VolumeChecker
	volumeBtn         vpad.SelectButton
	goButton          vpad.TriggerButton
	charaList         []*chara.Player
	ratingMatrix      [][]*rating.Rating
	winRectArray      []image.Rectangle
	selectArray       []vpad.SelectButton
	selectedIndex     int
	selectChanged     bool
	lenChara          int
	curtain           *Curtain
	isStarting        bool
	isClosing         bool
	fpsTextOp         *text.DrawOptions
	speedTextOpList   []*text.DrawOptions
	staminaTextOpList []*text.DrawOptions
	powerTextOpList   []*text.DrawOptions
}

// Initialize initializes all resources.
func (s *SelectScene) Initialize() error {
	s.bg = images.SelectBackground
	s.bgViewPort = &view.Viewport{}
	s.bgViewPort.SetSize(s.bg.Bounds().Dx(), s.bg.Bounds().Dy())
	s.bgViewPort.SetVelocity(1.0)
	s.bgViewPort.SetLoop(true)
	s.disc = music.Select
	s.clickSe = se.MenuSelect
	s.selectVoice = se.CharacterSelectVoice

	s.charaList = []*chara.Player{chara.Kurona, chara.Koma, chara.Shishimaru}
	s.lenChara = len(s.charaList)
	s.winRectArray = make([]image.Rectangle, s.lenChara)
	s.selectArray = make([]vpad.SelectButton, s.lenChara)
	s.selectedIndex = 0
	chara.InitializeCharacter() // インデックスを更新したら選択キャラクターも初期化しよう

	for i := 0; i < len(s.charaList); i++ {
		ratingList := []*rating.Rating{}
		sp, po, st := s.charaList[i].GetDisplayParameters()
		spRate := rating.NewRating(images.RatingStar, ratingOffsetInitialX+i*ratingOffsetX, ratingOffsetY, 10)
		spRate.SetValue(sp)
		ratingList = append(ratingList, spRate)
		poRate := rating.NewRating(images.RatingStar, ratingOffsetInitialX+i*ratingOffsetX, ratingOffsetY+40, 10)
		poRate.SetValue(po)
		ratingList = append(ratingList, poRate)
		stRate := rating.NewRating(images.RatingStar, ratingOffsetInitialX+i*ratingOffsetX, ratingOffsetY+80, 10)
		stRate.SetValue(st / 15.0)
		ratingList = append(ratingList, stRate)
		s.ratingMatrix = append(s.ratingMatrix, ratingList)
	}

	iw := images.CharaWindow.Bounds().Dx()
	ih := images.CharaWindow.Bounds().Dy()
	for i := range s.selectArray {
		window := vpad.NewSelectButton(images.CharaWindow, vpad.JustPressed, vpad.SelectColor)
		x := windowMargin + (iw+windowSpacing)*int(i)
		y := windowMargin*2 + 60
		window.SetLocation(x, y)
		s.selectArray[i] = window
		s.winRectArray[i] = image.Rectangle{
			Min: image.Point{x, y},
			Max: image.Point{x + iw, y + ih},
		}
	}
	s.selectArray[s.selectedIndex].SetSelectState(true)

	s.goButton = vpad.NewTriggerButton(images.CharaSelectButton, vpad.JustReleased, vpad.SelectColor)
	s.goButton.SetLocation(view.ScreenWidth-220, view.ScreenHeight-90)
	s.goButton.SetTriggerButton([]ebiten.Key{ebiten.KeySpace})
	s.volumeBtn = vpad.NewSelectButton(images.VolumeOffButton, vpad.JustPressed, vpad.SelectColor)
	s.volumeBtn.SetLocation(view.ScreenWidth-58, 10)
	s.volumeBtn.SetSelectImage(images.VolumeOnButton)
	s.volumeBtn.SetSelectKeys([]ebiten.Key{ebiten.KeyV})

	s.iChecker = &input.SelectInputChecker{GoBtn: s.goButton}
	s.vChecker = &input.VolumeInputChecker{VolumeBtn: s.volumeBtn}

	s.fontNormal = fonts.GamerFontM
	s.msgWindow = ui.NewMessageWindow(windowMargin+5, windowMargin+13, view.ScreenWidth-windowMargin*2-80, 42, frameWidth)
	s.msgWindow.SetColors(
		color.RGBA{64, 64, 64, 255},
		color.RGBA{192, 192, 192, 255},
		color.RGBA{33, 228, 68, 255})

	s.curtain = NewCurtain()
	s.isStarting = false
	s.isClosing = false

	s.fpsTextOp = &text.DrawOptions{}
	s.fpsTextOp.GeoM.Translate(10, view.ScreenHeight-15)
	s.fpsTextOp.ColorScale.ScaleWithColor(color.White)

	for i := range s.ratingMatrix {
		speedTextOp := &text.DrawOptions{}
		speedTextOp.GeoM.Translate(float64(ratingOffsetInitialX+ratingOffsetX*i-85), ratingOffsetY+23)
		speedTextOp.ColorScale.ScaleWithColor(color.White)
		s.speedTextOpList = append(s.speedTextOpList, speedTextOp)

		staminaTextOp := &text.DrawOptions{}
		staminaTextOp.GeoM.Translate(float64(ratingOffsetInitialX+ratingOffsetX*i-85), ratingOffsetY+40+23)
		staminaTextOp.ColorScale.ScaleWithColor(color.White)
		s.staminaTextOpList = append(s.staminaTextOpList, staminaTextOp)

		powerTextOp := &text.DrawOptions{}
		powerTextOp.GeoM.Translate(float64(ratingOffsetInitialX+ratingOffsetX*i-85), ratingOffsetY+80+23)
		powerTextOp.ColorScale.ScaleWithColor(color.White)
		s.powerTextOpList = append(s.powerTextOpList, powerTextOp)
	}

	return nil
}

// Update updates the status of this scene.
func (s *SelectScene) Update(state *GameState) {
	if s.isStarting || s.isClosing {
		s.curtain.Update()

		if s.curtain.IsFinished() {
			if s.isClosing {
				err := state.SceneManager.GoTo(&Stage01Scene{})
				if err != nil {
					log.Println("failed to go to the Stage01Scene: ", err)
				}
			} else if s.isStarting {
				s.isStarting = false
			}
		}
		return
	}

	s.updateVolume()

	if !s.selectVoice.IsPlaying() {
		s.disc.SetVolume(0.5)
	}

	s.selectChanged = false
	s.iChecker.Update()
	if s.iChecker.TriggeredLeft() {
		if s.selectedIndex > 0 {
			s.selectChanged = true
			s.selectedIndex--
		}
	}
	if s.iChecker.TriggeredRight() {
		if s.selectedIndex < s.lenChara-1 {
			s.selectChanged = true
			s.selectedIndex++
		}
	}

	if !s.selectChanged {
		for i := range s.selectArray {
			s.selectArray[i].Update()
			if s.selectArray[i].IsSelected() && s.selectedIndex != i {
				s.selectChanged = true
				s.selectedIndex = i
				break
			}
		}
	}

	if s.selectChanged {
		chara.Selected = s.charaList[s.selectedIndex]
		s.clickSe.Play()
	}

	for i := range s.selectArray {
		if i == s.selectedIndex {
			s.selectArray[i].SetSelectState(true)
		} else {
			s.selectArray[i].SetSelectState(false)
		}
	}

	s.goButton.Update()
	if s.iChecker.TriggeredStart() {
		s.isClosing = true
		s.curtain.Start(true)
		s.clickSe.Play()
		return
	}

	s.bgViewPort.Move(view.UpperRight)
}

// updateVolume updates the volume on/off state of music and sounds.
// If you add some sounds, please add this logic.
func (s *SelectScene) updateVolume() {
	s.vChecker.Update()

	if s.vChecker.JustVolumeOn() {
		s.setVolume(true)
		s.disc.Play()
	} else if s.vChecker.JustVolumeOff() {
		s.setVolume(false)
	}
}

func (s *SelectScene) setVolume(flag bool) {
	s.disc.SetVolumeFlag(flag)
	s.clickSe.SetVolumeFlag(flag)
	s.selectVoice.SetVolumeFlag(flag)
}

// Draw draws background and characters.
func (s *SelectScene) Draw(screen *ebiten.Image) {
	s.drawBackground(screen)
	s.drawWindows(screen)
	for i := range s.selectArray {
		s.selectArray[i].Draw(screen)
	}
	s.drawCharacters(screen)
	text.Draw(screen, fmt.Sprintf("FPS: %3.1f", ebiten.ActualFPS()), fonts.GamerFontSS, s.fpsTextOp)
	s.goButton.Draw(screen)
	s.volumeBtn.Draw(screen)

	if s.isStarting || s.isClosing {
		s.curtain.Draw(screen)
	}
}

func (s *SelectScene) drawBackground(screen *ebiten.Image) {
	x16, y16 := s.bgViewPort.Position()
	offsetX, offsetY := float64(x16)/16, float64(y16)/16

	// Draw bgImage on the screen repeatedly.
	const repeat = 3
	w := s.bg.Bounds().Dx()
	h := s.bg.Bounds().Dy()
	for j := 0; j < repeat; j++ {
		for i := 0; i < repeat; i++ {
			op := &ebiten.DrawImageOptions{}
			screenWidth := screen.Bounds().Dx()
			op.GeoM.Translate(float64(screenWidth)-float64(w*(i+1)), float64(h*j))
			op.GeoM.Translate(offsetX, offsetY)
			screen.DrawImage(s.bg, op)
		}
	}
}

func (s *SelectScene) drawWindows(screen *ebiten.Image) {
	s.msgWindow.DrawWindow(screen, messages.SelectStart)
}

func (s *SelectScene) drawCharacters(screen *ebiten.Image) {
	for i := range s.charaList {
		s.drawChara(screen, i)
		s.drawMessage(screen, i)
	}

	for i := range s.ratingMatrix {
		text.Draw(screen, "スピード:", fonts.GamerFontM, s.speedTextOpList[i])
		text.Draw(screen, "パワー:", fonts.GamerFontM, s.staminaTextOpList[i])
		text.Draw(screen, "スタミナ:", fonts.GamerFontM, s.powerTextOpList[i])

		for j := range s.ratingMatrix[i] {
			s.ratingMatrix[i][j].Draw(screen)
		}
	}
}

func (s *SelectScene) drawChara(screen *ebiten.Image, i int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale) // important: you have to scale before translating.
	op.GeoM.Translate(s.takeHorizontalCenterPosition(i))
	screen.DrawImage(s.charaList[i].StandingImage, op)
}

func (s *SelectScene) takeHorizontalCenterPosition(i int) (x, y float64) {
	rect := s.winRectArray[i]
	width := s.charaList[i].StandingImage.Bounds().Dx()
	x = float64((rect.Max.X-rect.Min.X)/2 + rect.Min.X - (width*scale)/2)
	y = float64(rect.Min.Y + margin)
	return x, y
}

func (s *SelectScene) drawMessage(screen *ebiten.Image, i int) {
	rect := s.winRectArray[i]
	splitlen := (rect.Max.X - rect.Min.X) / fontSize
	startPoint := s.takeTextPosition(i)
	lineNum := 1

	rows := strings.Split(s.charaList[i].Description, "\n")
	x := startPoint.X
	y := startPoint.Y
	for _, row := range rows {
		runes := []rune(row)

		for i := 0; i < len(runes); i += splitlen {
			y = y + fontSize + lineSpacing
			tOp := &text.DrawOptions{}
			tOp.GeoM.Translate(float64(x), float64(y))
			tOp.ColorScale.ScaleWithColor(color.White)
			if i+splitlen < len(runes) {
				text.Draw(screen, string(runes[i:(i+splitlen)]), s.fontNormal, tOp)
			} else {
				text.Draw(screen, string(runes[i:]), s.fontNormal, tOp)
			}
			lineNum++
		}
	}
}

func (s *SelectScene) takeTextPosition(i int) image.Point {
	rect := s.winRectArray[i]
	x := rect.Min.X + margin
	height := s.charaList[i].StandingImage.Bounds().Dy()
	y := rect.Min.Y + margin*2 + height*scale
	return image.Point{X: x, Y: y}
}

// StartMusic starts playing music
func (s *SelectScene) StartMusic(isVolumeOn bool) {
	s.volumeBtn.SetSelectState(isVolumeOn)
	if isVolumeOn {
		s.disc.SetVolume(0.3)
		s.disc.Play()
		s.selectVoice.Play()
	}
	s.isStarting = true
	s.curtain.Start(false)
}

// StopMusic stops playing music and sound effects
func (s *SelectScene) StopMusic() error {
	var err, e error
	e = s.selectVoice.Close()
	if e != nil {
		err = fmt.Errorf("%v, %v", err, e)
	}
	e = s.disc.Stop()
	if e != nil {
		err = fmt.Errorf("%v, %v", err, e)
	}

	return err
}

func (s *SelectScene) IsVolumeOn() bool {
	return s.vChecker.IsVolumeOn()
}
