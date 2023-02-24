package fizzbuzz

import "strconv"

func FizzBuzz(i int) string {
	// TODO
	ans := ""
	if i%5 != 0 && i%3 != 0 {
		ans = strconv.Itoa(i)
	} else {
		if i%3 == 0 {
			ans += "Fizz"
		}
		if i%5 == 0 {
			ans += "Buzz"
		}
	}
	return ans
}
