package main

import "fmt"

// walrus operator not allowed outside main
// jwtToken :=3000
var jwtToken int = 3000

// if the first letter of variable is Capital it is public
// if not then not public ..accessible inside file local
const LoginToken string = "165asd16wqd"

func main() {
	fmt.Println("Variables")
	var username string = "yo-404"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	var smallfloat float64 = 255.64589494
	fmt.Println(smallfloat)
	fmt.Printf("variable is of type: %T \n", smallfloat)

	// defailt values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n", anotherVariable)

	var EmptyString string
	fmt.Println(EmptyString)
	fmt.Printf("variable is of type: %T \n", EmptyString)

	// implicit type

	var name = "yogesh"
	fmt.Println(name)

	// no var style
	numberOfUser := 52
	fmt.Println(numberOfUser)

	fmt.Println(jwtToken)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n", LoginToken)
}
