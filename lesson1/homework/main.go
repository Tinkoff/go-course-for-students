package main

import (
	"bufio"
	"fmt"
	"lecture01_homework/fizzbuzz"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for i := 0; i <= 100; i++ {
		fmt.Fprintln(w, fizzbuzz.FizzBuzz(i))
	}
}
