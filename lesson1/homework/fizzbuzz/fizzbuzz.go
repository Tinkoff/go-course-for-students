package fizzbuzz

import "strconv"

/*
Напишите программу, которая выводит на экран числа от 1 до 100.

При этом вместо чисел, кратных трем, программа должна выводить слово Fizz, а вместо чисел, кратных пяти — слово Buzz.

Если число кратно пятнадцати, то программа должна выводить слово FizzBuzz.
*/

func FizzBuzz(i int) string {
	switch {
	case i%3 == 0 && i%15 != 0:
		return "Fizz"
	case i%5 == 0 && i%15 != 0:
		return "Buzz"
	case i%15 == 0:
		return "FizzBuzz"
	default:
		return strconv.Itoa(i)
	}
}
