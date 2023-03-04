package fizzbuzz

import "strconv"

func FizzBuzz(i int) string {
	ans := ""
	if i%3 == 0 {
		ans += "Fizz"
	}
	if i%5 == 0 {
		ans += "Buzz"
	}
	if i%3 != 0 && i%5 != 0 {
		ans += strconv.Itoa(i)
	}
	return ans
}
