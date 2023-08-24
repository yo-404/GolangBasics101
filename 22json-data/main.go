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
	// EncodeJSON()
	DecodeJSON()

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

func DecodeJSON() {

	jsonData := []byte(`
	{
		"studentname": "Yogesh",
		"Age": 23,
		"Language": "GO",
		"tags": ["student","fresher"]
	}

	`)
	var students details
	checkValid := json.Valid(jsonData)
	if checkValid {
		fmt.Println("json was valid")
		json.Unmarshal(jsonData, &students)
		fmt.Printf("%#v\n", students)
	} else {
		fmt.Println("json was not valid")
	}

	// there are some cases where u just want to add data to key value pair

	var mydataMap map[string]interface{}
	json.Unmarshal(jsonData, &mydataMap)
	fmt.Printf("%#v\n", mydataMap)
	fmt.Printf("\n\n")

	// order is not guaranteed in map
	for key, value := range mydataMap {
		fmt.Printf("Key is %v and value is %v and type is %T\n", key, value, value)
	}

}
