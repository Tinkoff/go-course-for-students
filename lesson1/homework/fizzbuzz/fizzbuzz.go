package fizzbuzz

import "strconv"

const FIZZ = "Fizz"

const BUZZ = "Buzz"

func FizzBuzz(i int) string {
	if i%15 == 0 {
		return FIZZ + BUZZ
	}
	if i%5 == 0 {
		return BUZZ
	}
	if i%3 == 0 {
		return FIZZ
	}
	return strconv.Itoa(i)
}
