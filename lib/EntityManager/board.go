package EntityManager

import (
	"fmt"

	"github.com/c2r0b/santorini.git/lib/EntityManager/cell"
	"github.com/c2r0b/santorini.git/lib/character"
)

type Board struct {
	xSize int
	ySize int
	field [][]cell.Cell
}

// field cell getter
func (b Board) GetCell(x int, y int) *cell.Cell {
	return &b.field[x][y]
}

func NewBoard(xSize int, ySize int) Board {
	var b Board
	b.xSize = xSize
	b.ySize = ySize
	b.field = make([][]cell.Cell, b.xSize)
	for i := range b.field {
		b.field[i] = make([]cell.Cell, b.ySize)
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
	for i := 0; i < len(b.field); i++ {
		// Print top border for each row
		if i == 0 {
			fmt.Print("┌")
			for j := 0; j < len(b.field[0]); j++ {
				fmt.Print("────")
				if j != len(b.field[0])-1 {
					fmt.Print("┬")
				}
			}
			fmt.Println("┐")
		}

		// Print cells in the row with borders
		fmt.Print("│")
		for j := 0; j < len(b.field[0]); j++ {
			fmt.Printf(" %2s ", b.field[j][i].Print())
			if j != len(b.field[0])-1 {
				fmt.Print("│")
			}
		}
		fmt.Println("│")

		// Print bottom border for each row
		if i != len(b.field)-1 {
			fmt.Print("├")
			for j := 0; j < len(b.field[0]); j++ {
				fmt.Print("────")
				if j != len(b.field[0])-1 {
					fmt.Print("┼")
				}
			}
			fmt.Println("┤")
		} else {
			fmt.Print("└")
			for j := 0; j < len(b.field[0]); j++ {
				fmt.Print("────")
				if j != len(b.field[0])-1 {
					fmt.Print("┴")
				}
			}
			fmt.Println("┘")
		}
	}
}

func (b Board) IsNear(p *character.Character, x int, y int) bool {
	// player cannot stand still
	if p.X == x && p.Y == y {
		return false
	}

	// player can move only to adjacent cells
	if x > p.X+1 || x < p.X-1 || y > p.Y+1 || y < p.Y-1 {
		return false
	}

	return true
}

func (b Board) IsPositionValid(character *character.Character, x int, y int) bool {
	if !b.field[x][y].IsEmpty() {
		return false
	}
	if !b.IsNear(character, x, y) {
		return false
	}
	if b.field[character.X][character.Y].Height < b.field[x][y].Height-1 {
		return false
	}
	return true
}

func (b Board) MoveCharacter(character *character.Character, x int, y int) {
	b.field[character.X][character.Y].RemoveCharacter()
	b.field[x][y].SetCharacter(character)
	character.X = x
	character.Y = y
}

func (b Board) SetCharacter(character *character.Character) {
	b.field[character.X][character.Y].SetCharacter(character)
}

func (b Board) Build(x int, y int) error {
	return b.field[x][y].Build()
}
