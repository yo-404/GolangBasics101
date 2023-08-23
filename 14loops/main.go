package main

import "fmt"

func main() {
	fmt.Println("loops in golang")

	days := []string{"Sunday", "monday", "tuesday", "wednesday"}
	fmt.Println("days are ", days)
	// looping through
	//for
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
	fmt.Println("\n")

	for i := range days {
		fmt.Println(days[i])
	}

	for index, days := range days {
		fmt.Printf("index is %v and value is %v \n", index, days)
	}

	for _, days := range days {
		fmt.Printf("Day is %v \n", days)
	}

	fmt.Println("\n \n ")
	value := 1
	// similar to while loop
	for value < 10 {

		if value == 7 {
			goto jump
		}

		if value == 5 {
			value++
			continue
			// break
			// break to stop the loop
			// continue skips over 1 interation
		}

		fmt.Println("value is :", value)
		value++
	}

	// goto statements

jump:
	fmt.Println("jump statments using goto")

}
