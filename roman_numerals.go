package gokatas

import "strings"

func FromRomanNumerals(roman string) int {
	roman, subSum := romanSumForMap(roman, romanSubtractionValues())
	roman, addSum := romanSumForMap(roman, romanAdditionValues())

	return subSum + addSum
}

func IsvalidRomanNumeral(roman string) bool {
	roman = removeSubtractionalNumerals(roman)

	if !isValidRomanNumeralSyntactically(roman) {
		return false
	}

	return isValidRomanNumeralSemantically(roman)
}

func removeSubtractionalNumerals(roman string) string {
	roman, _ = romanSumForMap(roman, romanSubtractionValues())
	return roman
}

func isValidRomanNumeralSyntactically(roman string) bool {
	for _, r := range roman {
		if !isValidRomanRune(r) {
			return false
		}
	}

	return true
}

func isValidRomanRune(r rune) bool {
	return romanAdditionValues()[string(r)] != 0
}

func isValidRomanNumeralSemantically(roman string) bool {
	numbers := make([]int, len(roman))
	for i, r := range roman {
		number := romanAdditionValues()[string(r)]
		numbers[i] = number
	}

	last := 9999
	for _, n := range numbers {
		if n > last {
			return false
		}
		last = n
	}

	return true
}

func romanSubtractionValues() map[string]int {
	return map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
}

func romanAdditionValues() map[string]int {
	return map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
}

func romanSumForMap(roman string, dict map[string]int) (string, int) {
	sum := 0
	for k, v := range dict {
		sum += strings.Count(roman, k) * v
		roman = strings.Replace(roman, k, "", -1)
	}
	return roman, sum
}
