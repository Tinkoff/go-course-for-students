package main

import (
	"fmt"
	"lecture01_homework/fizzbuzz"
)

func main() {
	// TODO тут напишите цикл с вызовом FizzBuzz
	for i := 1; i < 101; i++ {
		fmt.Println(fizzbuzz.FizzBuzz(i))
	}
	// fmt.Println(fizzbuzz.FizzBuzz(10))
	// fmt.Println("Fizz")
}
