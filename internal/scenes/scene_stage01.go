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
	state     gameState
	player    *chara.Player
	disc      *music.Disc
	prairie   *ebiten.Image
	fieldView view.Viewport
}

// Initialize initializes all resources.
func (s *Stage01Scene) Initialize() error {
	s.disc = music.Stage01
	s.player = chara.Selected
	s.player.SetPosition(chara.Position{X: 10, Y: 50})
	s.prairie = images.TilePrairie
	s.fieldView = view.Viewport{}
	s.fieldView.SetSize(s.prairie.Size())
	return nil
}

// Update updates the status of this scene.
func (s *Stage01Scene) Update(state *GameState) error {
	s.updateGameStatus(state)
	s.updatePlayerStatus(state)
	return nil
}

func (s *Stage01Scene) updateGameStatus(state *GameState) {
	switch s.state {
	case wait:
		if state.Input.StateForKey(ebiten.KeySpace) == 1 {
			s.state = run
		}
	case run:
		if state.Input.StateForKey(ebiten.KeySpace) == 1 {
			s.state = pause
		}
		// TODO: とりあえずゲームオーバーの練習
		if s.player.Position.X+50 > ScreenWidth-50 && s.state != gameover {
			s.state = gameover
		}
	case pause:
		if state.Input.StateForKey(ebiten.KeySpace) == 1 {
			s.state = run
		}
	case gameover:
		if state.Input.StateForKey(ebiten.KeySpace) == 1 {
			state.SceneManager.GoTo(&TitleScene{})
		}
	default:
		// unknown state
	}
}

func (s *Stage01Scene) updatePlayerStatus(state *GameState) {
	s.player.Update()
	s.fieldView.Move(view.Left)
}

// Draw draws background and characters.
func (s *Stage01Scene) Draw(screen *ebiten.Image) {
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
	firstLaneHeight  = 100
	secondLaneHeight = 300
	thirdLaneHeight  = 500
)

var laneHeights = []int{firstLaneHeight, secondLaneHeight, thirdLaneHeight}

func (s *Stage01Scene) drawFieldParts(screen *ebiten.Image) {
	// todo: 遅い方のViewPortも用意して、キャラクターの状態がwalkの場合はそっちを使う
	x16, y16 := s.fieldView.Position()
	offsetX, offsetY := float64(x16)/16, float64(y16)/16

	w, _ := s.prairie.Size()
	for _, h := range laneHeights {
		for i := 0; i < repeat; i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(w*i), float64(h))
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
