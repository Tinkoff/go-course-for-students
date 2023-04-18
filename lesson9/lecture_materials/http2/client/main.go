package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "https://www.google.com", nil)
	if err != nil {
		log.Fatalf("can't create request: %s", err.Error())
	}

	fmt.Println(req.Proto)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("can't do request: %s", err.Error())
	}

	fmt.Println(res.Proto)
}
