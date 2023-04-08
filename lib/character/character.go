package character

import (
	"github.com/c2r0b/santorini.git/lib/player"
)

type Character struct {
	player *player.Player
	X      int
	Y      int
}

func New(player *player.Player) Character {
	return Character{
		player: player,
	}
}
