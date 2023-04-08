package main

import (
	"github.com/c2r0b/santorini.git/lib/game"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Logger = log.With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msg("Hello, Santorini!")

	gameManager := game.New()
	gameManager.Start()
}
