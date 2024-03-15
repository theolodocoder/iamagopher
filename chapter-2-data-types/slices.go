package main

import "fmt"

func main() {
	aSlice := make([]float64, 3)
	fmt.Println(aSlice)
	// two dimensional
	bSlice := make([][]float64, 2)
	fmt.Println(bSlice)

	fmt.Println(bSlice, len(bSlice), cap(bSlice))
	aSlice = append(aSlice, 22.0)

	bSlice = append(bSlice, aSlice)
	fmt.Println(bSlice)
	cSlice := make([]complex64, 3)
	cSlice = append(cSlice, 44i)
	fmt.Println(cSlice)

	// visiting all elements of a 2d slice
	for _, i := range bSlice {
		for _, k := range i {
			fmt.Print(k, " ")
		}
		fmt.Println()
	}
}
