package fizzbuzz

import "strconv"

func FizzBuzz(i int) string {
	var answer string = ""
	if i%3 == 0 {
		answer += "Fizz"
	}
	if i%5 == 0 {
		answer += "Buzz"
	}
	if answer == "" {
		return strconv.Itoa(i)
	}
	return answer
}
