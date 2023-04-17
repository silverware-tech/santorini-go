package EntityManager

import (
	"fmt"
	"github.com/c2r0b/santorini.git/lib/EntityManager"
	"github.com/c2r0b/santorini.git/lib/EntityManager/cell"
	"github.com/c2r0b/santorini.git/lib/character"
	"github.com/c2r0b/santorini.git/lib/utility"
	"testing"
)

func TestIsValidMove(t *testing.T) {

	var board = EntityManager.NewBoard(5, 5)
	board.SetCharacter(&character.Character{CharacterId: "A", Position: utility.Point{2, 2}})
	board.SetCharacter(&character.Character{CharacterId: "B", Position: utility.Point{2, 1}})
	board.SetCharacter(&character.Character{CharacterId: "C", Position: utility.Point{0, 0}})
	board.GetCell(utility.Point{2, 0}).Height = cell.L3
	// Defining the columns of the table
	var tests = []struct {
		startPoint utility.Point
		endPoint   utility.Point
		expected   bool
	}{
		// the table itself
		{utility.Point{2, 2}, utility.Point{2, 0}, false},
		{utility.Point{2, 2}, utility.Point{2, 1}, false},
		{utility.Point{0, 0}, utility.Point{-1, 0}, false},
		{utility.Point{2, 2}, utility.Point{3, 3}, true},
		{utility.Point{2, 2}, utility.Point{3, 4}, false},
	}
	// The execution loop
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Check IsValidMove %d", i), func(t *testing.T) {
			ans := board.IsValidMove(tt.startPoint, tt.endPoint)
			if ans != tt.expected {
				t.Errorf("got %t, want %t", ans, tt.expected)
			}
		})
	}
}

func TestIsValidBuild(t *testing.T) {

	var board = EntityManager.NewBoard(5, 5)
	board.SetCharacter(&character.Character{CharacterId: "A", Position: utility.Point{2, 2}})
	board.SetCharacter(&character.Character{CharacterId: "B", Position: utility.Point{2, 1}})
	board.SetCharacter(&character.Character{CharacterId: "C", Position: utility.Point{0, 0}})

	// Defining the columns of the table
	var tests = []struct {
		startPoint utility.Point
		endPoint   utility.Point
		expected   bool
	}{
		// the table itself
		{utility.Point{2, 2}, utility.Point{2, 0}, false},
		{utility.Point{2, 2}, utility.Point{2, 1}, false},
		{utility.Point{0, 0}, utility.Point{-1, 0}, false},
		{utility.Point{2, 2}, utility.Point{3, 3}, true},
		{utility.Point{2, 2}, utility.Point{3, 4}, false},
	}
	// The execution loop
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Check IsValidBuild %d", i), func(t *testing.T) {
			ans := board.IsValidBuild(tt.startPoint, tt.endPoint)
			if ans != tt.expected {
				t.Errorf("got %t, want %t", ans, tt.expected)
			}
		})
	}
}
