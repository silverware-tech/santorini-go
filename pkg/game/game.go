package game

import (
	"fmt"
	"santorini/main/pkg/board"
	"santorini/main/pkg/character"
	"santorini/main/pkg/player"
	"strconv"

	"github.com/rs/zerolog/log"
)

const SIZE = 5
const MAX_PLAYERS = 4
const MAX_CHARACTERS_PER_PLAYER = 2
const MAX_CHARACTERS = MAX_PLAYERS * MAX_CHARACTERS_PER_PLAYER

type Game struct {
	Players    [MAX_PLAYERS]player.Player
	Characters [MAX_CHARACTERS]character.Character
	Board      board.Board
}

func New(numberOfPlayers int) Game {
	var players [MAX_PLAYERS]player.Player
	var characters [MAX_CHARACTERS]character.Character

	// generate game board
	board := board.New()

	// generate characters list (2 for each group)
	for i := 0; i < numberOfPlayers; i++ {
		players[i] = player.New("Player "+strconv.Itoa(i+1), false)

		for j := 0; j < MAX_CHARACTERS_PER_PLAYER; j++ {
			characters[i+j] = character.New(&players[i])
		}
	}
	fmt.Print(players)
	g := Game{players, characters, board}
	return g
}

func (g *Game) Start() {
	for i := 0; i < MAX_PLAYERS; i++ {
		if g.Players[i] {
			var characterToMove int
			var oldX, oldY, newX, newY int

			g.Board.Print()
			log.Info().Msgf("%s turn", g.Players[i].Name)

			log.Info().Msg("Move from (X):")
			fmt.Scan(&oldX)
			log.Info().Msg("Move from (Y):")
			fmt.Scan(&oldY)

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

func (g *Game) IsOver() bool {
	// TODO: implement
	return false
}
