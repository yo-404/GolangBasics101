package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to basic control flow statements")
	fmt.Println("If else ")

	var totalMarks int
	// reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter marks")
	fmt.Scan(&totalMarks)

	var result string

	if totalMarks < 40 {
		result = "needs improvement"
	} else if totalMarks < 60 {
		result = "average"
	} else if totalMarks < 80 {
		result = "Good"
	} else {
		result = "excellent"
	}

	fmt.Println("result is :", result)

	if totalMarks%2 == 0 {
		fmt.Println("\n \n no.is even")
	} else {
		fmt.Println("\n no.is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("number is less than 10")
	} else {
		fmt.Println("Number is Not less than 10")
	}

	// for checking err

	// if err != nil {

	// }
}
