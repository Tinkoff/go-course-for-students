package fizzbuzz

import "strconv"

func FizzBuzz(i int) string {
	var str = strconv.Itoa(i)
	if i%3 == 0 {
		str = "Fizz"
	}
	if i%5 == 0 {
		str = "Buzz"
	}
	if i%15 == 0 {
		str = "FizzBuzz"
	}
	return str
}
