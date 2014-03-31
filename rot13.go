package gokatas

import "strings"

const alphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func ROT13(input string) string {
	input = strings.ToUpper(input)
	input = replaceUmlauts(input)
	input = doROT(input)
	return input
}

func doROT(input string) string {
	output := make([]rune, len(input))
	for i, r := range input {
		indexInAlphabet := strings.IndexRune(alphabet, r)
		if indexInAlphabet >= 0 {
			newIndex := (indexInAlphabet + 13) % len(alphabet)
			output[i] = rune(alphabet[newIndex])
		} else {
			output[i] = r
		}
	}
	return string(output)
}

func replaceUmlauts(input string) string {
	input = strings.Replace(input, "Ö", "OE", -1)
	input = strings.Replace(input, "Ä", "AE", -1)
	input = strings.Replace(input, "Ü", "UE", -1)
	input = strings.Replace(input, "ß", "SS", -1)
	return input
}
