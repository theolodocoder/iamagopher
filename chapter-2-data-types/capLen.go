package main

import (
	"fmt"
)

func main() {
	// if only length is defined then capacity == len
	a := make([]int, 4)
	fmt.Println("L:", len(a), "C:", cap(a))

	b := []int{0, 1, 2, 3, 4}
	fmt.Println("L:", len(b), "C:", cap(b))

	// addind capacity
	aSlice := make([]int, 4, 4)
	fmt.Println(aSlice)

	// add an element to aSlice even though the len == cap
	aSlice = append(aSlice, 5)
	// the capacity is then doubled to fit the appended el
	fmt.Println("L:", len(aSlice), "C:", cap(aSlice))
}
