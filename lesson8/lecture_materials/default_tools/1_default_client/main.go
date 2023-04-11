package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// http://google.com/robots.txt
	resp, err := http.Get("http://www.testingmcafeesites.com/testcat_ac.html")
	if err != nil {
		panic(err)
	}

	// https://andrii-kushch.medium.com/is-it-necessary-to-close-the-body-in-the-http-response-object-in-golang-171c44c9394d
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
