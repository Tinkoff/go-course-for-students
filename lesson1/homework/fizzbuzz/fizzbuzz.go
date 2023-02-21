package fizzbuzz

import "fmt"

func FizzBuzz(i int) string {
	s := ""
	for j := 1; j <= i; j++ {
		if j%15 == 0 {
			s += "FizzBuzz"
		} else if j%3 == 0 {
			s += "Fizz"
		} else if j%5 == 0 {
			s += "Buzz"
		} else {
			s = fmt.Sprint(s, j)
		}
		s += " "
	}
	return s
}
