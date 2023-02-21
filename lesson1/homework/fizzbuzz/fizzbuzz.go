package fizzbuzz

import "strconv"

func FizzBuzz(i int) string {
	if i%3 != 0 && i%5 != 0 {
		return strconv.Itoa(i)
	}

	var result string
	if i%3 == 0 {
		result += "Fizz"
	}
	if i%5 == 0 {
		result += "Buzz"
	}
	return result
}
