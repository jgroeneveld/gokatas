package gokatas

import (
	gc "launchpad.net/gocheck"
)

// http://de.scribd.com/doc/140927641/Function-Kata-From-Roman-Numerals
type RomanNumeralsSuite struct{}

func (s *RomanNumeralsSuite) Test_FromRomanNumerals(c *gc.C) {
	c.Skip("Not yet implemented")

	tests := []struct {
		in  string
		out int
	}{
		{"I", 1},
		{"II", 2},
		{"IV", 4},
		{"V", 5},
		{"IX", 9},
		{"XLII", 42},
		{"XCIX", 99},
		{"MMXIII", 2013},
	}

	for _, test := range tests {
		c.Check(FromRomanNumerals(test.in), gc.Equals, test.out,
			gc.Commentf("for input '%s'", test.in))
	}
}

func (s *RomanNumeralsSuite) Test_IsValidRomanNumeral(c *gc.C) {
	c.Skip("Not yet implemented")
	tests := []struct {
		in  string
		out bool
	}{
		{"I", true},
		{"XCIX", true},
		{"I I", false}, // syntactic
		{"IC", false},  // semantic error
	}

	for _, test := range tests {
		c.Check(IsvalidRomanNumeral(test.in), gc.Equals, test.out,
			gc.Commentf("for input '%s'", test.in))
	}
}

var _ = gc.Suite(&RomanNumeralsSuite{})
