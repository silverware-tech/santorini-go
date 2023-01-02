package game

import (
	"fmt"
	"santorini/main/pkg/board"
	"santorini/main/pkg/player"
	"strconv"
)

const SIZE = 5
const MAX_GROUPS = 4
const PLAYERS_PER_GROUP = 2
const MAX_PLAYERS = MAX_GROUPS * PLAYERS_PER_GROUP

type Game struct {
	players [MAX_PLAYERS]player.Player
	board   board.Board
}

func New(numberOfPlayers int) Game {
	var players [MAX_PLAYERS]player.Player

	// generate game board
	board := board.New()

	// generate characters list (2 for each group)
	for i := 0; i < numberOfPlayers*2; i++ {
		group := i / 2
		players[i] = player.New(0, 0, group, strconv.Itoa(i+1))
	}
	fmt.Print(players)
	g := Game{players, board}
	return g
}

func (g *Game) GetPlayers() []player.Player {
	return g.players[0 : MAX_PLAYERS]
}

// get players grouped by groups
func (g *Game) GetGroups() [][]player.Player {
	groups := make([][]player.Player, len(g.players)/2)
	for _, player := range g.players {
		groups[player.GetGroup()] = append(groups[player.GetGroup()], player)
	}
	return groups
}

func (g *Game) GetBoard() board.Board {
	return g.board
}

func (g *Game) IsOver() bool {
	// TODO: implement
	return false
}
