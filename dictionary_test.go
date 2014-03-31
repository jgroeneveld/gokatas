package gokatas

import (
	"errors"

	gc "launchpad.net/gocheck"
)

type DictionarySuite struct{}

// http://de.scribd.com/doc/141109135/Function-Kata-ToDictionary
func (s *DictionarySuite) Test_ToDictionary(c *gc.C) {
	tests := []struct {
		in  string
		out Dictionary
		err error
	}{
		{"a=1;b=2;c=3", Dictionary{"a": "1", "b": "2", "c": "3"}, nil},
		{"a=1;a=2", Dictionary{"a": "2"}, nil},
		{"a=", Dictionary{"a": ""}, nil},
		{"=1", nil, errors.New("Invalid Input")},
		{"", Dictionary{}, nil},
		{"a==1", Dictionary{"a": "=1"}, nil},
	}

	for _, test := range tests {
		dict, err := ToDictionary(test.in)

		c.Check(dict, gc.DeepEquals, test.out,
			gc.Commentf("for input %s", test.in))

		c.Check(err, gc.DeepEquals, test.err,
			gc.Commentf("for input %s", test.in))
	}
}

var _ = gc.Suite(&DictionarySuite{})
