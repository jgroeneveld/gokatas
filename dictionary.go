package gokatas

import "errors"
import "strings"

type Dictionary map[string]string

func ToDictionary(input string) (Dictionary, error) {
	dict := make(Dictionary)

	if len(input) == 0 {
		return dict, nil
	}

	assignments := strings.Split(input, ";")

	for _, a := range assignments {
		parts := strings.Split(a, "=")
		key := ""
		value := ""

		if len(parts) >= 1 && len(parts[0]) > 0 {
			key = parts[0]

			if len(parts) >= 2 {
				value = strings.Join(parts[1:], "=")
			}

			dict[key] = value
		} else {
			return nil, errors.New("Invalid Input")
		}
	}

	return dict, nil
}
