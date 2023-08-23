package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("File handling in go Lang \n \n ")
	content := "This is a sample texst to be written in a file "
	// to create a file
	file, err := os.Create("./myfile.txt")
	checkNilError(err)
	// to write into a file
	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("Length is:", length)
	defer file.Close()
	readFile("./myfile.txt")
}

// to read from a file
func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)

	fmt.Println("Textdata inside file is \n", string(databyte))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
