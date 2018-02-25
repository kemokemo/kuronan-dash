package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	dash "github.com/kemokemo/kuronan-dash/lib"
)

const (
	exitOK = iota
	exitFailed
)

func main() {
	os.Exit(run())
}

func run() int {
	game := dash.Game{}
	err := game.Init()
	if err != nil {
		log.Println("Failed to initialize", err)
		return exitFailed
	}
	defer func() {
		e := game.Close()
		if err != nil {
			log.Println("Failed to close", e)
		}
	}()

	err = ebiten.Run(game.Update, dash.ScreenWidth, dash.ScreenHeight, 1, "Kuronan Dash!")
	if err != nil {
		log.Println("Failed to run.", err)
		return exitFailed
	}
	return exitOK
}
