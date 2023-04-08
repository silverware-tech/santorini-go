package game

import (
	"fmt"
	"github.com/c2r0b/santorini.git/lib/EntityManager"
	"github.com/c2r0b/santorini.git/lib/character"
	"github.com/c2r0b/santorini.git/lib/player"
	"strconv"

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

func AskValue(message string, min, max int) int {
	var value int
	for {
		log.Info().Msg(message)
		scan, err := fmt.Scan(&value)
		if scan != 1 || err != nil {
			log.Error().Msg("Error during scan")
		}
		if min <= value && value <= max {
			log.Info().Msg("Ok")
			break
		}
		log.Info().Msg(fmt.Sprintf("The inserted value must be between %d and %d", min, max))
	}
	return value
}

func New() Game {
	var numberOfPlayers = AskValue("Number of players:", MIN_PLAYERS, MAX_PLAYERS)

	var players = make([]player.Player, numberOfPlayers)
	var characters = make([]character.Character, numberOfPlayers*MAX_CHARACTERS_PER_PLAYER)

	// generate characters list (2 for each group)
	for i := 0; i < numberOfPlayers; i++ {
		players[i] = player.New("Player "+strconv.Itoa(i+1), false)

		for j := 0; j < MAX_CHARACTERS_PER_PLAYER; j++ {
			characters[i+j] = character.New(&players[i])
		}
	}
	fmt.Print(players)
	// generate game EntityManager
	entityManager := EntityManager.New(X_SIZE, Y_SIZE, characters)
	g := Game{players, entityManager}
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
