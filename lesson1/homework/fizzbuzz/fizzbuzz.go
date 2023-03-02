package fizzbuzz

import (
	"strconv"
)

func FizzBuzz(i int) string {
	// TODO
	if i%15 == 0 {
		return "FizzBuzz"
	} else if i%5 == 0 {
		return "Buzz"
	} else if i%3 == 0 {
		return "Fizz"
	}
	return strconv.Itoa(i)
}
