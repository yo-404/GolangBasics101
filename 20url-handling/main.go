package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://github.com:8080/yo-404/botkube-docs/blob/spell-check/docusaurus.config.js/info?coursename=opensource&paymentid=49984asd849xzc"

func main() {
	fmt.Println(" url handling in golang")
	fmt.Println(myurl)

	// parsing the url
	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.User)

	queryparams := result.Query()
	fmt.Printf("the type of queryparam is : %T", queryparams)
	fmt.Println(queryparams)
	// to extract individual value
	fmt.Println("coursename:", queryparams["coursename"])

	for _, value := range queryparams {
		fmt.Println("params are", value)
	}

	// creating url from info chunks
	// referncing is important
	partsofUrl := &url.URL{
		Scheme:   "https",
		Host:     "github.com",
		Path:     "/yo-404",
		RawQuery: "/botkube-docs",
	}

	anotherurl := partsofUrl.String()
	fmt.Println("ur newly created url is : \n ", anotherurl)
}
