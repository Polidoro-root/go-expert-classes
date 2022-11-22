package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func PostRequest() {
	c := http.Client{}

	jsonVar := bytes.NewBuffer([]byte(`{"name": "joao"}`))

	resp, err := c.Post("http://google.com", "application/json", jsonVar)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	io.CopyBuffer(os.Stdout, resp.Body, nil)
}

func main() {
	c := http.Client{}

	req, err := http.NewRequest("GET", "http://google.com", nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/json")

	resp, err := c.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	println(string(body))
}
