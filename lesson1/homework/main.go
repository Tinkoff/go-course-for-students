package main

import (
	"fmt"
	"lecture01_homework/fizzbuzz"
)

func main() {
	// TODO тут напишите цикл с вызовом FizzBuzz
	for i := 0; i <= 100; i++ {
		fmt.Println(fizzbuzz.FizzBuzz(i))
	}
}
