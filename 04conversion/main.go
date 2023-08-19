package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Conversions of variables/ streams")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your gaming experience from 1-5")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating us :", input)

	// lets add +1 to the rating given to us by the user
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	// using strings.TrimSpace to trim the \n (enter) from the input
	if err != nil {
		fmt.Println(err)
		// panic(err)
		// ends the program
	} else {
		fmt.Println("Added 1 to our rating from our side")
		fmt.Println("\n New rating is :", numRating+1)
	}
}
