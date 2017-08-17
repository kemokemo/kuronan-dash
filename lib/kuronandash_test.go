package kuronandash

import (
	"log"
	"os"
	"testing"

	"github.com/hajimehoshi/ebiten/audio"
)

var (
	testContext *audio.Context
)

func TestMain(t *testing.M) {
	setup()
	exitCode := t.Run()
	os.Exit(exitCode)
}

func setup() {
	var err error
	testContext, err = audio.NewContext(sampleRate)
	if err != nil {
		log.Println("Failed to create audio context", err)
	}
}
