package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var input int
	rand.Seed(time.Now().UTC().UnixNano())
	number := rand.Intn(100)

	for i := 0; i < 5; i++ {
		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &input)
		if input < number {
			fmt.Println("น้อยไป")
		}
		if input > number {
			fmt.Println("มากไป")
		}
		if input == number {
			fmt.Println("เจอละ")
			return
		}
		if i == 4 {
			fmt.Println("จบการทำงาน", number)
		}
	}
}
