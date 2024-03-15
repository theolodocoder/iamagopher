package main

import (
	"fmt"
	"math"
)

var Global int = 1023
var AnotherGlobal = -1334

func variables() {
	var j int
	i := AnotherGlobal + Global

	fmt.Println("Initial j value : ", j)
	j = Global

	k := math.Abs(float64(AnotherGlobal))
	fmt.Printf("Global=%d, i=%d, j=%d, k=%.2f\n", Global, i, j, k)
}
