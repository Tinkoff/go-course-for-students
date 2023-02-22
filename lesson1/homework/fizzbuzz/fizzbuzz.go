package fizzbuzz

import (
	"strconv"
	"strings"
)

func FizzBuzz(i int) string {
	var sb strings.Builder
	if i%3 == 0 {
		sb.WriteString("Fizz")
	}
	if i%5 == 0 {
		sb.WriteString("Buzz")
	}
	if sb.Len() == 0 {
		sb.WriteString(strconv.Itoa(i))
	}
	return sb.String()
}
