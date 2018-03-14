package string

import (
	"strconv"
)

func ConvertString(input []int) string {
	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			//for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
			input[i], input[j] = input[j], input[i]

		}

		output := ""
		for _, v := range input {
			output += strconv.Itoa(v)
		}
		return output
	}
	return ""
}
