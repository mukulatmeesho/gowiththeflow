package main

import (
	"fmt"
)

func main() {
	fmt.Print("Looping Go \n")
	var languages = [3]string{"java", "go", "javascript"}
	fmt.Println(len(languages))

	for d := 0; d < len(languages); d++ {
		fmt.Println(languages[d])
	}

	fmt.Println("For loop using range")

	for i := range languages {
		fmt.Println(languages[i])
	}

	fmt.Println("Using indexing")

	for index, language := range languages {
		fmt.Printf("Index of %v has value %v\n", index, language)
	}
	fmt.Println("While but for :) ")

	rogueValue := 1
	for rogueValue < 10 {
		//if rogueValue == 6 {
		//	break
		//}
		if rogueValue == 2 {
			goto lfo
		}
		if rogueValue == 6 {
			rogueValue++
			continue
		}
		fmt.Println(rogueValue)
		rogueValue++
	}

lfo:
	fmt.Println("Jumping to another method")
}
