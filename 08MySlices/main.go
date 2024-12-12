package main

import (
	"fmt"
	"sort"
)

func remove(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	fmt.Print("Slices ...\n")

	var fruitList = []string{"Apple", "Banana", "Orange"}

	fmt.Println(fruitList)

	fruitList = append(fruitList, "Grapes", "Lemon")

	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 123
	highScores[1] = 323
	highScores[2] = 923
	highScores[3] = 523

	highScores = append(highScores, 555, 43, 223)
	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)

	//remove a value from slices based on index
	highScores = remove(highScores, 2)
	fmt.Println(highScores)

}
