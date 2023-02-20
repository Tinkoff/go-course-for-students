package fizzbuzz_test

import (
	"github.com/stretchr/testify/assert"
	"lecture01_homework/fizzbuzz"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	var res string

	res = fizzbuzz.FizzBuzz(9)
	assert.Equal(t, res, "Fizz")

	res = fizzbuzz.FizzBuzz(25)
	assert.Equal(t, res, "Buzz")

	res = fizzbuzz.FizzBuzz(30)
	assert.Equal(t, res, "FizzBuzz")

	res = fizzbuzz.FizzBuzz(11)
	assert.Equal(t, res, "11")

	res = fizzbuzz.FizzBuzz(26)
	assert.Equal(t, res, "26")

	res = fizzbuzz.FizzBuzz(31)
	assert.Equal(t, res, "31")
}
