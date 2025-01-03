package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://www.youtube.com:3000/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=30"

func main() {
	fmt.Println("welcome to url handling in golang")
	fmt.Println(myUrl)

	parse, _ := url.Parse(myUrl)
	fmt.Println(parse)
	fmt.Println(parse.Scheme)
	fmt.Println(parse.Host)
	fmt.Println(parse.Path)
	fmt.Println(parse.Port())
	fmt.Println(parse.RawQuery)

	queryParams := parse.Query()
	fmt.Printf("the type of query params arr: %T\n", queryParams)
	fmt.Println(queryParams["list"])

	for _, val := range queryParams {
		//fmt.Println(key, val)
		fmt.Println("parsm val is ", val)
	}
	partsOfUrl := &url.URL{ //reference of url
		Scheme:  "https",
		Host:    "google.com",
		Path:    "/search",
		RawPath: "q=thisismemukul",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)

}
