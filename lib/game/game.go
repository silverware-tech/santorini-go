package game

import (
	"fmt"

	"github.com/c2r0b/santorini.git/lib/EntityManager"
	"github.com/c2r0b/santorini.git/lib/EntityManager/cell"
	"github.com/c2r0b/santorini.git/lib/character"
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
	entityManager := EntityManager.New(X_SIZE, Y_SIZE, setup.getCharacters())
	g := Game{setup.Players, entityManager}
	return g
}

func (game *Game) Start() {
	i := 0
	for {
		i = i % len(game.Players)
		var newX, newY int
		var characterId string
		var player = game.Players[i]

		log.Info().Msgf("%s turn", player.Name)
		for {
			game.EntityManager.PrintBoard()
			for {
				log.Info().Msg("Choose a character:")
				fmt.Scan(&characterId)
				if player.HasCharacter(characterId) {
					break
				}
				log.Info().Msg("Invalid character")
			}

			var character, err = player.GetCharacter(characterId)

			log.Info().Msg("Where to move (X):")
			fmt.Scan(&newX)
			log.Info().Msg("Where to move (Y):")
			fmt.Scan(&newY)

			log.Info().Msgf("Move character %v to position (%v,%v)\n", character.CharacterId, newX, newY)

			err = game.EntityManager.Move(character, newX, newY)

			if err != nil {
				log.Info().Msg(err.Error())
			} else {
				log.Info().Msgf("Player %v moved", game.Players[i].Name)
				if game.IsOver(character) {
					log.Info().Msgf("Player %v won", game.Players[i].Name)
					return
				}

				// ask user where to build a block to increase the height of a cell
				for {
					game.EntityManager.PrintBoard()
					log.Info().Msg("Where to build (X):")
					fmt.Scan(&newX)
					log.Info().Msg("Where to build (Y):")
					fmt.Scan(&newY)

					log.Info().Msgf("Build a block at position (%v,%v)\n", newX, newY)

					err = game.EntityManager.Build(character, newX, newY)

					if err != nil {
						log.Info().Msg(err.Error())
					} else {
						log.Info().Msgf("Player %v built a block", game.Players[i].Name)
						break
					}
				}
				i++
				break
			}
		}
	}
}

func (game *Game) IsOver(character *character.Character) bool {
	return game.EntityManager.Board.GetCell(character.X, character.Y).Height == cell.L3
}
