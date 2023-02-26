package main

import (
	"fmt"
	"lecture01_homework/fizzbuzz"
)

func main() {
    for i := 1; i <= 100; i++ {
        fmt.Printf("%v ", fizzbuzz.FizzBuzz(i));
    }
}
