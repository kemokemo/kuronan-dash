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
		StandingImage:  images.KuronaStanding,
		Description:    messages.DescKurona,
		attackImage:    images.AttackScratch,
		specialImage:   images.KuronaSpBack,
		animation:      anime.NewStepAnimation(images.KuronaAnimation, 5),
		jumpSe:         se.Jump,
		dropSe:         se.Drop,
		typeSe:         se.KuronaSe,
		stamina:        NewStamina(90, 6),
		tension:        NewTension(50, 4),
		power:          2.0,
		atkMaxDuration: 5,
		spMaxDuration:  41,
		vc:             move.NewKuronaVc(),
	}

	Koma = &Player{
		StandingImage:  images.KomaStanding,
		Description:    messages.DescKoma,
		attackImage:    images.AttackKomaFist,
		specialImage:   images.KomaSpBack,
		animation:      anime.NewStepAnimation(images.KomaAnimation, 5),
		jumpSe:         se.Jump,
		dropSe:         se.Drop,
		typeSe:         se.KomaSe,
		stamina:        NewStamina(100, 8),
		tension:        NewTension(70, 7),
		power:          5.0,
		atkMaxDuration: 9,
		spMaxDuration:  41,
		vc:             move.NewKomaVc(),
	}

	Shishimaru = &Player{
		StandingImage:  images.ShishimaruStanding,
		Description:    messages.DescShishimaru,
		attackImage:    images.AttackShishimaruFist,
		specialImage:   images.ShishimaruSpBack,
		animation:      anime.NewStepAnimation(images.ShishimaruAnimation, 5),
		jumpSe:         se.Jump,
		dropSe:         se.Drop,
		typeSe:         se.ShishimaruSe,
		stamina:        NewStamina(120, 12),
		tension:        NewTension(90, 6),
		power:          3.5,
		atkMaxDuration: 7,
		spMaxDuration:  41,
		vc:             move.NewShishimaruVc(),
	}

	Selected = Kurona
	return nil
}

func InitializeCharacter() {
	Selected = Kurona
}
