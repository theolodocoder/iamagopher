package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Need an Integer")
		return
	}

	index := args[1]
	i, err := strconv.Atoi(index)
	// return if there is an error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Using Index...", i)

	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Original Slice", aSlice)

	// check index is greater than the array len
	if i > len(aSlice)-1 {
		fmt.Println("Cannot delete ", i)
		return
	}
	aSlice = append(aSlice[:i], aSlice[i+1:]...)
	fmt.Println("Slice after first del technique", aSlice)

	// check index is greater than the array len
	if i > len(aSlice)-1 {
		fmt.Println("Cannot delete ", i)
		return
	}

	// del technique 2
	// set the el @ index to the last element
	aSlice[i] = aSlice[len(aSlice)-1]
	// remove last el
	aSlice = aSlice[:len(aSlice)-1]
	fmt.Println("Slice after second del technique", aSlice)
}
