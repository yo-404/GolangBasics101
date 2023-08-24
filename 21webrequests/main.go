package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// get Request
	fmt.Println("web requests in GO lang")
	// performGetRequest()
	// performPostJsonRequest()
	performPostFormOperation()

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

func performPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	// fake json payload

	requestBody := strings.NewReader(`
		{
			"name":"Yogesh",
			"Age":"23",
			"OS":"linux"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	// webserver sends back the same response sent by the user
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}

func performPostFormOperation() {
	const myurl = "http://localhost:8000/postform"

	// fake form data

	data := url.Values{}
	data.Add("firstname", "yogesh")
	data.Add("location", "pune")
	data.Add("age", "23")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}
