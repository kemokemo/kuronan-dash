// Copy from github.com/hajimehoshi/ebiten/example/blocks

package scenes

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	mplus "github.com/hajimehoshi/go-mplusbitmap"

	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/assets/music"

	chara "github.com/kemokemo/kuronan-dash/internal/character"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

// Stage01Scene is the scene for the 1st stage game.
type Stage01Scene struct {
	state    gameState
	player   *chara.Player
	disc     *music.Disc
	bg       *ebiten.Image
	prairie  *ebiten.Image
	mtNear   *ebiten.Image
	mtFar    *ebiten.Image
	viewPort view.Viewport
}

// Initialize initializes all resources.
func (s *Stage01Scene) Initialize() error {
	s.disc = music.Stage01
	s.player = chara.Selected
	err := s.player.SetLanes(laneHeights)
	if err != nil {
		return err
	}
	s.bg = images.SkyBackground
	s.prairie = images.TilePrairie
	s.mtNear = images.MountainNear
	s.mtFar = images.MountainFar

	s.viewPort = view.Viewport{}
	s.viewPort.SetSize(s.prairie.Size())
	return nil
}

// Update updates the status of this scene.
func (s *Stage01Scene) Update(state *GameState) error {
	switch s.state {
	case wait:
		if state.Input.StateForKey(ebiten.KeySpace) == 1 {
			s.state = run
			s.player.Start()
		}
	case run:
		if state.Input.StateForKey(ebiten.KeySpace) == 1 {
			s.state = pause
			s.player.Pause()
		} else if s.player.Position.X+50 > ScreenWidth-50 && s.state != gameover {
			// TODO: とりあえずゲームオーバーの練習
			s.state = gameover
			s.player.Stop()
		} else {
			s.player.Update()
			s.viewPort.Move(view.Left)
		}
	case pause:
		if state.Input.StateForKey(ebiten.KeySpace) == 1 {
			s.state = run
			s.player.ReStart()
		}
	case gameover:
		if state.Input.StateForKey(ebiten.KeySpace) == 1 {
			state.SceneManager.GoTo(&TitleScene{})
		}
	default:
		// unknown state
	}
	return nil
}

// Draw draws background and characters.
func (s *Stage01Scene) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.bg, &ebiten.DrawImageOptions{})

	s.drawFieldParts(screen)

	err := s.player.Draw(screen)
	if err != nil {
		log.Println("failed to draw a character:", err)
		return
	}

	text.Draw(screen, fmt.Sprintf("Now Playing: %s", s.disc.Name),
		mplus.Gothic12r, 12, 15, color.White)

	s.drawWithState(screen)
	// TODO: 衝突判定とSE再生
	err = s.checkCollision()
	if err != nil {
		log.Println("failed to check collisions:", err)
		return
	}
}

const repeat = 3

const (
	firstLaneHeight  = 200
	secondLaneHeight = firstLaneHeight + 170
	thirdLaneHeight  = secondLaneHeight + 170
)

var laneHeights = []int{firstLaneHeight, secondLaneHeight, thirdLaneHeight}

func (s *Stage01Scene) drawFieldParts(screen *ebiten.Image) {
	var k float64
	if s.player.GetState() == chara.Dash {
		k = 2
	} else {
		k = 1
	}

	x16, y16 := s.viewPort.Position()
	offsetX, offsetY := float64(x16)*k/16, float64(y16)*k/16

	wP, hP := s.prairie.Size()
	wMF, hMF := s.mtFar.Size()
	for _, h := range laneHeights {
		for i := 0; i < repeat; i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(wMF*i), float64(h-hMF+hP))
			op.GeoM.Translate(offsetX*0.5, offsetY*0.5)
			screen.DrawImage(s.mtFar, op)
		}
	}

	wMN, hMN := s.mtNear.Size()
	for _, h := range laneHeights {
		for i := 0; i < repeat; i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(wMN*i), float64(h-hMN+hP))
			op.GeoM.Translate(offsetX, offsetY)
			screen.DrawImage(s.mtNear, op)
		}
	}

	for _, h := range laneHeights {
		for i := 0; i < repeat; i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(wP*i), float64(h))
			op.GeoM.Translate(offsetX, offsetY)
			screen.DrawImage(s.prairie, op)
		}
	}
}

func (s *Stage01Scene) drawWithState(screen *ebiten.Image) {
	switch s.state {
	case wait:
		text.Draw(screen, "Spaceキーを押してゲームを開始してね！", mplus.Gothic12r, ScreenWidth/2-100, ScreenHeight/2, color.White)
	case pause:
		text.Draw(screen, "一時停止中だよ！", mplus.Gothic12r, ScreenWidth/2-100, ScreenHeight/2, color.White)
	case gameover:
		text.Draw(screen, "ゲームオーバー！Spaceを押してタイトルへ戻ってね！", mplus.Gothic12r, ScreenWidth/2-100, ScreenHeight/2, color.White)
	default:
		// nothing to draw
	}
}

func (s *Stage01Scene) checkCollision() error {
	// TODO: プレイヤーと障害物との衝突判定などをするよ
	return nil
}

// Close stops music
func (s *Stage01Scene) Close() error {
	err := s.disc.Stop()
	if err != nil {
		return fmt.Errorf("failed to stop music:%v", err)
	}
	return nil
}

// StartMusic starts playing music
func (s *Stage01Scene) StartMusic() error {
	return s.disc.Play()
}

// StopMusic stops playing music
func (s *Stage01Scene) StopMusic() error {
	return s.disc.Stop()
}
