package fizzbuzz

import "fmt"

func FizzBuzz(i int) (res string) {
	if i%3 == 0 {
		res += "Fizz"
	}
	if i%5 == 0 {
		res += "Buzz"
	}
	if len(res) == 0 {
		res = fmt.Sprintf("%d", i)
	}
	return
}
