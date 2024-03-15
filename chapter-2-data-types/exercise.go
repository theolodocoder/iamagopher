package main

import (
	"fmt"
)

func arrToSlice(arr1 [3]int, arr2 [4]int) []int {
	copiedSlice := make([]int, len(arr1)+len(arr2))
	copy(copiedSlice, arr1[:])
	copy(copiedSlice[len(arr1):], arr2[:])
	return copiedSlice
}

func main() {
	arr1 := [3]int{1, 2, 3}
	arr2 := [4]int{4, 5, 6, 7}
	sl := arrToSlice(arr1, arr2)
	fmt.Println(sl)

	// ================
	x := [3]int{1, 2, 3}
	// y :=make([]int, 2)
	// num represents the number of el copied
	// num := copy(y, x[1:])
	copy(x[1:], x[:])
	fmt.Println(x)

	str := "😊"
	fmt.Println(len(str))
}
