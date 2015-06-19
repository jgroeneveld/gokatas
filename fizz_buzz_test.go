package gokatas

import (
	"testing"

	"github.com/jgroeneveld/trial/assert"
)

// http://de.scribd.com/doc/140817312/Function-Kata-FizzBuzz
func TestFizzBuzz_Length(t *testing.T) {
	result := FizzBuzz()

	assert.Equal(t, 100, len(result))
}

func TestFizzBuzz_Contents(t *testing.T) {
	result := FizzBuzz()
	expectedStart := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7"}
	expectedMid := []string{"14", "FizzBuzz", "16"}

	assert.DeepEqual(t, expectedStart, result[:7])
	assert.DeepEqual(t, expectedMid, result[13:16])
}
