package character

import (
	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/assets/messages"
	"github.com/kemokemo/kuronan-dash/assets/se"
	"github.com/kemokemo/kuronan-dash/internal/anime"
	"github.com/kemokemo/kuronan-dash/internal/move"
)

// player characters
var (
	Kurona     *Player
	Koma       *Player
	Shishimaru *Player

	// Selected is the selected player.
	Selected *Player
)

// NewPlayers load all player characters.
func NewPlayers() error {
	Kurona = &Player{
		StandingImage: images.KuronaStanding,
		Description:   messages.DescKurona,
		attackImage:   images.AttackScratch,
		animation:     anime.NewStepAnimation(images.KuronaAnimation, 5),
		jumpSe:        se.Jump,
		dropSe:        se.Drop,
		typeSe:        se.KuronaSe,
		stamina:       NewStamina(90, 6),
		tension:       NewTension(50, 4),
		power:         2.0,
		maxDuration:   5,
		vc:            move.NewKuronaVc(),
	}

	Koma = &Player{
		StandingImage: images.KomaStanding,
		Description:   messages.DescKoma,
		attackImage:   images.AttackKomaFist,
		animation:     anime.NewStepAnimation(images.KomaAnimation, 5),
		jumpSe:        se.Jump,
		dropSe:        se.Drop,
		typeSe:        se.KomaSe,
		stamina:       NewStamina(100, 8),
		tension:       NewTension(70, 7),
		power:         5.0,
		maxDuration:   9,
		vc:            move.NewKomaVc(),
	}

	Shishimaru = &Player{
		StandingImage: images.ShishimaruStanding,
		Description:   messages.DescShishimaru,
		attackImage:   images.AttackKomaFist,
		animation:     anime.NewStepAnimation(images.ShishimaruAnimation, 5),
		jumpSe:        se.Jump,
		dropSe:        se.Drop,
		typeSe:        se.ShishimaruSe,
		stamina:       NewStamina(120, 12),
		tension:       NewTension(90, 6),
		power:         3.5,
		maxDuration:   7,
		vc:            move.NewShishimaruVc(),
	}

	Selected = Kurona
	return nil
}

func InitializeCharacter() {
	Selected = Kurona
}
