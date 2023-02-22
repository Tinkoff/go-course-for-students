package fizzbuzz

import "strconv"

func FizzBuzz(i int) string {
	result := strconv.Itoa(i)

	if i%15 == 0 {
		result = "FizzBuzz"
	} else if i%5 == 0 {
		result = "Buzz"
	} else if i%3 == 0 {
		result = "Fizz"
	}

	return result
}
