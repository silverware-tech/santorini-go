package game

import (
	"fmt"

	"github.com/c2r0b/santorini.git/lib/character"
	"github.com/c2r0b/santorini.git/lib/players"
	"github.com/rs/zerolog/log"
)

/**
Oggetto per la costruzione iniziale del gioco.
In AskSetup bisogna mettere tutte quelle cose di input iniziali:
- Quanti giocatori
- Nomi giocatori
- Se sono umani o AI
- Posizione dei character sulla mappa per ogni giocatore
*/

type Setup struct {
	Players []players.Player
}

var colors = map[uint8]string{
	0: "\x1b[1;31m",
	1: "\x1b[1;34m",
	2: "\x1b[1;32m",
	3: "\x1b[1;33m",
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
	var numberOfPlayers = AskInt("Number of players: ", MIN_PLAYERS, MAX_PLAYERS)

	var playerList = make([]players.Player, numberOfPlayers)

	// hashset for the positions
	var positions = make(map[utility.Point]bool)

	// generate characters list (2 for each group)
	for i := 0; i < numberOfPlayers; i++ {
		var playerName = AskName()
		var isIa = AskAi(playerName)

		var player players.Player
		if isIa {
			player = players.NewAi(playerName)
		} else {
			player = players.NewHuman(playerName)
		}
		for j := 0; j < MAX_CHARACTERS_PER_PLAYER; j++ {
			id := string(rune('A' + (i*MAX_CHARACTERS_PER_PLAYER + j)))

			// ask for unique position X and Y for the character on the board
			var x, y int
			for {
				x = AskInt("Insert X position for character "+id+":", 0, X_SIZE-1)
				y = AskInt("Insert Y position for character "+id+":", 0, Y_SIZE-1)

				// check if the position is already occupied
				if !positions[utility.Point{x, y}] {
					positions[utility.Point{x, y}] = true
					break
				}
				log.Info().Msg("The position is already occupied")
			}

			player.AddCharacter(character.New(id, x, y, colors[uint8(i)]))
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
