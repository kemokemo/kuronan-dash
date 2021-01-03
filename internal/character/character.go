package character

import (
	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/assets/messages"
	"github.com/kemokemo/kuronan-dash/assets/se"
	"github.com/kemokemo/kuronan-dash/internal/anime"
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
		stamina:       NewStamina(130, 6),
	}

	Koma = &Player{
		StandingImage: images.KomaStanding,
		Description:   messages.DescKoma,
		animation:     anime.NewStepAnimation(images.KomaAnimation, 5),
		jumpSe:        se.Jump,
		dropSe:        se.Drop,
		stamina:       NewStamina(160, 11),
	}

	Shishimaru = &Player{
		StandingImage: images.ShishimaruStanding,
		Description:   messages.DescShishimaru,
		animation:     anime.NewStepAnimation(images.ShishimaruAnimation, 5),
		jumpSe:        se.Jump,
		dropSe:        se.Drop,
		stamina:       NewStamina(200, 17),
	}

	Selected = Kurona
	return nil
}
