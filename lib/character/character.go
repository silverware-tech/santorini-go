package character

import "github.com/c2r0b/santorini.git/lib/utility"

type Character struct {
	Position    utility.Point
	CharacterId string
	Color       string
}

func New(characterId string, X, Y int, color string) Character {
	return Character{
		Position:    utility.Point{X: X, Y: Y},
		CharacterId: characterId,
		Color:       color,
	}
}
