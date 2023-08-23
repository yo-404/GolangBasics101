package main

import "fmt"

// main functions acts as an entry point for the program
func main() {
	fmt.Println("Functions")
	greet()

	// functions inside function are not allowed
	// func greet1() {
	// 	fmt.Println("Another method")
	// }

	// anonymous function example
	func() {
		fmt.Println("anonymous function")
	}()

	// ******************************
	greet1()

	result := adder(3, 5)
	fmt.Println("result is ", result)

	prores := proAdder(3, 5, 4, 9, 6, 4, 5)
	fmt.Println("proadder result is ", prores)

	resadd, myMessage := proADDy(3, 8, 4, 9, 47, 6, 3)
	fmt.Println("result for addition of all values is :", resadd)
	fmt.Println("message recieved is", myMessage)

}

func greet1() {
	fmt.Println("Another method")
}

func greet() {
	fmt.Println("Hello !! have a nice day ahead")
}

// using functions signature here ☻ to determine the value to be passed back onto the function
func adder(val1 int, val2 int) int {
	return val1 + val2
}

//  3 ... are  variadic function is a function of indefinite arity, i.e., one which accepts a variable number of arguments.
func proAdder(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func proADDy(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "added all the values which were provided"
}
