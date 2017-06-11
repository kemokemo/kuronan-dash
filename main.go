package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	dash "github.com/kemokemo/go-kuronandash/lib"
)

func main() {
	os.Exit(run())
}

func run() int {
	game := dash.Game{}
	game.Init()
	update := game.Update

	err := ebiten.Run(update, 800, 600, 1, "Kuronan Dash!")
	if err != nil {
		log.Println("Failed to run.", err)
		return 1
	}
	return 0
}
