package fizzbuzz

import (
	"strconv"
)

func FizzBuzz(i int) string {
	mod3 := i % 3
	mod5 := i % 5
	if mod3 == 0 && mod5 == 0 {
		return "FizzBuzz"
	} else if mod3 == 0 {
		return "Fizz"
	} else if mod5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(i)
}
