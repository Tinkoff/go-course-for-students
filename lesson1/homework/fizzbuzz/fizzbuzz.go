package fizzbuzz

import "fmt"

func FizzBuzz(i int) string {
	result := fmt.Sprintf("%d", i)
	if i%3 == 0 {
		result = "Fizz"
	}
	if i%5 == 0 {
		result = "Buzz"
	}
	if i%15 == 0 {
		result = "FizzBuzz"
	}
	return result
}
