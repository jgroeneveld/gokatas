package gokatas

import (
	"github.com/jgroeneveld/trial/assert"
	"testing"
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
	assert.DeepEqual(t, LOCData{
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
	assert.True(t, isWhitespaceLine(""))
	assert.True(t, isWhitespaceLine("   "))
	assert.True(t, isWhitespaceLine("	"))
	assert.False(t, isWhitespaceLine("	//"))
	assert.False(t, isWhitespaceLine("x	//"))
}
