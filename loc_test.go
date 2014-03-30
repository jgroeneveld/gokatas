package gokatas

import gc "launchpad.net/gocheck"

type LOCSuite struct{}

func (s *LOCSuite) Test_GetLOC(c *gc.C) {
	codeToTest := `/* this is the code to test
the first lines are a documentation comment
*/

// after this comes a function
func name(params) int {
	// indented comment
	return 42
}`

	loc := GetLOC(codeToTest)

	c.Check(loc, gc.Equals, LOCData{
		Total:      9,
		Code:       3,
		Comments:   5,
		Whitespace: 1,
	})
}

func (s *LOCSuite) Test_stripLines(c *gc.C) {
	unstripped := `first
  second
     further in
third`

	stripped := `first
second
further in
third`

	c.Check(stripLines(unstripped), gc.Equals, stripped)
}

func (s *LOCSuite) Test_isWhitespaceLine(c *gc.C) {
	c.Check(isWhitespaceLine(""), gc.Equals, true)
	c.Check(isWhitespaceLine("   "), gc.Equals, true)
	c.Check(isWhitespaceLine("	"), gc.Equals, true)
	c.Check(isWhitespaceLine("	//"), gc.Equals, false)
	c.Check(isWhitespaceLine("x	//"), gc.Equals, false)
}

var _ = gc.Suite(&LOCSuite{})
