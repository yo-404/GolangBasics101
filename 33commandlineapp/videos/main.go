package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// creating flag named "get" for videos
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	// inputs for get command
	getAll := getCmd.Bool("all", false, "Get all videos")
	getID := getCmd.String("id", "", "youtube video ID")

	// videos add command

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	// arguments for add command

	addID := addCmd.String("id", "", "youtube video id")
	addTitle := addCmd.String("title", "", "youtube video title")
	addUrl := addCmd.String("url", "", "youtube video URL")
	addImageUrl := addCmd.String("imageurl", "", "youtube image url")
	addDesc := addCmd.String("desc", "", "youtube video description")

	// validation
	if len(os.Args) < 2 {
		fmt.Println("Expected 'get' or 'add' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get": //if its the 'get' command
		HandleGet(getCmd, getAll, getID)
	case "add":
		HandleAdd(addCmd, addID, addTitle, addUrl, addImageUrl, addDesc)
	default:
	}

}

func HandleGet(getCmd *flag.FlagSet, all *bool, id *string) {

	getCmd.Parse(os.Args[2:])

	if *all == false && *id == "" {
		fmt.Print("id is required or specify --all for all videos")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	if *all {
		//return all videos
		videos := getVideos()

		fmt.Printf("ID \t Title \t URL \t ImageURL \t Description \n")
		for _, video := range videos {
			fmt.Printf("%v \t %v \t %v \t %v \t %v \n", video.Id, video.Title, video.Url, video.Imageurl, video.Description)
		}

		return
	}

	if *id != "" {
		videos := getVideos()
		id := *id
		for _, video := range videos {
			if id == video.Id {
				fmt.Printf("ID \t Title \t URL \t ImageURL \t Description \n")
				fmt.Printf("%v \t %v \t %v \t %v \t %v \n", video.Id, video.Title, video.Url, video.Imageurl, video.Description)
			}
		}
	}

}

func ValidateVideo(addCmd *flag.FlagSet, id *string, title *string, url *string, imageUrl *string, description *string) {

	addCmd.Parse(os.Args[2:])

	if *id == "" || *title == "" || *url == "" || *imageUrl == "" || *description == "" {
		fmt.Print("all fields are required for adding a video")
		addCmd.PrintDefaults()
		os.Exit(1)
	}

}

func HandleAdd(addCmd *flag.FlagSet, id *string, title *string, url *string, imageUrl *string, description *string) {

	ValidateVideo(addCmd, id, title, url, imageUrl, description)

	video := video{
		Id:          *id,
		Title:       *title,
		Description: *description,
		Imageurl:    *imageUrl,
		Url:         *url,
	}

	videos := getVideos()
	videos = append(videos, video)

	saveVideos(videos)

}
