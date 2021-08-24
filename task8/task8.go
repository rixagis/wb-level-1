package main

import (
	"fmt"
	"math/rand"
)

// устанавливает последний бит в 1
func setToOne (number int64) int64 {
	return number | 1
}

// устанавливает последний бит в 0
func setToZero (number int64) int64 {
	return number &^ 1
}

func main() {
	// проверка для 5 случайных значений
	for i := 0; i < 5; i++ {
		value := rand.Int63()
		fmt.Printf("Initial number = %x, last bit set to 0: %x, last bit set to 1: %x\n", value, setToZero(value), setToOne(value))
	}
}