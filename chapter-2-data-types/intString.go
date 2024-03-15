package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	args := os.Args[1]
	n, err := strconv.Atoi(args)
	if err == nil {
		input := strconv.Itoa(n)
		fmt.Printf("Itoa() %s of type %T\n", input, input)
		input = strconv.FormatInt(int64(n), 10)
		fmt.Printf("FormatInt() %s of type %T\n", input, input)
		input = string(n)
		fmt.Printf("string() %s of type %T\n", input, input)
	} else {
		fmt.Println(err)
	}

}
