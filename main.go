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
	game.Start()
}
