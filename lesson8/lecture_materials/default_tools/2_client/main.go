package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// ctx := context.WithTimeout(context.Background(), 5*time.Second)
	// NewRequestWithContext
	req, err := http.NewRequest(http.MethodPost, "http://google.com/robots.txt", nil)
	if err != nil {
		panic(err)
	}

	c := http.Client{Timeout: 10 * time.Second}

	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
