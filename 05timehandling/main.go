package main

import (
	"fmt"
	"time"
)

func main() {
	// after sleep and tick
	fmt.Println("Welcome to time handling in golang")
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006"))
	fmt.Println(presentTime.Format("Jan 02 2006"))
	fmt.Println(presentTime.Format("Jan 02 2006 IST 15:04:05"))
	fmt.Println(presentTime.Format("Monday Jan 02 2006 IST 15:04:05"))

	createDate := time.Date(2021, time.December, 10, 20, 18, 0, 0, time.UTC)

	fmt.Println("date for today is ", createDate)
	fmt.Println("formatted date", createDate.Format("Monday 02 Jan 2006"))

	fmt.Scanf("h")
	// os.Exit(0)
}
