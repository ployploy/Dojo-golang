package main

import (
	"fmt"
	"sync"
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
	var wg sync.WaitGroup //เอาไว้เชคตัวรูทีนส่งกี่ครั้งเชคเท่านั้นห้ามขาดห้ามเกิน
	maxGoroutines := 4
	worker := 0
	for i := 1; i <= volumn; i++ {

		// cashier receive order
		time.Sleep(5 * time.Millisecond)
		coffee := fmt.Sprintf("order: %d", i)

		wg.Add(1)
		worker++

		go func() {
			defer wg.Done()
			//barista brew coffee
			time.Sleep(100 * time.Millisecond)
			coffee = fmt.Sprintf("%s %s", coffee, "espresso")

			//waiter serve coffee
			time.Sleep(5 * time.Millisecond)
			container = append(container, fmt.Sprintf("%s %s", coffee, "ready 🙂"))
			worker--
		}()

		if worker >= maxGoroutines || i >= volumn {
			wg.Wait()

		}
	}
	return
}
