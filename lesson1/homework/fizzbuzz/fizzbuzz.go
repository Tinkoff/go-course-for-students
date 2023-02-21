package fizzbuzz

import "fmt"

func FizzBuzz(i int) (res string) {
	var check bool
	if i%3 == 0 {
		res += "Fizz"
		check = true
	}
	if i%5 == 0 {
		res += "Buzz"
		check = true
	}
	if !check {
		res = fmt.Sprintf("%d", i)
	}
	return
}
