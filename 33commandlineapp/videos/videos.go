package main

import (
	"encoding/json"
	"os"
)

type video struct {
	Id          string `json:"id" `
	Title       string `json:"title"`
	Description string `json:"description"`
	Imageurl    string `json:"imageurl"`
	Url         string `json:"url"`
}

func getVideos() (videos []video) {
	// ioutils deprecated
	fileBytes, err := os.ReadFile("./videos.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &videos)
	if err != nil {
		panic(err)
	}

	return videos
}

func saveVideos(videos []video) {
	videoBytes, err := json.Marshal(videos)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./videos-updated.json", videoBytes, 0644)
	if err != nil {
		panic(err)
	}
}
