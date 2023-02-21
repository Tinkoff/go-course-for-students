package fizzbuzz

import "strconv"

const (
	Fizz = "Fizz"
	Buzz = "Buzz"
)

func FizzBuzz(i int) string {
	var result string
	if i%3 == 0 {
		result += Fizz
	}
	if i%5 == 0 {
		result += Buzz
	}
	if len(result) == 0 {
		result = strconv.Itoa(i)
	}
	return result
}
