package players

import (
	"errors"
	"fmt"
	"github.com/c2r0b/santorini.git/lib/EntityManager"
	"github.com/c2r0b/santorini.git/lib/character"
	"github.com/c2r0b/santorini.git/lib/utility"
	"github.com/rs/zerolog/log"
	"math/rand"
	"time"
)

// This type implement the Player Interface

type RandomAi struct {
	Name       string
	Characters []character.Character
	Rand       *rand.Rand
}

func (ai *RandomAi) New(Name string) Player {
	s := rand.NewSource(time.Now().Unix())
	return &RandomAi{
		Name,
		make([]character.Character, 0),
		rand.New(s),
	}
}

func (ai *RandomAi) NewWithCharacters(Name string, characters []character.Character) Player {
	s := rand.NewSource(time.Now().Unix())
	return &RandomAi{
		Name,
		characters,
		rand.New(s),
	}
}

func (ai *RandomAi) GetName() string {
	return ai.Name
}

func (ai *RandomAi) DoTurn(em EntityManager.EntityManager) (*character.Character, utility.Point, utility.Point, error) {
	em.PrintBoard()

	var selectedCharacter *character.Character
	var destPoint, buildPoint utility.Point

	var characterIndex = ai.Rand.Intn(len(ai.Characters))

	var i int
	for i = 0; i < utility.MAX_CHARACTERS_PER_PLAYER; i++ {
		selectedCharacter = &ai.Characters[(characterIndex+i)%utility.MAX_CHARACTERS_PER_PLAYER]
		log.Debug().Msgf("Character %s", selectedCharacter.CharacterId)

		movePoints := em.GetAvailableMove(selectedCharacter.Position)
		log.Debug().Msgf("Available moves %v", movePoints)

		// The selected player is currently unusable
		if len(movePoints) == 0 {
			log.Debug().Msg("No available moves for this character")
			continue
		}
		destPoint = movePoints[ai.Rand.Intn(len(movePoints))]
		log.Debug().Msgf("selected %s", destPoint.Print())
		break // The player is usable go to build phase
	}

	if i >= utility.MAX_CHARACTERS_PER_PLAYER {
		return selectedCharacter, selectedCharacter.Position, selectedCharacter.Position, errors.New("no available movements")
	}

	if em.Board.IsWinner(destPoint) {
		return selectedCharacter, destPoint, destPoint, nil
	}

	buildPoints := em.GetAvailableBuild(selectedCharacter.Position, destPoint)
	log.Debug().Msgf("Available build %v", buildPoints)
	buildPoint = buildPoints[ai.Rand.Intn(len(buildPoints))]
	log.Debug().Msgf("selected %s", buildPoint.Print())

	return selectedCharacter, destPoint, buildPoint, nil
}

func (ai *RandomAi) GetCharacters() []character.Character {
	return ai.Characters
}

func (ai *RandomAi) AddCharacter(character character.Character) {
	ai.Characters = append(ai.Characters, character)
}

/**

END Interface

*/

func (ai *RandomAi) Print() string {
	return fmt.Sprintf("Player is RandomAI with Name: %s, Characters: %s", ai.Name, ai.PrintCharacters())
}

func (ai *RandomAi) PrintCharacters() string {
	var characters string
	for i := 0; i < len(ai.Characters); i++ {
		characters += ai.Characters[i].CharacterId
		if i != len(ai.Characters)-1 {
			characters += ", "
		}
	}
	return characters
}
