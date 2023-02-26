package main

import (
	"fmt"
	"lecture01_homework/fizzbuzz"
)

func main() {
	for value := 1; value <= 100; value++ {
		fmt.Println(fizzbuzz.FizzBuzz(value))
	}
}
