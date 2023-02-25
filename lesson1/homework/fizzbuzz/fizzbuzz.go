package fizzbuzz

import "fmt"

func FizzBuzz(i int) string {

	var result string = ""

	if i%3 == 0 {
		result += "Fizz"
	}
	if i%5 == 0 {
		result += "Buzz"
	}

	if result != "" {
		return result
	} else {
		return string(i)
	}
}

func main() {
	for value := 0; value <= 100; value++ {
		fmt.Println(FizzBuzz(value))
	}
}
