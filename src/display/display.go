package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter number fahrenheit:")
	var fahrenheit int
	fmt.Scanf("%d", &fahrenheit)
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("Answer: %.d\n", celsius)

	number := strconv.Itoa(celsius)
	numtexts := strings.Split(number, "")

	displayDigitalNumber(numtexts)
	fmt.Println("Bankok")
}

func displayDigitalNumber(numbers []string) {
	for i := 0; i < len(numbers); i++ {
		displayFirstLine(numbers[i])
	}
	println()
	for i := 0; i < len(numbers); i++ {
		displaySecondLine(numbers[i])
	}
	println()
	for i := 0; i < len(numbers); i++ {
		displayThirdLine(numbers[i])
	}
	println()
}

func displayFirstLine(numberText string) {
	switch numberText {
	case "0", "2", "3", "5", "6", "7", "8", "9":
		fmt.Print(" _  ")
	case "1", "4", "15":
		fmt.Print("    ")
	}
}

func displaySecondLine(numberText string) {
	switch numberText {
	case "0":
		fmt.Print("| | ")
	case "4", "8", "9":
		fmt.Print("|_| ")
	case "1", "7":
		fmt.Print("  | ")
	case "5", "6":
		fmt.Print("|_  ")
	case "2", "3":
		fmt.Print(" _| ")
	}
}

func displayThirdLine(numberText string) {
	switch numberText {
	case "0", "6", "8":
		fmt.Print("|_| ")
	case "1", "4", "7":
		fmt.Print("  | ")
	case "2":
		fmt.Print("|_  ")
	case "3", "5", "9":
		fmt.Print(" _| ")

	}

}
