package main

import (
	"fmt"
	"santorini/main/pkg/player"
)

const SIZE = 5

func main() {
	var field [SIZE][SIZE]int

	player1 := player.New(0, 0, 0, "Player1")
	player2 := player.New(1, 0, 1, "Player2")
	player3 := player.New(2, 0, 2, "Player3")

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			field[i][j] = 0
		}
	}

	fmt.Println("Hello, Santorini!", player1, player2, player3)

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			fmt.Print(field[i][j])
		}
		fmt.Println()
	}

	x := 0
	for y := 0; y < 3; y++ {
		if y == 2 {
			y++
		}

		fmt.Printf("Move %v to position (%v,%v)\n", player1, x, y)

		err := player1.Move(x, y)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Player moved", player1)
		}
	}
}
