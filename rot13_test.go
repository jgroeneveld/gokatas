package gokatas

import (
	gc "launchpad.net/gocheck"
)

type ROT13Suite struct{}

// http://de.scribd.com/doc/142096464/Function-Kata-ROT
func (s *ROT13Suite) Test_ROT13(c *gc.C) {
	tests := []struct {
		in  string
		out string
	}{
		{"Hello, World", "URYYB, JBEYQ"},
		{"ä", "NR"},
	}

	for _, test := range tests {
		c.Check(ROT13(test.in), gc.Equals, test.out,
			gc.Commentf("for input '%s'", test.in))
	}
}

var _ = gc.Suite(&ROT13Suite{})
