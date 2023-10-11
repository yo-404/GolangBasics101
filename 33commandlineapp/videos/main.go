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

}
