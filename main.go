package main

import (
	"fmt"
	"os"
	"santorini/main/pkg/game"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

    log.Logger = log.With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})

	var numberOfPlayers int

	log.Info().Msg("Hello, Santorini!")

	log.Info().Msg("Number of players:")
	fmt.Scan(&numberOfPlayers)

	game := game.New(numberOfPlayers)

	for !game.IsOver() {
		groups := game.GetGroups()
		for i := 0; i < numberOfPlayers; {
			var characterToMove int
			var newX, newY int

			group := groups[i]

			game.GetBoard().Print()
			log.Info().Msgf("Player %v turn", group[0].GetGroup())

			log.Info().Msg("What character to move (1,2):")
			fmt.Scan(&characterToMove)
			player := group[characterToMove-1]

			log.Info().Msg("Where to move (X):")
			fmt.Scan(&newX)
			log.Info().Msg("Where to move (Y):")
			fmt.Scan(&newY)

			log.Info().Msgf("Move %v to position (%v,%v)\n", player, newX, newY)

			err := player.Move(newX, newY)

			if err != nil {
				log.Info().Msg(err.Error())
			} else {
				log.Info().Msgf("Player %v moved", player.GetName())
				i++
			}
		}
	}
}
