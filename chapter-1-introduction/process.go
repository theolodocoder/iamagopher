package main

import (
	"os"
	"strconv"
)

func process() {
	var total, nInts, nFloats int
	invalid := make([]string, 0)
	arguments := os.Args

	for _, v := range arguments[1:] {

		_, err := strconv.Atoi(v)
		if err == nil {
			total++
			nInts++
			continue
		}

		_, err = strconv.ParseFloat(v, 64)
		if err == nil {
			total++
			nFloats++
			continue
		}

		invalid = append(invalid, v)
	}
}
