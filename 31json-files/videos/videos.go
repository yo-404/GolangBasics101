package main

import (
	"encoding/json"
	"os"
)

// mapping the struct/slice  with json data
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

	// fileContent := string(fileBytes)
	// fmt.Println(fileContent)

	// video1 := video{
	// 	Id:          "101wqsa",
	// 	Title:       "Learning basics of Golang ",
	// 	Description: "logo of golang with a gopher",
	// 	Imageurl:    "q/5MtWmsPBQOUAAAAAElFTkSuQmCC",
	// 	Url:         "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
	// }

	// video2 := video{
	// 	Id:          "123qsa",
	// 	Title:       "implementing json ",
	// 	Description: "logo of golang with a gopher",
	// 	Imageurl:    "DVPsgAAAABJRU5ErkJggg==",
	// 	Url:         "https://www.youtube.com/watch?v=Lp2qcCrdBLA",
	// }

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
