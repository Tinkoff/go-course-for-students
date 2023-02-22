package fizzbuzz

import "strconv"

func FizzBuzz(i int) (response string) {
	if i%15 == 0 {
		response = "FizzBuzz"
	} else if i%5 == 0 {
		response = "Buzz"
	} else if i%3 == 0 {
		response = "Fizz"
	} else {
		response = strconv.Itoa(i)
	}
	return
}
