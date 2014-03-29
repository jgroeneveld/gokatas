package gokatas

import (
	gc "launchpad.net/gocheck"
)

// http://de.scribd.com/doc/140927641/Function-Kata-From-Roman-Numerals
type FromRomanNumeralsSuite struct{}

func (s *FromRomanNumeralsSuite) Test_Numerals(c *gc.C) {
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
			gc.Commentf("for input %s", test.in))
	}
}

var _ = gc.Suite(&FromRomanNumeralsSuite{})
