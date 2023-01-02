package cell

import (
	"santorini/main/pkg/player"
	"santorini/main/pkg/customError"
)

type CellStatus int
const (
	GROUND CellStatus = iota
	L1
	L2
	L3
	DOME
)

type Cell struct {
	Worker *player.Player
	Status   CellStatus
}

func New() Cell {
	return Cell{
		Worker: nil,
		Status: GROUND,
	}
}

func (c *Cell) buildOn() error {
	switch {
	case c.Worker != nil:
		return customError.CellBuildError{
			ErrorStr: "Can not build on occupied cell",
		}
	case c.Status == DOME:
		return customError.CellBuildError{
			ErrorStr: "Can not build on a complete tower",
		}
	case c.Status == GROUND:
		c.Status = L1
	case c.Status == L1:
		c.Status = L2
	case c.Status == L2:
		c.Status = L3
	case c.Status == L3:
		c.Status = DOME
	}

	return nil
}

func (c *Cell) setWorker(worker *player.Player) {
	c.Worker = worker
}
