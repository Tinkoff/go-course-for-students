package fizzbuzz

import "strconv"

func FizzBuzz(i int) string {
	var val string
	if i%3 == 0 {
		val += "Fizz"
	}
	if i%5 == 0 {
		val += "Buzz"
	}
	if val == "" {
		return strconv.Itoa(i)
	}
	return val
}
