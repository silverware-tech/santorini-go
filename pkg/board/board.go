package board

import (
	"fmt"
	"santorini/main/pkg/board/cell"
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
			fmt.Print(b.field[i][j])
		}
		fmt.Println()
	}
}
