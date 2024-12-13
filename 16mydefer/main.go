package main

import "fmt"

func main() {
	defer fmt.Println("one")
	fmt.Println("Defer in Go lang")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	defer fmt.Println()
	var name string = "Mukul"
	for i := 0; i < len(name); i++ {
		defer fmt.Printf(string(name[i]))
	}
	fmt.Println()
}
