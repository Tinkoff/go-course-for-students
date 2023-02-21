package fizzbuzz_test

import (
	"lecture01_homework/fizzbuzz"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestFizzBuzz2(t *testing.T) {
	testCases := []struct {
		name        string
		number      int
		expectedAns string
	}{
		{
			name:        "simple number",
			number:      4,
			expectedAns: "4",
		},
		{
			name:        "Fizz",
			number:      27,
			expectedAns: "Fizz",
		},
		{
			name:        "Buzz",
			number:      25,
			expectedAns: "Buzz",
		},
		{
			name:        "FizzBuzz",
			number:      45,
			expectedAns: "FizzBuzz",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := fizzbuzz.FizzBuzz(tc.number)
			assert.Equal(t, tc.expectedAns, res)
		})
	}
}