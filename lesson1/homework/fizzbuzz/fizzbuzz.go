package fizzbuzz

import (
	"strconv"
	"strings"
)

// Returns resultIfDividable if number is dividable by divider
// Empty string otherwise
func helpF(divider int, resultIfDividable string, number int) string {
	if number%divider == 0 {
		return resultIfDividable
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
