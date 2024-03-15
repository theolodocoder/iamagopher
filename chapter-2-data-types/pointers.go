package main

import "fmt"

type aStruct struct {
	filed1 complex128
	filed2 int
}

func processPointer(x *int) {
	*x = *x * *x
}

func returnPointer(x float64) *float64 {
	temp := 2 * x
	return &temp
}

func bothPointers(x *float64) *float64 {
	temp := 2 * *x
	return &temp
}

func main() {
	var f float64 = 12.123
	fmt.Println("Mem of f :", &f)

	// pointer to f
	fp := &f
	fmt.Println("Memory address of f:", fp)

	// destructure the pointer to get the value
	fmt.Println("Value of f:", *fp)

	var i int = 10
	ip := &i
	processPointer(ip)

	var k *aStruct
	fmt.Println(k)

}
