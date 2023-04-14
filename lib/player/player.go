package player

import (
	"errors"

	"github.com/c2r0b/santorini.git/lib/character"
)

type Player struct {
	Name       string
	IsHuman    bool
	Characters []character.Character
}

func New(Name string, IsHuman bool) Player {
	p := Player{Name, IsHuman, make([]character.Character, 0)}
	return p
}

func (p *Player) AddCharacter(character character.Character) {
	p.Characters = append(p.Characters, character)
}

func (p *Player) PrintCharacters() string {
	var characters string
	for i := 0; i < len(p.Characters); i++ {
		characters += p.Characters[i].CharacterId
		if i != len(p.Characters)-1 {
			characters += ", "
		}
	}
	return characters
}

func (p *Player) GetCharacter(characterId string) (*character.Character, error) {
	for i := 0; i < len(p.Characters); i++ {
		if p.Characters[i].CharacterId == characterId {
			return &p.Characters[i], nil
		}
	}
	return nil, errors.New("Character not found")
}

func (p *Player) HasCharacter(characterId string) bool {
	for i := 0; i < len(p.Characters); i++ {
		if p.Characters[i].CharacterId == characterId {
			return true
		}
	}
	return false
}

func (p *Player) GetCharacters() []character.Character {
	return p.Characters
}
