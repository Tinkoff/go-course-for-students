package main

import (
	"fmt"
	"lecture01_homework/fizzbuzz"
)

const LEFT = 1
const RIGHT = 100

func main() {
	for i := LEFT; i <= RIGHT; i++ {
		fmt.Println(fizzbuzz.FizzBuzz(i))
	}
}
