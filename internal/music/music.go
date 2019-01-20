package music

import (
	"log"

	"github.com/hajimehoshi/ebiten/audio"
)

const sampleRate = 44100

var audioContext *audio.Context

func init() {
	var err error
	audioContext, err = audio.NewContext(sampleRate)
	if err != nil {
		log.Println("Failed to create a new audio context: ", err)
	}
}
