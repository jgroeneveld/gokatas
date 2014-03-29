default: test

package = github.com/jgroeneveld/gokatas

.PHONY: default test format fizz_buzz roman_numerals happy_numbers

test:
	go test $(package)/...

fizz_buzz:
	go test -gocheck.f "FizzBuzz*"

roman_numerals:
	go test -gocheck.f "RomanNumerals*"

happy_numbers:
	go test -gocheck.f "HappyNumbers*"

format:
	goimports -w .
