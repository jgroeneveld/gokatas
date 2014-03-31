package gokatas

import (
	"strings"
)

func FromRomanNumerals(roman string) int {
	roman, subSum := romanSumForMap(roman, romanSubtractionValues())
	roman, addSum := romanSumForMap(roman, romanAdditionValues())

	return subSum + addSum
}

func IsvalidRomanNumeral(roman string) bool {
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
