package main

import (
	"fmt"
	"os"
	"strconv"
)

func controls() {
	if len(os.Args) != 2 {
		fmt.Println("please provide a cmd argument")
		return
	}
	argument := os.Args[1]

	switch argument {
	case "0":
		fmt.Println("Zero")
	case "1":
		fmt.Println("One")
	case "2", "3", "4":
		fmt.Println("2 or 3 or 4")
		fallthrough
	default:
		fmt.Println("Value : ", argument)
	}

	value, err := strconv.Atoi(argument)

	if err != nil {
		fmt.Println("Cannot convert to int", argument)
		return
	}

	switch {
	case value == 0:
		fmt.Println("Zero")
	case value > 0:
		fmt.Println("Positive number")
	case value < 0:
		fmt.Println("Negative number")
	default:
		fmt.Println("This should never happen", value)
	}

}
