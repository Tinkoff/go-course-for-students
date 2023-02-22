package fizzbuzz

import "strconv"

func FizzBuzz(i int) string {
	switch {
	case i%3 == 0 && i%15 != 0:
		return "Fizz"
	case i%5 == 0 && i%15 != 0:
		return "Buzz"
	case i%15 == 0:
		return "FizzBuzz"
	default:
		return strconv.Itoa(i)
	}
}
