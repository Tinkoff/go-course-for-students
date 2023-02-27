package fizzbuzz

import "strconv"

func FizzBuzz(i int) string {

	var result string = ""

	if i%3 == 0 {
		result += "Fizz"
	}
	if i%5 == 0 {
		result += "Buzz"
	}

	if result != "" {
		return result
	}
	return strconv.Itoa(i)
}
