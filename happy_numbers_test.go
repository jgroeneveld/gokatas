package gokatas

import (
	"github.com/jgroeneveld/trial/assert"
	"testing"
)

// http://de.scribd.com/doc/142257204/Function-Kata-Fro%CC%88hliche-Zahlen
func TestIsHappyNumber(t *testing.T) {
	assert.Equal(t, false, IsHappyNumber(-7))
	assert.Equal(t, false, IsHappyNumber(0))
	assert.Equal(t, false, IsHappyNumber(21))
	assert.Equal(t, false, IsHappyNumber(2))
	assert.Equal(t, true, IsHappyNumber(7))
	assert.Equal(t, true, IsHappyNumber(964))
	assert.Equal(t, false, IsHappyNumber(999))
}

func TestCalculateHappyNumbers(t *testing.T) {
	assert.DeepEqual(t, []int{7, 10, 13, 19}, CalculateHappyNumbers(5, 20))
	assert.DeepEqual(t, []int{13, 19, 23, 28, 31, 32, 44}, CalculateHappyNumbers(13, 44))
}

func Test_sumOfSquaresOfDigits(t *testing.T) {
	assert.Equal(t, 49, sumOfSquaresOfDigits("7"))
	assert.Equal(t, 82, sumOfSquaresOfDigits("19"))
	assert.Equal(t, 68, sumOfSquaresOfDigits("82"))
	assert.Equal(t, 100, sumOfSquaresOfDigits("68"))
	assert.Equal(t, 1, sumOfSquaresOfDigits("100"))
}
