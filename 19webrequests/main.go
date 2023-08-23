package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://tunelives.com/"

func main() {
	fmt.Println("Web requests in GoLang")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("response is of type %T \n ", response)
	fmt.Println(response)

	//User responsibility to close
	defer response.Body.Close()

	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("databyte in string are:", string(databytes))

}
