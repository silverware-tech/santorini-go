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

func (c *Cell) IsEmpty() bool {
	return c.Character == nil
}

// remove character from cell
func (c *Cell) RemoveCharacter() {
	c.Character = nil
}

// set character on cell
func (c *Cell) SetCharacter(Character *character.Character) {
	c.Character = Character
}

func (c *Cell) Print() string {
	var s string = ""

	if c.Character != nil {
		s += c.Character.CharacterId
	} else {
		s += " "
	}

	switch c.Height {
	case GROUND:
		s += "-"
	case L1:
		s += "1"
	case L2:
		s += "2"
	case L3:
		s += "3"
	case DOME:
		s += "4"
	}

	return s
}

// build
func (c *Cell) Build() error {
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
