package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Mod in golang")
	greet()
	r := mux.NewRouter()
	// r.HandleFunc("/", serveHome)
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/next", servenext).Methods("GET")
	// server creations
	log.Fatal(http.ListenAndServe(":3000", r))
}

func greet() {
	fmt.Println("hello mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang module tutorial</h1>"))
}

func servenext(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> New version upcoming soon</h1>"))
}

// go list -m all
// go mod verify
// go mod tidy
// go list -m -versions github.com/gorilla/mux
// go build .
// go mod why github.com/gorilla/mux
// go mod graph
// go mod edit -go 1.20
// go mod edit -module modulename
// go mod vendor    --- the following adds a vendor folder(not recommended though)
// go run -mod=vendor main.go
