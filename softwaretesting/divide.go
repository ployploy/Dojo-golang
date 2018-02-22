package test

import "strconv"

func divide(number int) string {
	if number == 102 {
		return "no show"
	}
	if number%3 == 0 {
		return strconv.Itoa(number)
	}
	return "no display"
}
