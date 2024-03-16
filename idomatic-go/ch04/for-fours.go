package main

import "fmt"

func main() {
	samples := []string{"hello", "crazy", "world"}

outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if string(r) == "l" {
				// skips inner loop processing and continues the outer loop
				continue outer
			}
		}
		fmt.Println()
	}
}
