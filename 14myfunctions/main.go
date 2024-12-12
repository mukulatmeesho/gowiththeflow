package main

import "fmt"

func greetings() {
	fmt.Println("nameste")
	func() {
		fmt.Println("Hello from the anonymous function!")
	}()

}

func printName(prompt string) {
	fmt.Printf("Hello %v, Welcome to Meesho\n", prompt)
	func(name string) {
		fmt.Printf("Hello %s, this is an anonymous function!\n", name)
	}("Mukul")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}

	return total, "Done"
}
func main() {
	fmt.Print("Functions in go\n")

	greetings()
	printName("mukul")
	result := adder(3, 6)
	fmt.Println("Value is ", adder(5, 7))
	fmt.Println("Value is ", result)

	resultPro, msg := proAdder(7, 8, 7, 8, 6, 7, 3)
	fmt.Println(msg)
	fmt.Println(resultPro)
}
