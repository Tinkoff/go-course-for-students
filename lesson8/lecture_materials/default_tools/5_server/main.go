package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/users", UsersHandler)

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hello, World!"))
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("method: %s\n", r.Method)
	fmt.Printf("query values: %v\n", r.URL.Query())
	fmt.Printf("headers: %v\n", r.Header)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// not necessary
	// defer r.Body.Close()

	fmt.Printf("body: %s\n", string(body))

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Add("Server", "matrix")
	_, _ = w.Write([]byte("My name is..."))
}
