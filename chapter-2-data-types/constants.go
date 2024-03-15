package main

import "fmt"

type Digit int
type Power2 int

const PI = 3.1415926

const (
	C1 = "C1C1C1"
	C2 = "C2C2C2"
	C3 = "C3C3C3"
)

func main() {
	const s1 = 123
	var v1 float32 = s1 * 12
	fmt.Println(v1)
	fmt.Println(PI)

	/*
			The constant generator iota is used for declaring a sequence of related values that
		use incrementing numbers without the need to explicitly type each one of them.
	*/
	const (
		Zero Digit = iota
		One
		Two
		Three
		Four
	)

	fmt.Println(One)
	fmt.Println(Two)

	const (
		pg_0 Power2 = 1 << iota
		_
		pg_2
		_
		pg_4
		_
		pg_6
	)

	fmt.Println(pg_0)
	fmt.Printf("%b\n", pg_2)
	fmt.Printf("%b\n", pg_4)
	fmt.Printf("%b\n", pg_6)

}
