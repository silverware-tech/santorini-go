package ai

import (
	"github.com/c2r0b/santorini.git/lib/character"
	"github.com/c2r0b/santorini.git/lib/game"
	"github.com/c2r0b/santorini.git/lib/players"
	"github.com/c2r0b/santorini.git/lib/utility"
	"testing"
)

func TestRandomAi(t *testing.T) {
	ai1Characters := []character.Character{
		{
			Position:    utility.Point{0, 0},
			CharacterId: "A",
		},
		{
			Position:    utility.Point{1, 1},
			CharacterId: "B",
		},
	}
	ai2Characters := []character.Character{
		{
			Position:    utility.Point{3, 3},
			CharacterId: "C",
		},
		{
			Position:    utility.Point{4, 4},
			CharacterId: "D",
		},
	}

	setup := game.Setup{
		Players: []players.Player{
			(&players.RandomAi{}).NewWithCharacters("AI 1", ai1Characters),
			(&players.RandomAi{}).NewWithCharacters("AI 2", ai2Characters),
		},
	}

	var gameManager game.Game
	for i := 0; i < 1000; i++ {
		gameManager = game.New(setup)
		gameManager.Start()
	}
}
