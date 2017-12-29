package gotask

import "strconv"

func Decimal2Binary(digit int) string {
	var reminder int
	var output string

	for digit >= 1 {
		reminder = digit % 2
		output = strconv.Itoa(reminder) + output
		digit /= 2
	}

	return output
}
