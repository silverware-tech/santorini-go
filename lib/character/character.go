package character

type Character struct {
	X           int
	Y           int
	CharacterId string
}

func New(characterId string, X, Y int) Character {
	return Character{
		X, Y, characterId,
	}
}
