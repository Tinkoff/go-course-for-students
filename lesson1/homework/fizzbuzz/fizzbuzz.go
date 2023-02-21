package fizzbuzz

import "fmt"

func FizzBuzz(i int) string {
	if i % 15 == 0 {
		return "FizzBuzz"
	} else if i % 3 == 0 {
		return "Fizz"
	} else if i % 5 == 0 {
		return "Buzz"
	}
	return fmt.Sprint(i)
}
