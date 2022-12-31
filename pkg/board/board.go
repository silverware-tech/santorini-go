package board

import (
	"fmt"
)

const SIZE = 5

type Board struct {
	field [SIZE][SIZE]int
}

func New() Board {
	var field [SIZE][SIZE]int

	// generate game field
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			field[i][j] = 0
		}
	}

	b := Board{field}
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
