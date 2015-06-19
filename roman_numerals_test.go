package gokatas

import (
	"testing"

	"github.com/jgroeneveld/trial/assert"
)

// http://de.scribd.com/doc/140927641/Function-Kata-From-Roman-Numerals
func TestFromRomanNumerals_Addition(t *testing.T) {
	assert.Equal(t, 1, FromRomanNumerals("I"))
	assert.Equal(t, 2, FromRomanNumerals("II"))
	assert.Equal(t, 5, FromRomanNumerals("V"))
	assert.Equal(t, 2013, FromRomanNumerals("MMXIII"))
}

func TestFromRomanNumerals_Subtraction(t *testing.T) {
	assert.Equal(t, 4, FromRomanNumerals("IV"))
	assert.Equal(t, 9, FromRomanNumerals("IX"))
	assert.Equal(t, 42, FromRomanNumerals("XLII"))
	assert.Equal(t, 99, FromRomanNumerals("XCIX"))
}

func TestIsValidRomanNumeral(t *testing.T) {
	assert.True(t, IsValidRomanNumeral("I"))
	assert.True(t, IsValidRomanNumeral("XCIX"))
	assert.False(t, IsValidRomanNumeral("I I"), "treating syntactic error as valid")
	assert.False(t, IsValidRomanNumeral("I,I"))
	assert.False(t, IsValidRomanNumeral("IAI"))
	assert.False(t, IsValidRomanNumeral("IC"), "treating semantic error as valid")
}
