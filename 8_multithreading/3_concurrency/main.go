package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

func main() {
	// m := sync.Mutex{}

	var number uint32

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// m.Lock()
		// number++
		atomic.AddUint32(&number, 1)
		// m.Unlock()
		fmt.Printf("You are the %d visitor!\n", number)
	})

	http.ListenAndServe(":3000", nil)
}
