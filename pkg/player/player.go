package player

type Player struct {
	Name    string
	IsHuman bool
}

func New(Name string, IsHuman bool) Player {
	p := Player{Name, IsHuman}
	return p
}
