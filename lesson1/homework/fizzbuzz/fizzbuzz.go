package fizzbuzz

import (
	"strconv"
)

const FIZZ_BASE = 3
const BUZZ_BASE = 5

func FizzBuzz(i int) string {
	var result string
	if i%FIZZ_BASE == 0 {
		result += "Fizz"
	}
	if i%BUZZ_BASE == 0 {
		result += "Buzz"
	}
	if i%FIZZ_BASE != 0 && i%BUZZ_BASE != 0 {
		result = strconv.Itoa(i)
	}
	return result
}
