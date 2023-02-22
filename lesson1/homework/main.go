package main

import (
	"fmt"
	"lecture01_homework/fizzbuzz"
)

func main() {
	for i := 1; i <= 100; i++ {
		if fizzbuzz.FizzBuzz(i) == "1" {
			fmt.Println(i)
		} else {
			fmt.Println(fizzbuzz.FizzBuzz(i))
		}
	}
	//fmt.Println(fizzbuzz.FizzBuzz(4))
}
