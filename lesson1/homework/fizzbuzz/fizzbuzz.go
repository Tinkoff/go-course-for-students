package fizzbuzz

import "strconv"

func FizzBuzz(i int) string {
	var str string
	if i%15 == 0 {
		str = "FizzBuzz"
	} else if i%3 == 0 {
		str = "Fizz"
	} else if i%5 == 0 {
		str = "Buzz"
	} else {
		str = strconv.Itoa(i)
	}
	return str
}
