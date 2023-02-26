package fizzbuzz_test

import (
	"github.com/stretchr/testify/assert"
	"lecture01_homework/fizzbuzz"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	test := func(input int, actual string) {
		res := fizzbuzz.FizzBuzz(input)
		assert.Equal(t, res, actual)
	}

	test(9, "Fizz")
	test(25, "Buzz")
	test(30, "FizzBuzz")
	test(11, "11")
	test(26, "26")
	test(31, "31")
}
