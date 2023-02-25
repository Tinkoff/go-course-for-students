package main

import (
	"fmt"
	"lecture01_homework/fizzbuzz"
)

func main() {
	const (
		iterateFrom int = 1
		iterateTo   int = 100
	)

	for i := iterateFrom; i <= iterateTo; i++ {
		fmt.Println(fizzbuzz.FizzBuzz(i))
	}
}
