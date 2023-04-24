package utility

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func (p Point) IsNear(dest Point) bool {
	// players cannot stand still and can move only to adjacent cells
	standStill := p == dest
	checkX := dest.X > p.X+1 || dest.X < p.X-1
	checkY := dest.Y > p.Y+1 || dest.Y < p.Y-1
	return !(standStill || checkX || checkY)
}

func (p Point) IsNotNear(dest Point) bool {
	return !p.IsNear(dest)
}

func (p Point) Print() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func AskDestination() Point {
	fmt.Print("Where to move ")
	dest := AskPoint()
	return dest
}

func AskBuild() Point {
	fmt.Print("Where to build ")
	dest := AskPoint()
	return dest
}

func AskPoint() Point {
	var x, y int
	fmt.Print("X:")
	fmt.Scan(&x)
	fmt.Print("Y:")
	fmt.Scan(&y)
	return Point{x, y}
}

func AddPoints(a, b Point) Point {
	return Point{a.X + b.X, a.Y + b.Y}
}
