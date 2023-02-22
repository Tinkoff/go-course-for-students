package fizzbuzz

import (
	"strconv"
)

const FizzBase = 3
const BuzzBase = 5

func FizzBuzz(i int) string {
	var result string
	if i%FizzBase == 0 {
		result += "Fizz"
	}
	if i%BuzzBase == 0 {
		result += "Buzz"
	}
	if i%FizzBase != 0 && i%BuzzBase != 0 {
		result = strconv.Itoa(i)
	}
	return result
}
