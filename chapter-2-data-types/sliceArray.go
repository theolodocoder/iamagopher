package main

import (
	"fmt"
)

func change(s []string) {
	s[0] = "change_func"
}

func main() {
	a := [4]string{"Zero", "One", "Two", "Three"}
	fmt.Println("a: ", a)

	var S0 = a[0:1]
	fmt.Println(S0)
	S0[0] = "S0"

	var S12 = a[1:3]
	S12[0] = "S12_0"
	S12[1] = "S12_1"
	fmt.Println("a: ", a)

	change(S12)
	fmt.Println("a: ", a)

	fmt.Println("Capacity of S0:", cap(S0), "Length of S0:", len(S0))
	// Adding 4 elements to S0
	S0 = append(S0, "N1")
	S0 = append(S0, "N2")
	S0 = append(S0, "N3")

	a[0] = "-N1"
	fmt.Println("a: ", a)
}
