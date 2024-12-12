package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(fmt.Sprintf("Error while reading input: %v", err)) // Panic if an error occurs
	}
	return strings.TrimSpace(input)
}

func getValidRating(prompt string) float64 {
	ratingInput := readInput(prompt)

	rating, err := strconv.ParseFloat(ratingInput, 64)
	checkErr(err, "invalid rating, please enter a number between 1 and 5")

	if rating < 1 || rating > 5 {
		panic("rating out of bounds, please rate between 1 and 5")
	}

	return rating
}

func checkErr(err error, errMsg string) {
	if err != nil {
		panic(errMsg)
	}
}
func main() {
	fmt.Println("Welcome to my pizza app!")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error encountered:", r)
			fmt.Println("The program will exit now.")
		}
	}()

	nameInput := readInput("Please enter your name: ")

	rating := getValidRating("Please rate our pizza (1 to 5): ")

	fmt.Printf("Thank you, %v! You rated our pizza a %.1f.\n", nameInput, rating)
	newRating := rating + 1
	if newRating > 5 {
		newRating = 5
	}
	fmt.Printf("Added 1 to your rating. Your updated rating is %.1f.\n", newRating)
}
