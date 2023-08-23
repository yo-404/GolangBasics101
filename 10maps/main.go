package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps")
	// can be used to create slices and maps both (make)
	lang := make(map[string]string)

	lang["JS"] = "javascript"
	lang["RB"] = "ruby"
	lang["PY"] = "python"
	lang["MD"] = "MarkDown"

	fmt.Println("values in maps are", lang)

	fmt.Println("JS is", lang["JS"])
	fmt.Println("RB is", lang["RB"])

	// deleting values

	delete(lang, "RB")

	fmt.Println("values in maps are", lang)

	// loops
	//similar to foreach loop
	for key, value := range lang {
		fmt.Printf("key %v value is %v \n", key, value)
		// use printf while using the placeholders
	}

	for _, value := range lang {
		fmt.Printf("key value is %v \n", value)
		// use printf while using the placeholders
	}
}
