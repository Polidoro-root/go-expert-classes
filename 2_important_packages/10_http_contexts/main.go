package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	ctx := context.Background()

	// ctx, cancel := context.WithTimeout(ctx, time.Second)
	// ctx, cancel := context.WithCancel(ctx)

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second))

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)

	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	println(string(body))
}
