package gokatas

import "strconv"

// http://de.scribd.com/doc/140817312/Function-Kata-FizzBuzz
func FizzBuzz() []string {
	result := make([]string, 100)

	for i := 0; i < 100; i++ {
		var value string
		vi := i + 1

		mod3 := vi%3 == 0
		mod5 := vi%5 == 0

		if mod3 && mod5 {
			value = "FizzBuzz"
		} else if mod3 {
			value = "Fizz"
		} else if mod5 {
			value = "Buzz"
		} else {
			value = strconv.Itoa(vi)
		}

		result[i] = value
	}

	return result
}
