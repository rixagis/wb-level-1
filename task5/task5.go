package main

import (
	"fmt"
	"math/rand"
	"time"
)

// горутина, пишущая случайные числа в канал
func writer (output chan<- int, done <-chan struct{}) {
	for {
		select {
		case <- done:
			close(output)
			break;
		default:
			value := rand.Intn(100)
			fmt.Println("Writer writes: ", value)
			output <- value
			time.Sleep(250 * time.Millisecond)
		}
	}
}

// горутина, читающая из канала
func reader (input <- chan int) {
	for value := range input {
		fmt.Println("Reader reads: ", value)
	}
}

func main() {
	done := make(chan struct{}) // контрольный канал для закрытия пишущей горутины
	exchange := make(chan int) // канал для обмена данными между горутинами
	N := 5 // количество секунд для работы программы

	go writer(exchange, done)
	go reader(exchange)

	time.Sleep(time.Duration(N) * time.Second)
	done <- struct{}{} // закрываем пишущую горутину
	
}