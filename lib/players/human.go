package players

import (
	"errors"
	"fmt"
	"github.com/c2r0b/santorini.git/lib/EntityManager"
	"github.com/c2r0b/santorini.git/lib/character"
	"github.com/c2r0b/santorini.git/lib/utility"
	"github.com/rs/zerolog/log"
)

// This type implement the Player Interface

type Human struct {
	Name       string
	Characters []character.Character
}

func (human *Human) New(Name string) Player {
	return &Human{Name, make([]character.Character, 0)}
}

func (human *Human) NewWithCharacters(Name string, characters []character.Character) Player {
	return &Human{Name: Name, Characters: characters}
}

func (human *Human) GetName() string {
	return human.Name
}

func (human *Human) DoTurn(em EntityManager.EntityManager) (*character.Character, utility.Point, utility.Point, error) {
	em.PrintBoard()

	var character *character.Character
	var destPoint, buildPoint utility.Point
	for {
		character = human.ChooseCharacter()
		destPoint = utility.AskDestination()
		var isValidMove = em.Board.IsValidMove(character.Position, destPoint)
		if isValidMove {
			log.Info().Msgf("Move %v to position %s\n", character, destPoint)
			break
		} else {
			log.Error().Msgf("Unable to move %s in %s", character.CharacterId, destPoint)
		}
	}

	if em.Board.IsWinner(destPoint) {
		return character, destPoint, destPoint, nil
	}

	for {
		buildPoint = utility.AskBuild()
		var isValidBuild = em.Board.IsValidBuild(character.Position, buildPoint)
		if isValidBuild {
			log.Info().Msgf("Build to position %s\n", buildPoint)
			break
		}
	}

	return character, destPoint, buildPoint, nil
}

func (human *Human) GetCharacters() []character.Character {
	return human.Characters
}

func (human *Human) AddCharacter(character character.Character) {
	human.Characters = append(human.Characters, character)
}

func (human *Human) Print() string {
	return fmt.Sprintf("Player is Human with Name: %s, Characters: %s", human.Name, human.PrintCharacters())
}

func (human *Human) ChooseCharacter() *character.Character {
	var characterId string
	for {
		log.Info().Msg("Choose a character:")
		fmt.Scan(&characterId)
		for i := 0; i < len(human.Characters); i++ {
			if human.Characters[i].CharacterId == characterId {
				return &human.Characters[i]
			}
		}
		log.Info().Msg("Invalid character")
	}
}

func (human *Human) PrintCharacters() string {
	var characters string
	for i := 0; i < len(human.Characters); i++ {
		characters += human.Characters[i].CharacterId
		if i != len(human.Characters)-1 {
			characters += ", "
		}
	}
	return characters
}

func (human *Human) GetCharacter(characterId string) (*character.Character, error) {
	for i := 0; i < len(human.Characters); i++ {
		if human.Characters[i].CharacterId == characterId {
			return &human.Characters[i], nil
		}
	}
	return nil, errors.New("character not found")
}

func (human *Human) HasCharacter(characterId string) bool {
	for i := 0; i < len(human.Characters); i++ {
		if human.Characters[i].CharacterId == characterId {
			return true
		}
	}
	return false
}
