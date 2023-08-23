package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch case \n")

	// rand.Seed(time.Now().UnixNano())
	diceNum := rand.Intn(6) + 1
	//ranges are non inclusive

	fmt.Println("value of dice is ", diceNum)

	switch diceNum {
	case 1:
		fmt.Println("Dice value 1 and u can start ur game")
	case 2:
		fmt.Println("You can move 2 spots")
	case 3:
		fmt.Println("You can move 3 spots")
	case 4:
		fmt.Println("You can move 4 spots")
		fallthrough
		// also continues to the next case using fallthrough
	case 5:
		fmt.Println("You can move 5 spots")
	case 6:
		fmt.Println("You can move 6 spots and roll the dice again")
	default:
		fmt.Println("Invalid move ")
	}
}
