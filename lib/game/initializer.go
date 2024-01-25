package game

import (
	"fmt"
	"github.com/c2r0b/santorini.git/lib/character"
	"github.com/c2r0b/santorini.git/lib/players"
	"github.com/c2r0b/santorini.git/lib/utility"
	"github.com/rs/zerolog/log"
)

/**
Used for the initial setup of the game. This module asks all the questions about game setup, like:
- How many players
- Players names
- If a player is AI or Human
- Ask the position of every player, generate the position of AI (random)
*/

type Setup struct {
	Players []players.Player
}

func AskInt(message string, min, max int) int {
	var value int
	for {
		fmt.Print(message)
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

func AskAi(name string) bool {
	var value int
	for {
		fmt.Print("Is player " + name + " ai? (0/1): ")
		scan, err := fmt.Scan(&value)
		if scan != 1 || err != nil {
			log.Error().Msg("Error during scan")
		} else {
			break
		}
	}
	return value != 0
}

func AskName() string {
	fmt.Scanln()
	fmt.Print("Player Name: ")
	var value string
	fmt.Scanln(&value)
	return value
}

func AskSetup() Setup {
	var numberOfPlayers = AskInt("Number of players: ", utility.MIN_PLAYERS, utility.MAX_PLAYERS)

	var playerList = make([]players.Player, numberOfPlayers)

	// hashset for the positions
	var positions = make(map[utility.Point]bool)

	// generate characters list (2 for each group)
	for i := 0; i < numberOfPlayers; i++ {
		var playerName = AskName()
		var isIa = AskAi(playerName)

		var player players.Player
		if isIa {
			player = (&players.RandomAi{}).New(playerName)
		} else {
			player = (&players.Human{}).New(playerName)
		}
		for j := 0; j < utility.MAX_CHARACTERS_PER_PLAYER; j++ {
			id := string(rune('A' + (i*utility.MAX_CHARACTERS_PER_PLAYER + j)))

			// ask for unique position X and Y for the character on the board
			var x, y int
			for {
				x = AskInt("Insert X position for character "+id+":", 0, utility.X_SIZE-1)
				y = AskInt("Insert Y position for character "+id+":", 0, utility.Y_SIZE-1)

				// check if the position is already occupied
				if !positions[utility.Point{X: x, Y: y}] {
					positions[utility.Point{X: x, Y: y}] = true
					break
				}
				log.Info().Msg("The position is already occupied")
			}

			player.AddCharacter(character.New(id, x, y, utility.COLORS[uint8(i)]))
		}

		log.Info().Msg(player.Print())
		playerList[i] = player
	}

	return Setup{playerList}
}

func (setup Setup) getCharacters() []character.Character {
	var characters []character.Character
	for _, player := range setup.Players {
		characters = append(characters, player.GetCharacters()...)
	}
	return characters
}
