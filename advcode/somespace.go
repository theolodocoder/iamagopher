package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func firstAndLast(s string) string {
	var firstNum, lastNum rune

	for _, c := range s {
		if unicode.IsDigit(c) {
			if firstNum == 0 {
				firstNum = c
			}
			lastNum = c
		}
	}

	firstNumStr := string(firstNum)
	lastNumStr := string(lastNum)

	totalStr := firstNumStr + lastNumStr
	return totalStr
}

func main() {
	var total int
	// open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer file.Close()

	// create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// read each file from the file
	for scanner.Scan() {
		line := scanner.Text()
		t := firstAndLast(line)
		v, err := strconv.Atoi(t)
		if err != nil {
			fmt.Println("Error :", err)
			return
		}
		total += v
	}

	fmt.Println(total)

}
