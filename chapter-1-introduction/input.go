package main

import "fmt"

func input() {
	fmt.Printf("Input your name ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Your name is", name)
}
