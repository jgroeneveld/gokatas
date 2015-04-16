package gokatas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// http://de.scribd.com/doc/142096464/Function-Kata-ROT
func TestROT13(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"Hello, World", "URYYB, JBEYQ"},
		{"ä", "NR"},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, ROT13(test.in), "for %q", test.in)
	}
}
