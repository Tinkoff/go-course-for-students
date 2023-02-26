package fizzbuzz

import "strconv"

var primeDict = []struct {
	prime int
	word  string
}{
	{3, "Fizz"},
	{5, "Buzz"},
}

func FizzBuzz(i int) string {
	var str = ""

	for _, e := range primeDict {
		if i%e.prime == 0 {
			str += e.word
		}
	}

	if str == "" {
		str = strconv.Itoa(i)
	}

	return str
}
