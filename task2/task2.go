package main

import (
	"fmt"
	"sync"
)

// горутина, отправляющая содержимое массива в канал
func send(input []int, output chan<- int) {
	defer close(output)
	for _, value := range input {
		output <- value
	}
}

// воркер, обрабатывающий данные из канала
func worker(workerNum int, input <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for inputValue := range input {
		fmt.Printf("Worker #%d: result = %d\n", workerNum, inputValue * inputValue)	
	}
}

func main(){
 	var data = []int{2, 4, 6, 8, 10} // входные данные по условию задачи
	var dataChannel = make(chan int) // канал входных данных для воркеров
	var wg sync.WaitGroup
	var numOfWorkers = 4
	
	// активируем горутины
	go send(data, dataChannel)
	for i := 0; i < numOfWorkers; i++ {
		go worker(i+1, dataChannel, &wg)
		wg.Add(1)
	}

	wg.Wait()

}