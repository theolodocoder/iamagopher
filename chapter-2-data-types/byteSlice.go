package main

import (
	"fmt"
)

func main() {
	b := make([]byte, 12)
	fmt.Println(b)
	fmt.Println(len(b), cap(b))
	b = []byte("This is a slice %")
	fmt.Println(b)
	fmt.Println(len(b), cap(b))
}
