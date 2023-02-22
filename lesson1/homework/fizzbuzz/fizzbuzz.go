package fizzbuzz

import (
	"strconv"
	"strings"
)

func FizzBuzz(i int) string {

	var res strings.Builder

	const FIZZ_MULT = 3
	const BUZZ_MULT = 5
	if i%FIZZ_MULT == 0 {
		res.WriteString("Fizz")
	}
	if i%BUZZ_MULT == 0 {
		res.WriteString("Buzz")
	}
	if len(res.String()) == 0 {
		res.WriteString(strconv.Itoa(i))
	}
	return res.String()
}
