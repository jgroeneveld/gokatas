default: all

.PHONY: default all format fizz_buzz roman_numerals happy_numbers coverage

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

rot13:
	go test -gocheck.f "ROT13*"

coverage:
	go test -test.coverprofile=/tmp/cov.out && go tool cover -html=/tmp/cov.out

format:
	goimports -w .
