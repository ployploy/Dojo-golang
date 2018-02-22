package test

func printDimond(input string) string {
	chars := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	count := len(chars)
	for index, value := range chars {
		if value == input {
			count = index + 1

		}
	}
	columnWidth := (count * 2)
	centerPoint := count

	printText := ""

	for rowTop := 0; rowTop < count; rowTop++ {
		for columnTop := 1; columnTop < columnWidth; columnTop++ {
			left := centerPoint - rowTop
			right := centerPoint + rowTop
			if columnTop == left {
				printText += chars[rowTop]
			} else if columnTop == right {
				printText += chars[rowTop]
			} else {
				printText += " "
			}
		}
		printText += "\n"
	}

	for rowDown := count - 1; rowDown > 0; rowDown-- {
		for columnDown := 1; columnDown < columnWidth; columnDown++ {
			left := centerPoint - (rowDown - 1)
			right := centerPoint + (rowDown - 1)
			if columnDown == left {
				printText += chars[rowDown-1]
			} else if columnDown == right {
				printText += chars[rowDown-1]
			} else {
				printText += " "
			}
		}
		printText += "\n"
	}

	return printText
}
