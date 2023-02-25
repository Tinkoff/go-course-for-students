package fizzbuzz

import "fmt"

func FizzBuzz(i int) string {
	switch {
	case i%15 == 0:
		return "FizzBuzz"
	case i%5 == 0:
		return "Buzz"
	case i%3 == 0:
		return "Fizz"
	default:
		return fmt.Sprintf("%v", i)
	}
}
