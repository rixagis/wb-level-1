package main

import (
	"fmt"
)

func main(){
	var (
		sliceWith100Length = make([]int, 100)		// Слайс с 100 элементами
		sliceWith100Capacity = make([]int, 0, 100)	// Слайс с 0 элементами и емкостью 100
	)

	fmt.Printf("Slice with 100 elements: len = %d, cap = %d\n", len(sliceWith100Length), cap(sliceWith100Length))
	fmt.Printf("Slice with 100 capacity: len = %d, cap = %d\n", len(sliceWith100Capacity), cap(sliceWith100Capacity))
}