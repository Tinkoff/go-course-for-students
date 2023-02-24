package fizzbuzz

import "strconv"

func FizzBuzz(i int) string {
	var ret string
	switch {
	case (i%3 != 0) && (i%5 != 0):
		ret = strconv.Itoa(i)
	case i%15 == 0:
		ret = "FizzBuzz"
	case i%3 == 0:
		ret = "Fizz"
	case i%5 == 0:
		ret = "Buzz"
	}
	return ret
}
