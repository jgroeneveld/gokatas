package gokatas

import (
	gc "launchpad.net/gocheck"
)

type FizzBuzzSuite struct{}

func (s *FizzBuzzSuite) Test_Length(c *gc.C) {
	result := FizzBuzz()

	c.Check(len(result), gc.Equals, 100)
}

func (s *FizzBuzzSuite) Test_Contents(c *gc.C) {
	result := FizzBuzz()
	expectedStart := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7"}
	expectedMid := []string{"14", "FizzBuzz", "16"}

	c.Check(result[:7], gc.DeepEquals, expectedStart)
	c.Check(result[13:16], gc.DeepEquals, expectedMid)
}

var _ = gc.Suite(&FizzBuzzSuite{})
