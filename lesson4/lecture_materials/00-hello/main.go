package main

import (
	"fmt"
	"time"
)

func main() {
	go hello("Максим")
	go hello("Люся")
	go hello("Миша")
	go hello("Илья")
	go hello("Петя")

	time.Sleep(100 * time.Millisecond)
}

func hello(name string) {
	fmt.Printf("Привет %s!\n", name)
}
