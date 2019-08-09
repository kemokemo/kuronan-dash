package music

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := LoadAudioContext()
	if err != nil {
		fmt.Println("failed to LoadAudioContext:", err)
		return
	}

	err = LoadMusic()
	if err != nil {
		fmt.Println("failed to LoadMusic:", err)
		return
	}
	defer func() {
		e := CloseMusic()
		fmt.Println("failed to CloseMusic:", e)
	}()

	os.Exit(m.Run())
}

func TestDisc(t *testing.T) {
	tests := []struct {
		name string
		disc *Disc
	}{
		{"Title", Title},
		{"Stage01", Stage01},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.disc == nil {
				t.Errorf("Disc '%v' is nil, loading error", tt.name)
			}
		})
	}
}
