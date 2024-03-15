package main

import "fmt"

// generic
func Print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()

}
func generics() {
	Ints := []int{1, 2, 3}
	Strings := []string{"one", "two", "3"}
	Print(Ints)
	Print(Strings)

}
