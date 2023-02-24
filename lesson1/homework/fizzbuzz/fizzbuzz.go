package fizzbuzz

import "strconv"

func FizzBuzz(i int) string {
	var ret string
	if (i%3 != 0) && (i%5 != 0) {
		ret = strconv.Itoa(i)
	} else {
		if i%3 == 0 {
			ret = "Fizz"
		}
		if i%5 == 0 {
			ret += "Buzz"
		}
	}
	return ret
}
