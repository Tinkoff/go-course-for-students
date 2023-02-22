package fizzbuzz

import "strconv"

// FizzBuzz returns Fizz if i is divided by 3, Buzz if i is divided by 5,
// FizzBuzz if i is divided by 15, and just i otherwise
//
// Note the number is divided by 15 iff it is divided by 3 and 5
func FizzBuzz(i int) string {
	var text string
	if i%3 == 0 {
		text += "Fizz"
	}
	if i%5 == 0 {
		text += "Buzz"
	}
	if len(text) == 0 {
		text = strconv.Itoa(i)
	}
	return text
}
