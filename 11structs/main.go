package main

import "fmt"

func main() {
	fmt.Println("WElcome to structs in golang")
	// there is no inheritance in golang
	// there is no super,parent etc

	yogesh := User{"Yogesh", "a@x.com", "M", "23"}

	fmt.Println(yogesh)
	fmt.Printf("yogesh's detail are %+v \n", yogesh)
	fmt.Printf("name is %v \n and email is %v", yogesh.Name, yogesh.Email)
}

// capital first word ie publically accessible

type User struct {
	Name   string
	Email  string
	gender string
	Age    string
}
