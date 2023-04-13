package game

import (
	"fmt"
	"github.com/c2r0b/santorini.git/lib/EntityManager"
	"github.com/c2r0b/santorini.git/lib/player"
	"github.com/rs/zerolog/log"
)

const X_SIZE = 5
const Y_SIZE = 5
const MIN_PLAYERS = 2
const MAX_PLAYERS = 4
const MAX_CHARACTERS_PER_PLAYER = 2
const MAX_CHARACTERS = MAX_PLAYERS * MAX_CHARACTERS_PER_PLAYER

type Game struct {
	Players       []player.Player
	EntityManager EntityManager.EntityManager
}

func New(setup Setup) Game {
	// generate game EntityManager
	entityManager := EntityManager.New(X_SIZE, Y_SIZE, setup.Characters)
	g := Game{setup.Players, entityManager}
	return g
}

func (game *Game) Start() {
	for {
		for i := 0; i < MAX_PLAYERS; i++ {
			var oldX, oldY, newX, newY int

			game.EntityManager.PrintBoard()
			log.Info().Msgf("%s turn", game.Players[i].Name)

			log.Info().Msg("Move from (X):")
			fmt.Scan(&oldX)
			log.Info().Msg("Move from (Y):")
			fmt.Scan(&oldY)

			log.Info().Msg("Where to move (X):")
			fmt.Scan(&newX)
			log.Info().Msg("Where to move (Y):")
			fmt.Scan(&newY)

			log.Info().Msgf("Move %v to position (%v,%v)\n", game.Players[i], newX, newY)

			err := game.EntityManager.Move(newX, newY)

			if err != nil {
				log.Info().Msg(err.Error())
			} else {
				log.Info().Msgf("Player %v moved", game.Players[i].Name)
				i++
			}

		}
	}
}

func (game *Game) IsOver() bool {
	// TODO: implement
	return false
}
