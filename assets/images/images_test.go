package images

import (
	"fmt"
	"os"
	"testing"

	"github.com/hajimehoshi/ebiten"
)

func TestMain(m *testing.M) {
	err := LoadImages()
	if err != nil {
		fmt.Println("failed to LoadImages:", err)
		return
	}

	os.Exit(m.Run())
}

func TestImages(t *testing.T) {
	tests := []struct {
		name string
		img  *ebiten.Image
	}{
		// Background
		{"TitleBackground", TitleBackground},
		{"SelectBackground", SelectBackground},
		// Field parts
		{"TilePrairie", TilePrairie},
		{"TilePrairie", TilePrairie},
		// Character standing images
		{"KuronaStanding", KuronaStanding},
		{"KomaStanding", KomaStanding},
		{"ShishimaruStanding", ShishimaruStanding},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.img == nil {
				t.Errorf("Image '%v' is nil, loading error", tt.name)
			}
		})
	}
}

func TestAnimation(t *testing.T) {
	tests := []struct {
		name string
		anim []*ebiten.Image
	}{
		{"KuronaAnimation", KuronaAnimation},
		{"KomaAnimation", KomaAnimation},
		{"ShishimaruAnimation", ShishimaruAnimation},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.anim == nil {
				t.Errorf("Animation '%v' is nil, loading error", tt.name)
			}
		})
	}
}
