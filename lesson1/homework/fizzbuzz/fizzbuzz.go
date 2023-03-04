package fizzbuzz

import "strconv"

func FizzBuzz(i int) string {
	var result string
	switch {
	case i%15 == 0:
		result = "FizzBuzz"
	case i%5 == 0:
		result = "Buzz"
	case i%3 == 0:
		result = "Fizz"
	default:
		result = strconv.Itoa(i)
	}
	return result
}
