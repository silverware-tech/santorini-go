package EntityManager

import (
	"errors"
	"fmt"
	"github.com/c2r0b/santorini.git/lib/utility"

	"github.com/c2r0b/santorini.git/lib/character"
)

type EntityManager struct {
	Board Board
	// Character character.Character
}

func (m EntityManager) PrintBoard() {
	m.Board.Print()
}

func (m EntityManager) Move(character *character.Character, destination utility.Point) (bool, error) {
	if m.Board.IsValidMove(character.Position, destination) {
		m.Board.MoveCharacter(character, destination)
	} else {
		return false, errors.New("position not valid")
	}

	return m.Board.IsOver(destination), nil
}

func New(xSize, ySize int, characters []character.Character) EntityManager {
	var m EntityManager
	m.Board = NewBoard(xSize, ySize)

	for _, c := range characters {
		c := c
		m.Board.SetCharacter(&c)
	}
	return m
}

func (m EntityManager) Build(character *character.Character, buildPoint utility.Point) error {
	if m.Board.IsValidBuild(character.Position, buildPoint) {
		m.Board.Build(buildPoint)
	} else {
		return errors.New(fmt.Sprintf("Build Position %s not valid", buildPoint.Print()))
	}
	return nil
}
