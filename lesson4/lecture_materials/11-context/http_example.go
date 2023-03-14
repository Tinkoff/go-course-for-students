package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	req, err := http.NewRequest("GET", "https://ya.ru", nil)
	if err != nil {
		log.Fatalf("%v", err)
	}

	ctx, cancel := context.WithTimeout(req.Context(), 30*time.Millisecond)
	defer cancel()

	req = req.WithContext(ctx)

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Printf("%v\n", res.StatusCode)
}
