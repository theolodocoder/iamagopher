package main

import "fmt"

func main() {
	b := []byte("some string")
	fmt.Printf("%s", b)

	// runes
	r := '$'
	fmt.Println()
	fmt.Println(r)
	fmt.Printf("%c", r)
	fmt.Printf("As a string: %s and as a character: %c\n", r, r)
	some := "string + $"
	for _, v := range some {
		fmt.Printf("%c ", v)
	}

	fmt.Println()
}
