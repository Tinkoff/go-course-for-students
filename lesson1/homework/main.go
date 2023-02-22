package main

import (
	"fmt"
	"lecture01_homework/fizzbuzz"
)

func main() {
	// a fizzbuzz game for the range [1; 100]
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzbuzz.FizzBuzz(i))
	}
}
