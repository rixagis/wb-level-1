package main

import (
	"fmt"
)

func main() {
	var (
		a = 5
		b = 10
	)
	
	fmt.Printf("До обмена: a = %d, b = %d\n", a, b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("После обмена: a = %d, b = %d", a, b)


}