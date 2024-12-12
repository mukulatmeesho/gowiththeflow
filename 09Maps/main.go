package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Let's GO with Maps\n")
	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["J"] = "Java"
	languages["PY"] = "Python"
	languages["GO"] = "GOLang"

	fmt.Println(len(languages))

	for key, val := range languages {
		fmt.Println(key, val)
	}

	value, ok := languages["GO"]
	if ok {
		fmt.Println("Found:", value)
	} else {
		fmt.Println("Key not found!")
	}

	delete(languages, "PY")
	fmt.Println(languages)

	multilanguage := map[string]string{
		"JS": "JavaScript",
		"GO": "GoLang",
		"PY": "Python",
	}
	fmt.Println(len(multilanguage))

	// Count occurrences of words
	text := "apple banana apple orange banana apple"
	wordCount := make(map[string]int)
	words := strings.Split(text, " ")
	for _, word := range words {
		wordCount[word]++
	}
	fmt.Println("\nWord frequency:", wordCount)

}
