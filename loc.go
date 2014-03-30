package gokatas

import (
	"regexp"
	"strings"
)

type LOCData struct {
	Total      int
	Code       int
	Comments   int
	Whitespace int
}

func GetLOC(code string) LOCData {
	return LOCData{
		Total:      getTotalLOC(code),
		Code:       getCodeLOC(code),
		Comments:   getCommentsLOC(code),
		Whitespace: getWhitespaceLOC(code),
	}
}

func getTotalLOC(code string) int {
	return len(strings.Split(code, "\n"))
}

func getCodeLOC(code string) int {
	return getTotalLOC(code) - getCommentsLOC(code) - getWhitespaceLOC(code)
}

func getCommentsLOC(code string) int {
	code = stripLines(code)
	count := 0

	startsCodeBlock := regexp.MustCompile(`/\*`)
	endsCodeBlock := regexp.MustCompile(`\*/`)
	lineComment := regexp.MustCompile(`^//`)

	findEnd := false

	for _, line := range strings.Split(code, "\n"){
		if findEnd {
			count++
			if endsCodeBlock.MatchString(line) {
				findEnd = false
			}
		} else {
			if startsCodeBlock.MatchString(line) {
				count++
				findEnd = true
			} else if lineComment.MatchString(line) {
				count++
			}
		}
	}

	return count
}

func getWhitespaceLOC(code string) int {
	count := 0

	for _, line := range strings.Split(code, "\n") {
		if isWhitespaceLine(line) {
			count++
		}
	}

	return count
}

func stripLines(code string) string {
	regex := regexp.MustCompile(`(?m)^\s*`)
	newCode := regex.ReplaceAllLiteralString(code, "")
	return newCode
}

func isWhitespaceLine(line string) bool {
	regex := regexp.MustCompile(`^\s*$`)
	return regex.MatchString(line)
}
