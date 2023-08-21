package main

import "fmt"

func main() {
	fmt.Println("welcome to arrays")
	var friends [4]string

	friends[0] = "John"
	friends[1] = "Jack"
	// friends[2] = "Joey"
	friends[3] = "Joanne"
	// friends[4] = "jean"

	fmt.Println("list of friends are :", friends)
	fmt.Println("length of the list is", len(friends))

	var fruits = [5]string{"apple", "banana", "mango"}
	fmt.Println("fruits is is as follows :", fruits)
	fmt.Println("length of fruit list is:", len(fruits))
}
