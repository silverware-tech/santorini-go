package main

import (
	"fmt"
	"santorini/main/pkg/game"
)

func main() {
	var numberOfPlayers int

	fmt.Println("Hello, Santorini!")

	fmt.Print("Number of players:")
	fmt.Scan(&numberOfPlayers)

	game := game.New(numberOfPlayers)

	for !game.IsOver() {
		groups := game.GetGroups()
		for i := 0; i < len(groups); {
			var characterToMove int
			var newX, newY int

			group := groups[i]

			game.GetBoard().Print()
			fmt.Println("Player", group[0].GetGroup(), "turn")

			fmt.Printf("What character to move (%v,%v):", group[0].GetName(), group[1].GetName())
			fmt.Scan(&characterToMove)
			player := group[characterToMove-1]

			fmt.Print("Where to move (X):")
			fmt.Scan(&newX)
			fmt.Print("Where to move (Y):")
			fmt.Scan(&newY)

			fmt.Printf("Move %v to position (%v,%v)\n", player, newX, newY)

			err := player.Move(newX, newY)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Player moved", player)
				i++
			}
		}
	}
}
