package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Print("Switch case in go \n")

	diceNumber := rand.Intn(6) + 1

	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("you can open")
	case 2:
		fmt.Println("2")
	case 6:
		fmt.Println("6")
	default:
		fmt.Println("What was that")
	}
}
