package main

import "fmt"

func main() {
	var a = 1
	age := 9
	var name string = "Mukul"
	fmt.Println(a)
	fmt.Println(age)
	fmt.Println(name)
	fmt.Printf("My name is %v, I am %v years old\n", name, age)

	// 1. Boolean Type
	var isActive = true
	fmt.Printf("Value %v\nType %T\n", isActive, isActive)

	// 2. Numeric Types
	var num int = 10
	var num8 int8 = 100
	var num16 int16 = 10000
	var num32 int32 = 1000000
	var num64 int64 = 1234567890
	fmt.Println(num, num8, num16, num32, num64)

}
