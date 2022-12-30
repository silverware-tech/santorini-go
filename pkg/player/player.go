package player

type Player struct {
	x     int
	y     int
	group int
	name  string
}

func New(x int, y int, group int, name string) Player {
	p := Player{x, y, group, name}
	return p
}

func (p Player) Move() {

}
