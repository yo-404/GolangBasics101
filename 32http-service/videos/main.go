package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Starting server ....")
	http.HandleFunc("/get", HandleGetVideos)
	http.HandleFunc("/", Hellow)
	http.HandleFunc("/update", HandleUpdateVideos)

	http.ListenAndServe(":8080", nil)
}

func Hellow(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hellow and welcome to golang basics"))
}

func HandleGetVideos(w http.ResponseWriter, r *http.Request) {

	videos := getVideos()

	videoBytes, err := json.Marshal(videos)

	if err != nil {
		panic(err)
	}

	w.Write(videoBytes)
}

func HandleUpdateVideos(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		var videos []video
		err = json.Unmarshal(body, &videos)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "Bad Request")
		}

		saveVideos(videos)
	} else {
		w.WriteHeader(405)
		fmt.Fprintf(w, "Method not Supported!")
	}
}
