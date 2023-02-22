package fizzbuzz

import "fmt"

// FizzBuzz
// Function to implement fizzbuzz game
// It accepts an integer and returns the result in accordance with the rules of the game
// Read more: https://ru.wikipedia.org/wiki/Fizz_buzz
func FizzBuzz(n int) string {
	if n%15 == 0 {
		return "FizzBuzz"
	}
	if n%3 == 0 {
		return "Fizz"
	}
	if n%5 == 0 {
		return "Buzz"
	}
	return fmt.Sprint(n)
}
