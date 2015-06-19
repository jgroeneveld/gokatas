package gokatas

import (
	"errors"
	"testing"

	"github.com/jgroeneveld/trial/assert"
)

// http://de.scribd.com/doc/141109135/Function-Kata-ToDictionary
func TestToDictionary(t *testing.T) {
	dict, err := ToDictionary("a=1;b=2;c=3")
	assert.Nil(t, err)
	assert.DeepEqual(t, Dictionary{"a": "1", "b": "2", "c": "3"}, dict)

	dict, err = ToDictionary("a=1;a=2")
	assert.Nil(t, err)
	assert.DeepEqual(t, Dictionary{"a": "2"}, dict)

	dict, err = ToDictionary("a=")
	assert.Nil(t, err)
	assert.DeepEqual(t, Dictionary{"a": ""}, dict)

	dict, err = ToDictionary("")
	assert.Nil(t, err)
	assert.DeepEqual(t, Dictionary{}, dict)

	dict, err = ToDictionary("a==1")
	assert.Nil(t, err)
	assert.DeepEqual(t, Dictionary{"a": "=1"}, dict)

	dict, err = ToDictionary("=1")
	assert.DeepEqual(t, errors.New("Invalid Input"), err)
}
