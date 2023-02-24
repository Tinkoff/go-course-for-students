package fizzbuzz

import "strconv"

func FizzBuzz(i int) string {
	switch {
	case i%15 == 0:
		return "FizzBuzz"
	case i%3 == 0:
		return "Fizz"
	case i%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(i)
	}
}
