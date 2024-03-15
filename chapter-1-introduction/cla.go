package main

import (
	"fmt"
	"os"
	"strconv"
)

func cla() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("One argument expected")
		return
	}

	var min, max float64
	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			continue
		}
		if i == 1 {
			min = n
			max = n
			fmt.Println(n)
			continue
		}

		if n < min {
			min = n
			fmt.Println(n)
		}

		if n > max {
			max = n
			fmt.Println(n)
		}

	}
	fmt.Println("Min", min)
	fmt.Println("Max", max)

}
