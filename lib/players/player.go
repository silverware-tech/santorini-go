package players

import (
	"github.com/c2r0b/santorini.git/lib/EntityManager"
	"github.com/c2r0b/santorini.git/lib/character"
	"github.com/c2r0b/santorini.git/lib/utility"
)

type Player interface {
	Print() string

	GetName() string
	DoTurn(em EntityManager.EntityManager) (*character.Character, utility.Point, utility.Point)

	GetCharacters() []character.Character
	AddCharacter(character character.Character)
}
