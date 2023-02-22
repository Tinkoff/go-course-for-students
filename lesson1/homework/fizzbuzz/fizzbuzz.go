package fizzbuzz

import "strconv"

func FizzBuzz(i int) string {
	switch remainder := i % 15; {
	case remainder == 0:
		return "FizzBuzz"
	case remainder%5 == 0:
		return "Buzz"
	case remainder%3 == 0:
		return "Fizz"
	default:
		return strconv.Itoa(i)
	}
}
