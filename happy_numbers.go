package gokatas

import (
	"math"
	"strconv"
)

func IsHappyNumber(number int) bool {
	return isHappyNumberLimited(number, []int{})
}

func CalculateHappyNumbers(start, end int) []int {
	happyNumbers := []int{}

	for i := start; i <= end; i++ {
		if IsHappyNumber(i) {
			happyNumbers = append(happyNumbers, i)
		}
	}

	return happyNumbers
}

func isHappyNumberLimited(number int, triedNumbers []int) bool {
	if number > 1 && !includesInt(triedNumbers, number) {
		numberAsString := strconv.Itoa(number)
		nextNumber := sumOfSquaresOfDigits(numberAsString)
		triedNumbers = append(triedNumbers, number)
		return isHappyNumberLimited(nextNumber, triedNumbers)
	}

	return number == 1
}

func sumOfSquaresOfDigits(number string) int {
	sum := 0
	for _, rune := range number {
		digit, err := strconv.Atoi(string(rune))
		if err != nil {
			panic(err)
		}
		sum += int(math.Pow(float64(digit), float64(2)))
	}
	return sum
}

func includesInt(slice []int, e1 int) bool {
	for _, e2 := range slice {
		if e2 == e1 {
			return true
		}
	}
	return false
}
