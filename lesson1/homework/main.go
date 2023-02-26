package main

import (
	"fmt"
	"strconv"
	"lecture01_homework/fizzbuzz"
)

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzbuzz.FizzBuzz(i))
	}
}
