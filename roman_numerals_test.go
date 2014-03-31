package gokatas

import (
	gc "launchpad.net/gocheck"
)

// http://de.scribd.com/doc/140927641/Function-Kata-From-Roman-Numerals
type RomanNumeralsSuite struct{}

func (s *RomanNumeralsSuite) Test_FromRomanNumeralsAddition(c *gc.C) {
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
		c.Check(FromRomanNumerals(test.in), gc.Equals, test.out,
			gc.Commentf("for input '%s'", test.in))
	}
}

func (s *RomanNumeralsSuite) Test_FromRomanNumeralsSubtraction(c *gc.C) {
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
		c.Check(FromRomanNumerals(test.in), gc.Equals, test.out,
			gc.Commentf("for input '%s'", test.in))
	}
}

func (s *RomanNumeralsSuite) Test_IsValidRomanNumeral(c *gc.C) {
	tests := []struct {
		in  string
		out bool
	}{
		{"I", true},
		{"XCIX", true},
		{"I I", false}, // syntactic error
		{"I,I", false},
		{"IAI", false},
		{"IC", false},  // semantic error
	}

	for _, test := range tests {
		c.Check(IsvalidRomanNumeral(test.in), gc.Equals, test.out,
			gc.Commentf("for input '%s'", test.in))
	}
}

var _ = gc.Suite(&RomanNumeralsSuite{})
