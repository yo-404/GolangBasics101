package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// pointers in nature
var wg sync.WaitGroup

func main() {
	// using go keyword to invoke go routines
	// go greet("hello")
	// greet("world")
	websiteList := []string{
		"https://google.com",
		"https://tunelives.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		wg.Add(1)
		getStatusCode(web)

	}
	wg.Wait()
}

// func greet(s string) {
// 	for i := 0; i < 10; i++ {
// 		time.Sleep(20 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	result, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n %d 200 status code for website %s", result.StatusCode, endpoint)

}
