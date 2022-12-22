package se

import (
	"fmt"
	"os"
	"testing"

	"github.com/kemokemo/kuronan-dash/assets/music"
)

func TestMain(m *testing.M) {
	music.LoadAudioContext()

	err := LoadSE()
	if err != nil {
		fmt.Println("failed to LoadSE:", err)
		return
	}
	defer func() {
		e := CloseSE()
		fmt.Println("failed to CloseSE:", e)
	}()

	os.Exit(m.Run())
}

func TestPlayer(t *testing.T) {
	tests := []struct {
		name   string
		player *Player
	}{
		{"Jump", JumpSe},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.player == nil {
				t.Errorf("SE Player '%v' is nil, loading error", tt.name)
			}
		})
	}
}
