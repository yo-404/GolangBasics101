package main

import "fmt"

func main() {
	fmt.Println("Defer statements in Go lang")
	defer fmt.Println("world")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("hello")
	// defer uses LIFO order
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// output - hello ,43210,two one  world
// In Golang, the defer keyword is used to delay the execution of a function until the surrounding function completes. The deferred function calls are executed in Last-In-First-Out (LIFO) order. That means the most recently deferred function is executed first, followed by the second most recently deferred function, and so on.
// The defer keyword is useful when we want to ensure that certain operations are performed before a function returns, regardless of whether an error occurs or not. This can help simplify error handling and make code more readable.
