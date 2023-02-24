package fizzbuzz

import "strconv"

func FizzBuzz(i int) string {
	if i%15 == 0 {
		return "FizzBuzz"
	}
	if i%3 == 0 {
		return "Fizz"
	}
	if i%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(i)
}
