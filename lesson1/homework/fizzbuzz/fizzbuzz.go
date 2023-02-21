package fizzbuzz

import (
	"strconv"
	"strings"
)

func helpF(d int, s string, i int) string {
	if i%d == 0 {
		return s
	}
	return ""
}

func FizzBuzz(i int) string {
	result := strings.Builder{}
	result.WriteString(helpF(3, "Fizz", i))
	result.WriteString(helpF(5, "Buzz", i))
	if result.Len() == 0 {
		return strconv.Itoa(i)
	}
	return result.String()
}
