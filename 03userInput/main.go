package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to GOlang tuts"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our struggling life:")

	//comma okay syntax or error ok syntax
	//similar to try catch

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating ", input)
	fmt.Printf("type of this rating is %T", input)

	input1, err := reader.ReadByte()
	fmt.Println("thanks for rating ", input1)
	fmt.Printf("type of this rating is %T", input1)
	fmt.Println("\n Error faced here ", err)

}
