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

type Ai struct {
	Name       string
	Characters []character.Character
}

func NewAi(Name string) *Ai {
	ai := Ai{Name, make([]character.Character, 0)}
	return &ai
}

func (ai *Ai) GetName() string {
	return ai.Name
}

func (ai *Ai) DoTurn(em EntityManager.EntityManager) (*character.Character, utility.Point, utility.Point) {
	log.Info().Msgf("Start AI")
	em.PrintBoard()

	var character *character.Character
	var destPoint, buildPoint utility.Point
	for {
		character = ai.ChooseCharacter()
		destPoint = utility.AskDestination()
		var isValidMove = em.Board.IsValidMove(character.Position, destPoint)
		if isValidMove {
			log.Info().Msgf("Move %v to position %s\n", character, destPoint)
			break
		}
	}

	if em.Board.IsOver(destPoint) {
		return character, destPoint, destPoint
	}

	for {
		buildPoint = utility.AskBuild()
		var isValidBuild = em.Board.IsValidBuild(character.Position, buildPoint)
		if isValidBuild {
			log.Info().Msgf("Build to position %s\n", buildPoint)
			break
		}
	}

	return character, destPoint, buildPoint
}

func (ai *Ai) GetCharacters() []character.Character {
	return ai.Characters
}

func (ai *Ai) AddCharacter(character character.Character) {
	ai.Characters = append(ai.Characters, character)
}

func (ai *Ai) ChooseCharacter() *character.Character {
	var characterId string
	for {
		fmt.Print("Choose a character: ")
		fmt.Scan(&characterId)
		for i := 0; i < len(ai.Characters); i++ {
			if ai.Characters[i].CharacterId == characterId {
				return &ai.Characters[i]
			}
		}
		log.Info().Msg("Invalid character")
	}
}

func (ai *Ai) Print() string {
	return fmt.Sprintf("Player is AI with Name: %s, Characters: %s", ai.Name, ai.PrintCharacters())
}

func (ai *Ai) PrintCharacters() string {
	var characters string
	for i := 0; i < len(ai.Characters); i++ {
		characters += ai.Characters[i].CharacterId
		if i != len(ai.Characters)-1 {
			characters += ", "
		}
	}
	return characters
}

func (ai *Ai) GetCharacter(characterId string) (*character.Character, error) {
	for i := 0; i < len(ai.Characters); i++ {
		if ai.Characters[i].CharacterId == characterId {
			return &ai.Characters[i], nil
		}
	}
	return nil, errors.New("Character not found")
}

func (ai *Ai) HasCharacter(characterId string) bool {
	for i := 0; i < len(ai.Characters); i++ {
		if ai.Characters[i].CharacterId == characterId {
			return true
		}
	}
	return false
}
