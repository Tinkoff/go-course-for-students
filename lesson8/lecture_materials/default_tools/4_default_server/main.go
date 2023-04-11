package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hello, World!"))
}
