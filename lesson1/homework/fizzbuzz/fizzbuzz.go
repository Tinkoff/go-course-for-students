package fizzbuzz

import (
    "strconv"
    "fmt"
)

func FizzBuzz(i int) string {
	if i % 15 == 0 {
        return "FizzBuzz"
    }
    if i % 5 == 0 {
        return "Buzz"
    }
    if i % 3 == 0 {
        return "Fizz"
    }
	return strconv.Itoa(i)
}

func main() {
    for i := 1; i <= 100; i++ {
        fmt.Printf("%v ", FizzBuzz(i));
    }
}
