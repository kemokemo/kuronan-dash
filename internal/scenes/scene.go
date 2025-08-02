package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kemokemo/kuronan-dash/assets"
)

// Scene is interface for the all scenes.
type Scene interface {
	// Initializeは、内部Stateや各種Assetsの初期化などを行う。
	// Volume調整はゲーム画面を跨いで用いるGame共通項目なので、SceneManagerが作って、各Sceneにわたす。
	// gameSoundControlChから再生の指示を受け取って、各SceneはGameの必要なサウンドの再生を行う。
	Initialize(gameSoundControlCh <-chan assets.GameSoundControl) error

	// Updateは、内部Stateや従属するAssets、Playerなどの状態更新を行う。
	Update(state *GameState)

	// Drawは、Updateによって更新された各種従属インスタンスの描画を行う。
	Draw(screen *ebiten.Image)

	// Startは、音楽の再生やgoroutineの開始などゲーム開始可能な状態にする。
	Start(gameSoundState bool)

	// Closeは、従属するgoroutineの停止などSceneの終了処理を行う。
	Close() error
}
