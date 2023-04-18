package character

type Character struct {
	X           int
	Y           int
	CharacterId string
	Color       string
}

func New(characterId string, X, Y int, Color string) Character {
	return Character{
		X, Y, characterId, Color,
	}
}
