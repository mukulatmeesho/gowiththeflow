package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.google.com"

func checkError(err error, errMsg string) {
	if err != nil {
		fmt.Printf(errMsg)
		panic(err)
	}
}

func main() {
	fmt.Println("web req in go")

	resp, err := http.Get(url)

	checkError(err, "Error reading file: ")
	fmt.Printf("Response is of type %T\n", resp)
	fmt.Printf("Response is of value %v\n", resp.Body)
	defer resp.Body.Close()

	dataBytes, err := io.ReadAll(resp.Body)
	checkError(err, "Error reading response body")

	content := string(dataBytes)
	fmt.Println(content)
}
