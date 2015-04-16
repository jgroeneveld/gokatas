package gokatas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLOC(t *testing.T) {
	codeToTest := `/* this is the code to test
the first lines are a documentation comment
*/

// after this comes a function
func name(params) int {
	// indented comment
	return 42
}`

	loc := GetLOC(codeToTest)
	assert.Equal(t, LOCData{
		Total:      9,
		Code:       3,
		Comments:   5,
		Whitespace: 1,
	}, loc)
}

func TestStripLines(t *testing.T) {
	unstripped := `first
  second
     further in
third`

	stripped := `first
second
further in
third`

	assert.Equal(t, stripped, stripLines(unstripped))
}

func TestIsWhitespaceLine(t *testing.T) {
	assert := assert.New(t)
	assert.True(isWhitespaceLine(""))
	assert.True(isWhitespaceLine("   "))
	assert.True(isWhitespaceLine("	"))
	assert.False(isWhitespaceLine("	//"))
	assert.False(isWhitespaceLine("x	//"))
}
