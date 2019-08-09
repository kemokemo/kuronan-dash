package character

import (
	"fmt"
	"os"
	"testing"

	"github.com/kemokemo/kuronan-dash/assets"
)

func TestMain(m *testing.M) {
	err := assets.LoadAssets()
	if err != nil {
		fmt.Println("failed to assets.LoadAssets:", err)
		return
	}
	defer func() {
		e := assets.CloseAssets()
		fmt.Println("failed to assets.CloseAssets:", e)
	}()

	err = NewPlayers()
	if err != nil {
		fmt.Println("failed to NewPlayers:", err)
		return
	}

	os.Exit(m.Run())
}

func TestNewPlayers(t *testing.T) {
	tests := []struct {
		name   string
		player *Player
	}{
		{"Kurona", Kurona},
		{"Koma", Koma},
		{"Shishimaru", Shishimaru},
		{"Selected", Selected},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.player == nil {
				t.Errorf("GamePlayer '%v' is nil, loading error", tt.name)
			}
		})
	}
}
