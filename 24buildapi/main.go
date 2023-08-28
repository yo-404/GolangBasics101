package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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


// controller - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to API creation and handling</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
}