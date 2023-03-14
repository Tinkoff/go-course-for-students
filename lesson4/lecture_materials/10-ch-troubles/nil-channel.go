package main

import "fmt"

func main() {
	var ch chan string

	select {
	case <-ch:
		fmt.Println("got msg from channel")
	}
}
