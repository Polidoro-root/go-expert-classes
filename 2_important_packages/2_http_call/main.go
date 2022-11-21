package main

import (
	"fmt"
	"io"
	"net/http"
)

func httpCall() {
	req, err := http.Get("https://google.com")

	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	println(string(res))
}

func main() {
	defer fmt.Println("First line")
	fmt.Println("Second line")
	fmt.Println("Third line")

}
