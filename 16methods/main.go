package main

import "fmt"

func main() {
	fmt.Println("WElcome to methods")
	// there is no inheritance in golang
	// there is no super,parent etc

	yogesh := User{"Yogesh", "a@x.com", "M", "23"}

	fmt.Println(yogesh)
	fmt.Printf("yogesh's detail are %+v \n", yogesh)
	fmt.Printf("name is %v \n and email is %v", yogesh.Name, yogesh.Email)
	fmt.Println("\n \n \n *****************")
	yogesh.GetStatus()
	yogesh.newMail()
	// doent change value in real struct as copy is being used . to change the mail in real structure u might have to use pointer
	fmt.Printf("name is %v \n and email is %v", yogesh.Name, yogesh.Email)
}

// capital first word ie publically accessible

type User struct {
	Name   string
	Email  string
	gender string
	Age    string
}

// use G as capital if u want to export
func (u User) GetStatus() {
	fmt.Println("gender of user:", u.gender)
}

func (u User) newMail() {
	u.Email = "test@xyz.com"
	fmt.Println("Email of user is", u.Email)
}
