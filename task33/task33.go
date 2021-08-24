package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// Производит quantity случайных чисел и отправляет их в output
func producer(output chan<- int, quantity int) {
	for i := 0; i < quantity; i++{
		output <- rand.Intn(100)
	}
	close(output)
}

// Проверяет числа из input на четность и, если четные, отправляет в output.
// Работает до закрытия канала input
func checker(input <-chan int, output chan<- int) {
	for value := range input {
		fmt.Println("Checker received:", value)
		if value % 2 == 0 {
			output <- value
		}
	}
	close(output)
}

// Печатает все числа, приходящие в input. Работает до закрытия канала input.
func printer(input <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for value := range input {
		fmt.Println("Printer received:", value)
	}
}

func main() {
	var (
		channelAll = make(chan int)	// канал для всех случайных чисел
		channelEven = make(chan int)	// канал для четных чисел
		wg sync.WaitGroup				// группа для ожидания завершения последней горутины
	)

	// запускаем горутины
	go producer(channelAll, 100)
	go checker(channelAll, channelEven)
	wg.Add(1)
	go printer(channelEven, &wg)

	wg.Wait()
}