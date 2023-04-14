package game

import (
	"fmt"
	"strconv"

	"github.com/c2r0b/santorini.git/lib/character"
	"github.com/c2r0b/santorini.git/lib/player"
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
	Players []player.Player
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

func AskSetup() Setup {

	var numberOfPlayers = AskValue("Number of players:", MIN_PLAYERS, MAX_PLAYERS)

	var players = make([]player.Player, numberOfPlayers)

	// generate characters list (2 for each group)
	for i := 0; i < numberOfPlayers; i++ {
		players[i] = player.New("Player "+strconv.Itoa(i+1), false)

		for j := 0; j < MAX_CHARACTERS_PER_PLAYER; j++ {
			id := string('A' + i*MAX_CHARACTERS_PER_PLAYER + j)
			players[i].AddCharacter(character.New(id, i, j))
		}
	}
	fmt.Println(players)

	return Setup{players}
}

func (setup Setup) getCharacters() []character.Character {
	var characters []character.Character
	for _, player := range setup.Players {
		characters = append(characters, player.GetCharacters()...)
	}
	return characters
}
