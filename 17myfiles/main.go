package main

import (
	"fmt"
	"io"
	"os"
)

func checkError(err error, errMsg string) {
	if err != nil {
		fmt.Printf(errMsg)
		panic(err)
	}
}

func main() {
	fmt.Println("GO with files")

	content := "I am learning go"

	file, err := os.Create("./content.txt")
	checkError(err, "")

	length, err := io.WriteString(file, content)
	checkError(err, "")

	fmt.Printf("Length is %v\n", length)

	readFile("./content.txt")

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
}

func readFile(filePath string) {
	file, err := os.Open(filePath)
	checkError(err, "Error opening file: ")
	defer file.Close()

	bytes := make([]byte, 1024)
	n, err := file.Read(bytes)
	checkError(err, "Error reading file: ")
	fmt.Printf("Read %d bytes: %s\n", n, string(bytes[:n]))

}
