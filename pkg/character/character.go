package character

import (
	"santorini/main/pkg/player"
)

type Character struct {
	player *player.Player
}

func New(player *player.Player) Character {
	return Character{
		player: player,
	}
}
