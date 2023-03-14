package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	fmt.Println("push to ch: ")
	ch <- "a"
	res := <-ch
	fmt.Println("get res: ", res)
}
