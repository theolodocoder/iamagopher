package main

import "fmt"

func loops() {
	// normal for loop
	for i := 0; i < 10; i++ {
		fmt.Print(i*i, " ")
	}
	fmt.Println()

	// idiomatic go :
	i := 0
	for ok := true; ok; ok = (i != 10) {
		fmt.Print(i*i, " ")
		i++
	}
	fmt.Println()

	// while loop simulation using for loops
	j := 0
	for {
		if j == 10 {
			break
		}
		fmt.Print(j*j, " ")
		j++
	}
	fmt.Println()

	aSlice := []int{-1, 2, 1, -1}
	for i, v := range aSlice {
		fmt.Println("index: ", i, "value: ", v)
	}
}
