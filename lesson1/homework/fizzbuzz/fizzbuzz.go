package fizzbuzz

import "strconv"

func helpF(d int, s string, i int) string {
	if i%d == 0 {
		return s
	}
	return ""
}

func FizzBuzz(i int) string {
	result := helpF(3, "Fizz", i) + helpF(5, "Buzz", i)
	if result == "" {
		return strconv.Itoa(i)
	}
	return result
}
