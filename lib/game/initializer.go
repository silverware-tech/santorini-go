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
	var value bool
	for {
		fmt.Print("Is player " + name + " ai? (true/false): ")
		scan, err := fmt.Scan(&value)
		if scan != 1 || err != nil {
			log.Error().Msg("Error during scan")
		} else {
			break
		}
	}
	return value
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
			player.AddCharacter(character.New(id, i, j))
		}

		log.Info().Msg(fmt.Sprint(player.Print()))
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
