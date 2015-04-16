package gokatas

import (
	"testing"

	"github.com/heroku/hk/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

// http://de.scribd.com/doc/140927641/Function-Kata-From-Roman-Numerals
func TestFromRomanNumerals_Addition(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"I", 1},
		{"II", 2},
		{"V", 5},
		{"MMXIII", 2013},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, FromRomanNumerals(test.in), "for %q", test.in)
	}
}

func TestFromRomanNumerals_Subtraction(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"IV", 4},
		{"IX", 9},
		{"XLII", 42},
		{"XCIX", 99},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, FromRomanNumerals(test.in), "for %q", test.in)
	}
}

func TestIsValidRomanNumeral(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"I", true},
		{"XCIX", true},
		{"I I", false}, // syntactic error
		{"I,I", false},
		{"IAI", false},
		{"IC", false}, // semantic error
	}

	for _, test := range tests {
		assert.Equal(t, test.out, IsvalidRomanNumeral(test.in), "for %q", test.in)
	}
}
