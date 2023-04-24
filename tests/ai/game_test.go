package ai

import (
	"github.com/c2r0b/santorini.git/lib/character"
	"github.com/c2r0b/santorini.git/lib/game"
	"github.com/c2r0b/santorini.git/lib/players"
	"github.com/c2r0b/santorini.git/lib/utility"
	"math/rand"
	"testing"
	"time"
)

func TestRandomAi(t *testing.T) {

	setup := game.Setup{
		Players: []players.Player{
			&players.RandomAi{
				Name: "AI 1",
				Characters: []character.Character{
					{
						Position:    utility.Point{0, 0},
						CharacterId: "A",
					},
					{
						Position:    utility.Point{1, 1},
						CharacterId: "B",
					},
				},
				Rand: rand.New(rand.NewSource(time.Now().Unix())),
			},
			&players.RandomAi{
				Name: "AI 2",
				Characters: []character.Character{
					{
						Position:    utility.Point{3, 3},
						CharacterId: "C",
					},
					{
						Position:    utility.Point{4, 4},
						CharacterId: "D",
					},
				},
				Rand: rand.New(rand.NewSource(time.Now().Unix())),
			},
		},
	}

	gameManager := game.New(setup)

	gameManager.Start()

}
