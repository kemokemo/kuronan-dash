package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	kuronandash "github.com/kemokemo/kuronan-dash/internal"
	"github.com/kemokemo/kuronan-dash/internal/scenes"
)

const (
	exitOK = iota
	exitFailed
)

func main() {
	os.Exit(run())
}

func run() int {
	game, err := kuronandash.NewGame()
	if err != nil {
		log.Println("failed to create a new game:", err)
		return exitFailed
	}
	defer func() {
		e := game.Close()
		if e != nil {
			log.Println("failed to close this game:", e)
		}
	}()

	err = ebiten.Run(game.Update, scenes.ScreenWidth, scenes.ScreenHeight,
		1, "Kuronan Dash!")
	if err != nil {
		log.Println("failed to run:", err)
		return exitFailed
	}
	return exitOK
}
