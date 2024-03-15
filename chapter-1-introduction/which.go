package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("No arguments provided")
		return
	}

	file := arguments[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	for _, directory := range pathSplit {

		fullPath := filepath.Join(directory, file)
		// does it exist?
		fileInfo, err := os.Stat(fullPath)
    fmt.Println(fileInfo)
		// fmt.Println(os.Stat("C:\\Users\\USER\\AppData\\Local\\Programs\\Python\\Python312\\python")) break
		if err == nil {
			fmt.Println("noerr", fileInfo)
			mode := fileInfo.Mode()
			// is this a regnoular file
			if mode.IsRegular() {
				// is it executable
				if mode&0111 != 0 {
					fmt.Println(fullPath)
					return
				}
			}
		}
	}
}
