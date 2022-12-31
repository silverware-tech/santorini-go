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

func (p *Player) GetX() int {
	return p.x
}

func (p *Player) GetY() int {
	return p.y
}

func (p *Player) GetGroup() int {
	return p.group
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) Move(x int, y int) error {
	// player cannot stand still
	if p.x == x && p.y == y {
		return customError.PlayerMoveError{
			PlayerName: p.name,
			PlayerX:    p.x,
			PlayerY:    p.y,
			ErrorStr:   "Can not move to player's actual position",
		}
	}

	// player can move only to adjacent cells
	if x > p.x+1 || x < p.x-1 || y > p.y+1 || y < p.y-1 {
		return customError.PlayerMoveError{
			PlayerName: p.name,
			PlayerX:    x,
			PlayerY:    y,
			ErrorStr:   "Players can move only to adjacent cells",
		}
	}

	// player can move only to empty cells
	// TODO: check if cell is empty

	// player can move only to cells with height difference <= 1
	// TODO: check if cell height difference is <= 1

	p.x = x
	p.y = y
	return nil
}
