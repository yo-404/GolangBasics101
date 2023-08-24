package main

import (
	"encoding/json"
	"fmt"
)

type details struct {
	// using json alias
	Name string `json:"studentname"`
	Age  int
	// - hides info
	Email    string `json:"-"`
	Language string
	// omits empty tags (no space between ',')
	Tags []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Json data handling")
	EncodeJSON()

}

func EncodeJSON() {

	Student := []details{
		{"Yogesh", 23, "yogesh@gmail.com", "GO", []string{"student", "fresher"}},
		{"Chirag", 23, "chirrya@gmail.com", "RUBY", nil},
		{"Anurag", 23, "bunny@gmail.com", "React", []string{"student", "fresher"}},
	}

	// packing the data as JSON

	// finalJson, err := json.Marshal(Student)
	finalJson, err := json.MarshalIndent(Student, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJson))
}
