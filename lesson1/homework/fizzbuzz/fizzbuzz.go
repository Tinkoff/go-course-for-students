package fizzbuzz

import "strconv"

func FizzBuzz(i int) string {
	// TODO
	ans := ""
	if i%5 != 0 && i%3 != 0 {
		return strconv.Itoa(i)
	}
	if i%3 == 0 {
		ans += "Fizz"
	}
	if i%5 == 0 {
		ans += "Buzz"
	}
	return ans
}
