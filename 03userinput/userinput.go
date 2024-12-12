package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	welcome := "Welcome to user input machine"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")

	// comma ok || error ok

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}
	//input = input[:len(input)-1]
	input = strings.TrimSpace(input) // Remove any leading/trailing whitespace or newline

	if input == "" {
		fmt.Println("You did not enter a name. Please try again.")
		return
	}
	fmt.Printf("Welcome %v! Let's go :)\n", input)
}
