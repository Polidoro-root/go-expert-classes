package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	http.ListenAndServe(":8080", mux)

}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Println("Start request")

	defer log.Println("Request done")

	select {
	case <-time.After(time.Second * 5):
		log.Println("Request processed with success")

		w.Write([]byte("Request processed with success"))
	case <-ctx.Done():
		log.Println("Request cancelled by the client")
	}
}
