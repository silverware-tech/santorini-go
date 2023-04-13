package EntityManager

import (
	"github.com/c2r0b/santorini.git/lib/character"
	"github.com/rs/zerolog/log"
)

type EntityManager struct {
	Board     Board
	Character character.Character
}

func (m EntityManager) PrintBoard() {
	m.Board.Print()
}

func (m EntityManager) Move(x int, y int) error {
	log.Info().Msg("TODO")

	return nil
}

func New(xSize, ySize int, characters []character.Character) EntityManager {

	return EntityManager{Board: NewBoard(xSize, ySize)}
}
