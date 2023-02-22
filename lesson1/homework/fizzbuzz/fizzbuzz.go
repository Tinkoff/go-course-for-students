package fizzbuzz

import "strconv"

func FizzBuzz(i int) string {
	multipleOfThree := i%3 == 0
	multipleOfFive := i%5 == 0

	if multipleOfThree && multipleOfFive {
		return "FizzBuzz"
	} else if multipleOfThree {
		return "Fizz"
	} else if multipleOfFive {
		return "Buzz"
	}

	return strconv.Itoa(i)
}
