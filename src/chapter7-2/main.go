package main

import (
	"fmt"
)

//slice ใช้เยอะกับพวกเรียก 3rd party
func main() {
	slice := make([]int, 3)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3

	fmt.Println(slice)

	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice2)

	fmt.Println("Slice with lenght and capacity")
	fmt.Printf("slice: length %v, capacity %v, %v\n",
		len(slice), cap(slice), slice)

	//append
	for i := 4; i < 10; i++ {
		slice = append(slice, i)
	}
	fmt.Printf("slice: length %v, capacity %v, %v\n",
		len(slice), cap(slice), slice)
}
