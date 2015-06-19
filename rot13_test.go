package gokatas

import (
	"github.com/jgroeneveld/trial/assert"
	"testing"
)

// http://de.scribd.com/doc/142096464/Function-Kata-ROT
func TestROT13(t *testing.T) {
	assert.Equal(t, "URYYB, JBEYQ", ROT13("Hello, World"))
	assert.Equal(t, "NR", ROT13("ä"))
}
