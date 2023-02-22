package fizzbuzz

import (
	"strconv"
)

func FizzBuzz(i int) string {
	const (
		fizzBuzz = "FizzBuzz"
		buzz     = "Buzz"
		fizz     = "Fizz"
	)
	switch {
	case i%15 == 0:
		return fizzBuzz
	case i%5 == 0:
		return buzz
	case i%3 == 0:
		return fizz
	default:
		return strconv.Itoa(i)
	}
}
