package players

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/c2r0b/santorini.git/lib/EntityManager"
	"github.com/c2r0b/santorini.git/lib/character"
	"github.com/c2r0b/santorini.git/lib/utility"

	"github.com/rs/zerolog/log"
)

// This type implement the Player Interface

type RandomAi struct {
	Name       string
	Characters []character.Character
	Rand       *rand.Rand
}

func NewRandomAi(Name string) *RandomAi {
	s := rand.NewSource(time.Now().Unix())
	ai := RandomAi{
		Name,
		make([]character.Character, 0),
		rand.New(s),
	}
	return &ai
}

func (ai *RandomAi) GetName() string {
	return ai.Name
}

func (ai *RandomAi) DoTurn(em EntityManager.EntityManager) (*character.Character, utility.Point, utility.Point) {
	log.Info().Msgf("Start AI")
	em.PrintBoard()

	var selectedCharacter *character.Character
	var destPoint, buildPoint utility.Point

	selectedCharacter = &ai.Characters[ai.Rand.Intn(len(ai.Characters))]
	log.Debug().Msgf("Character %s", selectedCharacter.CharacterId)

	movePoints := em.GetAvailableMove(selectedCharacter.Position)
	log.Debug().Msgf("Available moves %v", movePoints)
	destPoint = movePoints[ai.Rand.Intn(len(movePoints))]
	log.Debug().Msgf("selected %s", destPoint.Print())

	if em.Board.IsOver(destPoint) {
		return selectedCharacter, destPoint, destPoint
	}

	buildPoints := em.GetAvailableBuild(destPoint)
	log.Debug().Msgf("Available build %v", buildPoints)
	buildPoint = buildPoints[ai.Rand.Intn(len(buildPoints))]
	log.Debug().Msgf("selected %s", buildPoint.Print())

	return selectedCharacter, destPoint, buildPoint
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
