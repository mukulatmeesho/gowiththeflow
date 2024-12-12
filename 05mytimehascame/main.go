package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("Welcome to time machine\n")

	presentTime := time.Now()

	fmt.Println("Current time:", presentTime)

	fmt.Println("Formatted date (MM-DD-YYYY):", presentTime.Format("01-02-2006"))
	fmt.Println("Formatted date with weekday (MM-DD-YYYY Weekday):", presentTime.Format("01-02-2006 Monday"))
	fmt.Println("Formatted date with time (MM-DD-YYYY HH:MM:SS Weekday):", presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2001, time.April, 12, 0, 0, 0, 0, time.UTC)

	fmt.Println("Created date (default format):", createdDate)

	fmt.Println("Formatted created date (MM-DD-YYYY Weekday):", createdDate.Format("01-02-2006 Monday"))
}
