package main

import "fmt"

func main() {
	fmt.Print("Welcome to arrays\n")

	var languages [5]string
	languages[0] = "Java"
	languages[1] = "GOLang"
	languages[2] = "GOLang"
	languages[4] = "GOLang"
	fmt.Println("List ", languages)
	fmt.Println("List length ", len(languages))

	var fruits = [4]string{"Apple", "banana", "orange"}
	fmt.Println(fruits)

	fmt.Println("Lang List")
	for i := 0; i < len(languages); i++ {
		fmt.Println(languages[i])
	}

	exists := false
	for _, lang := range languages {
		if lang == "Python" {
			exists = true
			break
		}
	}
	fmt.Println("\nDoes Python exist in languages? ", exists)

	var copiedLanguages = languages
	copiedLanguages[0] = "C++"
	fmt.Println("\nOriginal Languages List: ", languages)
	fmt.Println("Copied Languages List: ", copiedLanguages)
}
