package main

import (
	"os"

	"github.com/c2r0b/santorini.git/lib/game"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Logger = log.With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msg("Hello, Santorini!")
	var setup = game.AskSetup()

	gameManager := game.New(setup)

	gameManager.Start()
}
