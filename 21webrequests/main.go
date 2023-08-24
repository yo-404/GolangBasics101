package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	// get Request
	fmt.Println("web requests in GO lang")
	performGetRequest()

}

func performGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code :", response.StatusCode)
	fmt.Println("Content length is :", response.ContentLength)

	var responsestring strings.Builder
	content, err := io.ReadAll(response.Body)
	byteCount, _ := responsestring.Write(content)
	if err != nil {
		panic(err)
	}
	// byte response is still stored into response string
	fmt.Println("Bytecount is ", byteCount)
	fmt.Println(responsestring.String())
	// content stored in byte format
	// fmt.Println(string(content))
	// using string builder

}
