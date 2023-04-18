package game

import (
	"github.com/c2r0b/santorini.git/lib/EntityManager"
	"github.com/c2r0b/santorini.git/lib/players"
	"github.com/rs/zerolog/log"
)

const X_SIZE = 5
const Y_SIZE = 5
const MIN_PLAYERS = 2
const MAX_PLAYERS = 4
const MAX_CHARACTERS_PER_PLAYER = 2
const MAX_CHARACTERS = MAX_PLAYERS * MAX_CHARACTERS_PER_PLAYER

type Game struct {
	Players       []players.Player
	EntityManager EntityManager.EntityManager
}

func New(setup Setup) Game {
	// generate game EntityManager
	entityManager := EntityManager.New(X_SIZE, Y_SIZE, setup.getCharacters())
	g := Game{setup.Players, entityManager}
	return g
}

func (game *Game) Start() {
	turn := 0
	for {
		var player = game.Players[turn%len(game.Players)]

		log.Info().Msgf("%s turn. %d", player.GetName(), turn)

		var character, moveDestination, buildPoint = player.DoTurn(game.EntityManager)

		endGame, err := game.EntityManager.Move(character, moveDestination)

		if err != nil {
			log.Info().Msg(err.Error())
			return
		}
		if endGame {
			log.Info().Msgf("Player %v won", player.GetName())
			return
		}

		// ask user where to build a block to increase the height of a cell
		err = game.EntityManager.Build(character, buildPoint)
		if err != nil {
			log.Info().Msg(err.Error())
			return
		}

		log.Info().Msgf("Player %v built a block", player.GetName())

		turn++
	}
}
