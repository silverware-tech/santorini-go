package board

import (
	"fmt"
	"santorini/main/pkg/board/cell"
	"santorini/main/pkg/customError"
)

const SIZE = 5

type Board struct {
	field [SIZE][SIZE]cell.Cell
}

func New() Board {
	var b Board

	// generate game field
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			b.field[i][j] = cell.New()
		}
	}

	return b
}

func (b Board) Print() {
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			fmt.Print(b.field[i][j].Height)
		}
		fmt.Println()
	}
}

func (p *Player) Move(x int, y int) error {
	// player cannot stand still
	if p.x == x && p.y == y {
		return customError.PlayerMoveError{
			PlayerName: p.name,
			PlayerX:    p.x,
			PlayerY:    p.y,
			ErrorStr:   "Can not move to player's actual position",
		}
	}

	// player can move only to adjacent cells
	if x > p.x+1 || x < p.x-1 || y > p.y+1 || y < p.y-1 {
		return customError.PlayerMoveError{
			PlayerName: p.name,
			PlayerX:    x,
			PlayerY:    y,
			ErrorStr:   "Players can move only to adjacent cells",
		}
	}

	// player can move only to empty cells
	// TODO: check if cell is empty

	// player can move only to cells with height difference <= 1
	// TODO: check if cell height difference is <= 1

	p.x = x
	p.y = y
	return nil
}
