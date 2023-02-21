package fizzbuzz

import (
	"strconv"
)

func FizzBuzz(i int) string {
	var ans string

	if i%3 == 0 {
		ans += "Fizz"
	}
	if i%5 == 0 {
		ans += "Buzz"
	}

	if len(ans) == 0 {
		ans = strconv.Itoa(i)
	}

	return ans
}
