package main

import "fmt"

const SIZE = 5

type Player struct {
	x int
	y int
	group int
	name string
}



func main() {
	var field[SIZE][SIZE]int
	player1 := Player{
		x:0,
		y:0,
		group: 0,
        name: "Player1",
    }
    player2 := Player{
        x:1,
        y:0,
        group: 1,
        name: "Player2",
    }
    player3 := Player{
        x:2,
        y:0,
        group: 2,
        name: "Player3",
    }

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
}