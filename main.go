package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func main() {
	os.Exit(run())
}

func run() int {
	err := ebiten.Run(update, 400, 300, 2, "Kuronan Dash!")
	if err != nil {
		log.Println("Failed to run.", err)
		return 1
	}
	return 0
}

func update(screen *ebiten.Image) error {
	ebitenutil.DebugPrint(screen, "My first implementation of -Kuronan Dash!- with golang.")
	return nil
}
