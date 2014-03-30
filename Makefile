default: all

.PHONY: default all format fizz_buzz roman_numerals happy_numbers

all:
	go test

fizz_buzz:
	go test -gocheck.f "FizzBuzz*"

roman_numerals:
	go test -gocheck.f "RomanNumerals*"

happy_numbers:
	go test -gocheck.f "HappyNumbers*"

loc:
	go test -gocheck.f "LOC*"

format:
	goimports -w .
