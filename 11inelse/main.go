package main

import "fmt"

func main() {

	num := 5
	var result string
	if num > 5 {
		result = "Num is greater"
	} else if num < 5 {
		result = "Regular user"
	} else {
		result = "its 5"
	}

	fmt.Println(result)

}
