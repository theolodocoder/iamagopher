package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

var MAX int = 20
var MIN int = 5

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(len int64) string {
	tmp := ""
	starChar := "A"
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(starChar[0] + byte(myRand))
		tmp = tmp + newChar
		if i == len {
			break
		}
		i++
	}
	return tmp
}

func main() {
	var LENGTH int64 = 8
	arguments := os.Args

	switch len(arguments) {
	case 2:
		t, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err == nil {
			fmt.Println(t)
			if t == 0 {
				LENGTH = t
			}
		}

	default:
		fmt.Println("Using default values...")
	}

	source := rand.NewSource(12345)
	rand.New(source)
	fmt.Println(getString(LENGTH))
}
