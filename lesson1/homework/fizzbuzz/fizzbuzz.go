package fizzbuzz

import "strconv"

func FizzBuzz(i int) string {
	var res string
	switch {
	case i%15 == 0:
		res = "FizzBuzz"
	case i%3 == 0:
		res = "Fizz"
	case i%5 == 0:
		res = "Buzz"
	default:
		res = strconv.Itoa(i)
	}
	return res
}
