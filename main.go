package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	kuronandash "github.com/kemokemo/kuronan-dash/internal"
	"github.com/kemokemo/kuronan-dash/internal/view"
)

const (
	exitOK = iota
	exitFailed
)

func main() {
	os.Exit(run())
}

func run() int {
	game, err := kuronandash.NewGame(fmt.Sprintf("Version: %s.%s", Version, Revision))
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

	err = ebiten.Run(game.Update, view.ScreenWidth, view.ScreenHeight,
		1, "Kuronan Dash!")
	if err != nil {
		log.Println("failed to run:", err)
		return exitFailed
	}
	return exitOK
}
