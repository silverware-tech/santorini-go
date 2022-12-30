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

func (p Player) Move(x int, y int) bool {
	// check if move is valid
	if p.x == x && p.y == y {
		return false
	}
	if x > p.x+1 || x < p.x-1 {
		return false
	}
	if y > p.y+1 || y < p.y-1 {
		return false
	}
	p.x = x
	p.y = y
	return true
}
