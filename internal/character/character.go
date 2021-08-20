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
		animation:     anime.NewStepAnimation(images.KuronaAnimation, 5),
		jumpSe:        se.Jump,
		dropSe:        se.Drop,
		stamina:       NewStamina(90, 6),
		vc:            move.NewKuronaVc(),
	}

	Koma = &Player{
		StandingImage: images.KomaStanding,
		Description:   messages.DescKoma,
		animation:     anime.NewStepAnimation(images.KomaAnimation, 5),
		jumpSe:        se.Jump,
		dropSe:        se.Drop,
		stamina:       NewStamina(100, 11),
		vc:            move.NewKomaVc(),
	}

	Shishimaru = &Player{
		StandingImage: images.ShishimaruStanding,
		Description:   messages.DescShishimaru,
		animation:     anime.NewStepAnimation(images.ShishimaruAnimation, 5),
		jumpSe:        se.Jump,
		dropSe:        se.Drop,
		stamina:       NewStamina(120, 17),
		vc:            move.NewShishimaruVc(),
	}

	Selected = Kurona
	return nil
}
