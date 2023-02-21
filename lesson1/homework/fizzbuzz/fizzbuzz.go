package fizzbuzz

import "strconv"

func FizzBuzz(i int) string {
	// TODO
	var ret string
	switch {
	case i%15 == 0:
		ret = "FizzBuzz"
	case i%5 == 0:
		ret = "Buzz"
	case i%3 == 0:
		ret = "Fizz"
	default:
		ret = strconv.Itoa(i)
	}
	return ret
}
