package EntityManager

import (
	"fmt"
	"github.com/c2r0b/santorini.git/lib/utility"

	"github.com/c2r0b/santorini.git/lib/EntityManager/cell"
	"github.com/c2r0b/santorini.git/lib/character"
)

type Board struct {
	xSize int
	ySize int
	field [][]cell.Cell
}

// field cell getter
func (b Board) GetCell(point utility.Point) *cell.Cell {
	if b.IsOutOfBound(point) {
		log.Panic().Msgf("Point %s is out of bound", point)
		unix.Exit(-1)
	}
	return &b.field[point.X][point.Y]
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

// IsValidMove check if the moveDest is a valid position, to be a valid position the destination must be
// in the board bounds, near the player, empty (no other player present) and the increase in height must be at max 1
func (b Board) IsValidMove(characterPosition, moveDest utility.Point) bool {
	if b.IsOutOfBound(moveDest) ||
		!characterPosition.IsNear(moveDest) ||
		!b.GetCell(moveDest).IsEmpty() ||
		b.GetCell(characterPosition).Height < b.GetCell(moveDest).Height-1 {
		return false
	}
	return true
}

/*
IsValidBuild check if the buildPosition is a valid position. A valid build position match some requirement:
- the destination must be in the board bounds
- near the player
- empty (no other player present)
- there is no DOME
*/
func (b Board) IsValidBuild(characterPosition, buildPosition utility.Point) bool {
	if b.IsOutOfBound(buildPosition) ||
		characterPosition.IsNotNear(buildPosition) ||
		b.GetCell(buildPosition).IsNotEmpty() ||
		b.GetCell(buildPosition).Height == cell.DOME {
		return false
	}
	return true
}

func (b Board) MoveCharacter(character *character.Character, moveDest utility.Point) {
	b.GetCell(character.Position).RemoveCharacter()
	b.GetCell(moveDest).SetCharacter(character)
	character.Position = moveDest
}

func (b Board) SetCharacter(character *character.Character) {
	b.GetCell(character.Position).SetCharacter(character)
}

func (b Board) Build(build utility.Point) {
	b.GetCell(build).Build()
}

func (b Board) IsOutOfBound(point utility.Point) bool {
	return point.X < 0 || point.X > b.xSize-1 ||
		point.Y < 0 || point.Y > b.ySize-1
}

func (b Board) IsOver(destination utility.Point) bool {
	return b.GetCell(destination).Height == cell.L3
}
