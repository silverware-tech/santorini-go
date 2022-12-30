package player

import "santorini/main/pkg/customError"

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

func (p *Player) Move(x int, y int) error {
	// check if move is valid
	if p.x == x && p.y == y {
		return customError.PlayerMoveError{
			PlayerName: p.name,
			PlayerX: p.x,
			PlayerY: p.y,
			ErrorStr: "Can not move to player's actual position",
		}
	}

	if x > p.x+1 || x < p.x-1 || y > p.y+1 || y < p.y-1 {
		return customError.PlayerMoveError{
			PlayerName: p.name,
			PlayerX: x,
			PlayerY: y,
			ErrorStr: "Players can move only to adjacent cells",
		}
	}

	p.x = x
	p.y = y
	return nil
}
