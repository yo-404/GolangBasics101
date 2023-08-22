package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices ")
	// not defining size
	var fruit = []string{"apple", "orange", "banana"}
	fmt.Println("type of fruitlist data is %T ", fruit)

	fruit = append(fruit, "Mango", "Lichi")

	fmt.Println("fruit list is as follows \n", fruit)

	//slicing
	fruit = append(fruit[1:])
	fmt.Println("fruit list is as follows \n", fruit)
	// fruit = append(fruit[1:3])
	// fmt.Println("fruit list is as follows \n", fruit)
	// the last range is always non inclusive
	fruit = append(fruit[:3])
	fmt.Println("fruit list is as follows \n", fruit)

	highScores := make([]int, 4)

	highScores[0] = 13
	highScores[1] = 198
	highScores[2] = 654
	highScores[3] = 574

	fmt.Println(highScores)
	// adding adnother value
	// highScores[4] = 691

	highScores = append(highScores, 546, 486, 48)
	// error since index out of bound

	fmt.Println(highScores)

	fmt.Println("are integers sorted :", sort.IntsAreSorted(highScores))
	//these methods are available in slices and not in arrays
	sort.Ints(highScores)
	fmt.Println(highScores)

	fmt.Println("are integers sorted :", sort.IntsAreSorted(highScores))

	// to remove a value from slice based on indexing

	var courses = []string{"reactjs", "javascript", "java", "python"}
	fmt.Println("\n \n \n ", courses)

	// using append to remove the values
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

}
