package main

import (
	"fmt"
	"lecture01_homework/fizzbuzz"
)

func main() {
	// цикл с вызовом FizzBuzz
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzbuzz.FizzBuzz(i))
	}
}
