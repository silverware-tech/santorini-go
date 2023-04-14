package EntityManager

import (
	"errors"

	"github.com/c2r0b/santorini.git/lib/character"
)

type EntityManager struct {
	Board     Board
	Character character.Character
}

func (m EntityManager) PrintBoard() {
	m.Board.Print()
}

func (m EntityManager) Move(character *character.Character, x int, y int) error {
	if (x > m.Board.xSize) || (y > m.Board.ySize) {
		return errors.New("out of bounds")
	}
	if m.Board.IsPositionValid(character, x, y) {
		m.Board.MoveCharacter(character, x, y)
	} else {
		return errors.New("position not valid")
	}
	return nil
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

func (m EntityManager) Build(character *character.Character, x, y int) error {
	if (x > m.Board.xSize) || (y > m.Board.ySize) {
		return errors.New("out of bounds")
	}
	// check if near
	if !m.Board.IsNear(character, x, y) {
		return errors.New("position not near")
	}
	return m.Board.Build(x, y)
}
