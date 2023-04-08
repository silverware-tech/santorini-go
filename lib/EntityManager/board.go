package EntityManager

import (
	"fmt"
	"github.com/c2r0b/santorini.git/lib/EntityManager/cell"
	"github.com/c2r0b/santorini.git/lib/character"
	"github.com/c2r0b/santorini.git/lib/customError"
)

type Board struct {
	xSize int
	ySize int
	field [][]cell.Cell
}

func NewBoard(xSize int, ySize int) Board {
	var b Board
	b.xSize = xSize
	b.ySize = ySize
	b.field = make([][]cell.Cell, b.ySize)
	for i := range b.field {
		b.field[i] = make([]cell.Cell, b.xSize)
	}
	// generate game field
	for i := 0; i < b.xSize; i++ {
		for j := 0; j < b.ySize; j++ {
			b.field[i][j] = cell.New()
		}
	}

	return b
}

func (b Board) Print() {
	for i := 0; i < b.xSize; i++ {
		for j := 0; j < b.ySize; j++ {
			fmt.Print(b.field[i][j].Height)
		}
		fmt.Println()
	}
}

func Move(p character.Character, x int, y int) error {
	// player cannot stand still
	if p.X == x && p.Y == y {
		return customError.PlayerMoveError{
			PlayerName: "p.name",
			PlayerX:    p.X,
			PlayerY:    p.Y,
			ErrorStr:   "Can not move to player's actual position",
		}
	}

	// player can move only to adjacent cells
	if x > p.X+1 || x < p.X-1 || y > p.Y+1 || y < p.Y-1 {
		return customError.PlayerMoveError{
			PlayerName: "p.name",
			PlayerX:    x,
			PlayerY:    y,
			ErrorStr:   "Players can move only to adjacent cells",
		}
	}

	// player can move only to empty cells
	// TODO: check if cell is empty

	// player can move only to cells with height difference <= 1
	// TODO: check if cell height difference is <= 1

	p.X = x
	p.Y = y
	return nil
}
