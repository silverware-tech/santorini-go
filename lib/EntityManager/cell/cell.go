package cell

import (
	"github.com/c2r0b/santorini.git/lib/character"
	"github.com/c2r0b/santorini.git/lib/customError"
)

type CellStatus int

const (
	GROUND CellStatus = iota
	L1
	L2
	L3
	DOME
)

type Cell struct {
	Character *character.Character
	Height    CellStatus
}

func New() Cell {
	return Cell{
		Character: nil,
		Height:    GROUND,
	}
}

func (c *Cell) buildOn() error {
	switch {
	case c.Character != nil:
		return customError.CellBuildError{
			ErrorStr: "Can not build on occupied cell",
		}
	case c.Height == DOME:
		return customError.CellBuildError{
			ErrorStr: "Can not build on a complete tower",
		}
	case c.Height == GROUND:
		c.Height = L1
	case c.Height == L1:
		c.Height = L2
	case c.Height == L2:
		c.Height = L3
	case c.Height == L3:
		c.Height = DOME
	}

	return nil
}

func (c *Cell) setWorker(Character *character.Character) {
	c.Character = Character
}

func (c *Cell) removeWorker() {
	c.Character = nil
}

func (c *Cell) getCharacter() *character.Character {
	return c.Character
}

func (c *Cell) getHeight() CellStatus {
	return c.Height
}
