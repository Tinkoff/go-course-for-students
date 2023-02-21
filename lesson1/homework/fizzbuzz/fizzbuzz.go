package fizzbuzz

import "strconv"

func FizzBuzz(i int) string {
	fizz, buzz := "Fizz", "Buzz"

	if i%3 == 0 && i%5 == 0 {
		return fizz + buzz
	} else if i%3 == 0 {
		return fizz
	} else if i%5 == 0 {
		return buzz
	}
	return strconv.Itoa(i)
}
