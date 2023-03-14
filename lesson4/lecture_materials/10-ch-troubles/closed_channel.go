package main

import "fmt"

func main() {
	ch := make(chan string, 5)
	close(ch)

	ch <- "ok"

	fmt.Printf("ok is: %v\n", <-ch)
}
