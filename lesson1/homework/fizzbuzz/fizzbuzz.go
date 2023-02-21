package fizzbuzz

import "fmt"

func FizzBuzz(i int) string {
	if i%15 == 0 {
		return "FizzBuzz"
	} else if i%5 == 0 {
		return "Buzz"
	} else if i%3 == 0 {
		return "Fizz"
	} else {
		return fmt.Sprintf("%d", i)
	}
}
