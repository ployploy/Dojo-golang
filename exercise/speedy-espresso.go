package main

import (
	"fmt"
	"time"
)

func main() {
	volumn := 200
	start := time.Now()

	container := order(volumn)
	for _, cup := range container {
		fmt.Println(cup)
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func order(volumn int) (container []string) {
	for i := 1; i <= volumn; i++ {
		//cashire receive order
		time.Sleep(5 * time.Millisecond)
		coffee := fmt.Sprint("order: %d", i)

		//barista brew coffee
		time.Sleep(100 * time.Millisecond)
		coffee = fmt.Sprintf("%s %s", coffee, "espresso")

		//waiter serve coffee
		time.Sleep(5 * time.Millisecond)
		container = append(container, fmt.Sprintf("%s %s", coffee, "ready :)"))
	}
	return
}
