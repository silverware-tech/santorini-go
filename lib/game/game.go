package game

import (
	"fmt"
	"github.com/c2r0b/santorini.git/lib/EntityManager"
	"github.com/c2r0b/santorini.git/lib/players"
	"github.com/c2r0b/santorini.git/lib/utility"
	"github.com/rs/zerolog/log"
)

type Game struct {
	Players       []players.Player
	EntityManager EntityManager.EntityManager
}

func New(setup Setup) Game {
	// generate game EntityManager
	entityManager := EntityManager.New(utility.X_SIZE, utility.Y_SIZE, setup.getCharacters())
	g := Game{setup.Players, entityManager}
	return g
}

func (game *Game) Start() {
	turn := 0
	for {
		var player = game.Players[turn%len(game.Players)]

		log.Info().Msgf("%s turn. %d", player.GetName(), turn)

		selectedCharacter, moveDestination, buildPoint, err := player.DoTurn(game.EntityManager)
		if err != nil {
			log.Error().Msg(err.Error())
			fmt.Printf("Player %s loose", player.GetName())
			return
		}

		endGame, err := game.EntityManager.Move(selectedCharacter, moveDestination)

		if err != nil {
			log.Error().Msg(err.Error())
			return
		}
		if endGame {
			log.Info().Msgf("Player %v won", player.GetName())
			return
		}

		// ask user where to build a block to increase the height of a cell
		err = game.EntityManager.Build(selectedCharacter, buildPoint)
		if err != nil {
			log.Info().Msg(err.Error())
			return
		}

		log.Info().Msgf("Player %v built a block", player.GetName())

		turn++
	}
}
