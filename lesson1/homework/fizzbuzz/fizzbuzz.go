package fizzbuzz

import "strconv"

func FizzBuzz(i int) string {
	// TODO
	if i%3 == 0 {
		if i%5 == 0 {
			return "FizzBuzz"
		} else {
			return "Fizz"
		}
	} else if i%5 == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(i)
	}
}
