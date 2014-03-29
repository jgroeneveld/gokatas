package gokatas

import (
	gc "launchpad.net/gocheck"
)

// http://de.scribd.com/doc/142257204/Function-Kata-Fro%CC%88hliche-Zahlen
type HappyNumbersSuite struct{}

func (s *HappyNumbersSuite) Test_IsHappyNumber(c *gc.C) {
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
		c.Check(IsHappyNumber(test.in), gc.Equals, test.out,
			gc.Commentf("for input %d", test.in))
	}
}

func (s *HappyNumbersSuite) Test_CalculateHappyNumbers(c *gc.C) {
	c.Check(CalculateHappyNumbers(5, 20), gc.DeepEquals, []int{7, 10, 13, 19})
	c.Check(CalculateHappyNumbers(13, 44), gc.DeepEquals, []int{13, 19, 23, 28, 31, 32, 44})
}

func (s *HappyNumbersSuite) Test_sumOfSquaresOfDigits(c *gc.C) {
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
		c.Check(sumOfSquaresOfDigits(test.in), gc.Equals, test.out,
			gc.Commentf("for input %s", test.in))
	}
}


var _ = gc.Suite(&HappyNumbersSuite{})
