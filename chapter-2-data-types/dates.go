package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	if len(os.Args) != 2 {
		fmt.Println("Usage: dates parse_string")
		return
	}

	dateString := os.Args[1]

	// check if the string is date only
	date, err := time.Parse("02 January 2006", dateString)
	if err == nil {
		fmt.Println("Full:", date)
		fmt.Println("Time:", date.Day(), date.Month(), date.Year())
	}

	// check if its a date + time
	date, err = time.Parse("02 January 2005 15:04", dateString)
	if err == nil {
		fmt.Println("Full:", date)
		fmt.Println("Date:", date.Day(), date.Month(), date.Year())
		fmt.Println("Time:", date.Hour(), date.Minute())
	}

	// check if its date+time but with months represented with number?
	date, err = time.Parse("02-01-2005 15:05", dateString)
	if err == nil {
		fmt.Println("Full:", date)
		fmt.Println("Date:", date.Day(), date.Month(), date.Year())
		fmt.Println("Time:", date.Hour(), date.Minute())
	}

	// check if its time only
	date, err = time.Parse("15:04", dateString)
	if err == nil {
		fmt.Println("Full", date)
		fmt.Println("Time:", date.Hour(), date.Minute())
	}

	t := time.Now().Unix()
	fmt.Println("Epoch Time : ", t)

	// converting epoch time to time.Tieme
	date = time.Unix(t, 0)
	fmt.Println("Date:", date.Day(), date.Month(), date.Year())
	fmt.Printf("Time: %d:%d\n", date.Hour(), date.Minute())
	duration := time.Since(start)
	fmt.Println("Execution time: ", duration)
}
