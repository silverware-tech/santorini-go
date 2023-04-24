package EntityManager

import (
	"errors"
	"fmt"
	"github.com/c2r0b/santorini.git/lib/character"
	"github.com/c2r0b/santorini.git/lib/utility"
	"github.com/rs/zerolog/log"
)

type EntityManager struct {
	Board Board
}

func (m EntityManager) PrintBoard() {
	m.Board.Print()
}

func (m EntityManager) Move(character *character.Character, destination utility.Point) (bool, error) {
	if m.Board.IsValidMove(character.Position, destination) {
		m.Board.MoveCharacter(character, destination)
	} else {
		return false, errors.New("position not valid")
	}

	return m.Board.IsWinner(destination), nil
}

func New(xSize, ySize int, characters []character.Character) EntityManager {
	var m EntityManager
	m.Board = NewBoard(xSize, ySize)

	for _, c := range characters {
		c := c
		m.Board.SetCharacter(&c)
	}
	return m
}

func (m EntityManager) Build(character *character.Character, buildPoint utility.Point) error {
	if m.Board.IsValidBuild(character.Position, buildPoint) {
		m.Board.Build(buildPoint)
	} else {
		return errors.New(fmt.Sprintf("Build Position %s not valid", buildPoint.Print()))
	}
	return nil
}

func (m EntityManager) GetAvailableMove(position utility.Point) []utility.Point {
	var points = m.Board.GetNearPoints(position)
	var available []utility.Point

	for _, point := range points {
		if m.Board.IsValidMove(position, point) {
			available = append(available, point)
		}
	}

	return available
}

func (m EntityManager) GetAvailableBuild(characterPosition, movePosition utility.Point) []utility.Point {
	var points = m.Board.GetNearPoints(movePosition)
	var available []utility.Point
	log.Debug().Msgf("NearBuilds %v", points)
	for _, point := range points {
		// Character is not really moved on the board so the previous position is occupied by himself
		// and IsValidBuild will give false on a good point
		if m.Board.IsValidBuild(movePosition, point) || utility.Equals(characterPosition, point) {
			available = append(available, point)
		}
	}

	return available
}
