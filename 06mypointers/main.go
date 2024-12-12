package main

import "fmt"

func main() {

	fmt.Print("Welcome to pointers class\n")
	//var ptr *int
	//fmt.Println(ptr)
	myNumber := 23
	var numPtr = &myNumber
	fmt.Println(numPtr)
	fmt.Println(*numPtr)
}
