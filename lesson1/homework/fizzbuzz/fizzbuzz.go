package fizzbuzz

import "strconv"

func FizzBuzz(i int) (result string) {
	switch {
	case i % 15 == 0: 
		result = "FizzBuzz"
	case i % 3 == 0:  
		result = "Fizz"
	case i % 5 == 0:  
		result = "Buzz"
	default:          
		result = strconv.Itoa(i)
	}
	return 
}
