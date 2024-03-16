package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// ##this concept is called shadowing
	// x := 10
	// if x > 5 {
	// 	fmt.Println(x)
	// 	x := 5
	// 	fmt.Println(x)
	// }
	// fmt.Println(x)
	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("low")
	} else if n > 5 {
		fmt.Println("big", n)
	} else {
		fmt.Println("good", n)
	}
}
