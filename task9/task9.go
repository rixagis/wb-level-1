package main

import (
	"fmt"
	"sync"
)

// записывает числа из массива в первый канал
func writer(input []int, output chan<- int) {
	defer close(output)
	for _, value := range input {
		fmt.Println("Writer writes: ", value)
		output <- value
	}
}

// читает числа из первого канала, умножает и записывает во второй
func multiplier(input <-chan int, output chan<- int) {
	defer close(output)
	for value := range input {
		fmt.Println("Multiplier multiplies: ", value, " and writes: ", value * 2)
		output <- value * 2
	}
}

// читает числа из второго канала
func reader(input <- chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for value := range input {
		fmt.Println("Reader reads: ", value)
	}
}

func main(){
	var data = []int{3, 4, 2, 4, 23, 43, 1, 23, 0, 35}	// изначальный массив с числами
	var chan1 = make(chan int)	// канал для записи чисел из массива
	var chan2 = make(chan int)	// канал для записи умноженных чисел
	var wg sync.WaitGroup		// группа для ожидания завершения горутин

	// запускаем горутины
	go writer(data, chan1)
	go multiplier(chan1, chan2)
	go reader(chan2, &wg)
	wg.Add(1)

	wg.Wait()
	
}