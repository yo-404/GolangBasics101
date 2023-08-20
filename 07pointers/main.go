package main

import "fmt"

func main() {
	fmt.Println("welcome to pointers")
	// var ptr *int
	// fmt.Println("value of pointer is ", ptr)

	// & - used for referencing
	myNum := 23
	var ptr = &myNum

	fmt.Println("value of actual pointer is ", ptr)
	fmt.Println("value of actual pointer is ", *ptr)

	// *ptr = value inside it
	*ptr = *ptr + 2
	fmt.Println("new val is", myNum)
}
