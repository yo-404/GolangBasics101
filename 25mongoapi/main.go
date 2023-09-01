package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yo-404/GolangBasics/25mongoapi/router"
)

func main() {
	fmt.Println("Mongo API")
	r := router.Router()
	fmt.Println("Server is running...")
	http.ListenAndServe(":4000", r)
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening to port 4000")
}
