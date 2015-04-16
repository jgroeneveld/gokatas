package gokatas

import (
	"testing"

	"github.com/heroku/hk/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

// http://de.scribd.com/doc/142257204/Function-Kata-Fro%CC%88hliche-Zahlen
func TestIsHappyNumber(t *testing.T) {
	tests := []struct {
		in  int
		out bool
	}{
		{-7, false},
		{0, false},
		{1, true},
		{2, false},
		{7, true},
		{964, true},
		{999, false},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, IsHappyNumber(test.in), "for %s", test.in)
	}
}

func TestCalculateHappyNumbers(t *testing.T) {
	assert.Equal(t, []int{7, 10, 13, 19}, CalculateHappyNumbers(5, 20))
	assert.Equal(t, []int{13, 19, 23, 28, 31, 32, 44}, CalculateHappyNumbers(13, 44))
}

func Test_sumOfSquaresOfDigits(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"7", 49},
		{"19", 82},
		{"82", 68},
		{"68", 100},
		{"100", 1},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, sumOfSquaresOfDigits(test.in), "for %s", test.in)
	}
}
