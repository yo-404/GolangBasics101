package main

import "fmt"

//  model for courses - file
type Course struct {
	CourseId    string  `json:"courseid`
	CourseName  string  `json:"coursname"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db

var courses []Course

// middleware
func (c *Course)IsEmpty() bool {
	return c.CourseId == "" && c.CourseName=""
}

func main() {
	fmt.Println("Welcome to api creations")
}
