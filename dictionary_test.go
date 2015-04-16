package gokatas

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// http://de.scribd.com/doc/141109135/Function-Kata-ToDictionary
func TestToDictionary(t *testing.T) {
	assert := assert.New(t)

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

		assert.Equal(test.out, dict, "for %q", test.in)
		assert.Equal(test.err, err, "for %q", test.in)
	}
}
