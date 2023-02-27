package main

import (
	"fmt"
	"lecture01_homework/fizzbuzz"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 || i%5 == 0 {
			fmt.Println(fizzbuzz.FizzBuzz(i))
		} else {
			fmt.Println(i)
		}
	}
}
