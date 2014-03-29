package gokatas

import (
	"testing"
	. "launchpad.net/gocheck"
)

func Test(t *testing.T) { TestingT(t) }

type FizzBuzzSuite struct{}
var _ = Suite(&FizzBuzzSuite{})

func (s *FizzBuzzSuite) Test_FizzBuzzLength(c *C) {
	result := FizzBuzz()

	c.Check(len(result), Equals,100)
}

func (s *FizzBuzzSuite) Test_FizzBuzzContents(c *C) {
	result := FizzBuzz()
	expectedStart := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7"}
	expectedMid := []string{"14", "FizzBuzz", "16"}

	c.Check(result[:7], DeepEquals, expectedStart)
	c.Check(result[13:16], DeepEquals, expectedMid)
}
