package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/kemokemo/kuronan-dash/internal"
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
		log.Println("Failed to create a new game", err)
		return exitFailed
	}
	defer func() {
		e := game.Close()
		if e != nil {
			log.Println("Failed to close", e)
		}
	}()

	err = ebiten.Run(game.Update, scenes.ScreenWidth, scenes.ScreenHeight,
		1, "Kuronan Dash!")
	if err != nil {
		log.Println("Failed to run.", err)
		return exitFailed
	}
	return exitOK
}
