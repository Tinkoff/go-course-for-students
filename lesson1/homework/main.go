package main

import (
	"fmt"
	"lecture01_homework/fizzbuzz"
)

func main() {
	// TODO тут напишите цикл с вызовом FizzBuzz
	// fmt.Println(fizzbuzz.FizzBuzz(10))
	for i := 1; i < 101; i++ {
		fmt.Println(fizzbuzz.FizzBuzz(i))
	}
}
