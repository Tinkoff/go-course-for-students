package fizzbuzz

import (
	"strconv"
)

func FizzBuzz(i int) string {
	switch 0 {
	case i % 15:
		return "FizzBuzz"
	case i % 3:
		return "Fizz"
	case i % 5:
		return "Buzz"
	default:
		return strconv.Itoa(i)
	}
}
