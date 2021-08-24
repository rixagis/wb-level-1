package main

import (
	"fmt"
	"sync"
)

// синхронизированный int для результата
type syncInt struct {
	value int
	mutex sync.Mutex
}

// метод для синхронизированного изменения syncInt
func (si *syncInt) add(increment int) {
	si.mutex.Lock()
	si.value += increment
	si.mutex.Unlock()
}

// горутина, отправляющая содержимое массива в канал
func send(input []int, output chan<- int) {
	defer close(output)
	for _, value := range input {
		output <- value
	}
}

// воркер, обрабатывающий данные из канала
func worker(input <-chan int, output *syncInt, wg *sync.WaitGroup) {
	defer wg.Done()
	for inputValue := range input {
		output.add(inputValue * inputValue)
	}
}

func main(){
	var data = []int{2, 4, 6, 8, 10} // входные данные по условию задачи
	var dataChannel = make(chan int) // канал входных данных для воркеров
	var result syncInt // переменная для результата
	var wg sync.WaitGroup
	var numOfWorkers = 4
	
	// активируем горутины
	go send(data, dataChannel)
	for i := 0; i < numOfWorkers; i++ {
		go worker(dataChannel, &result, &wg)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println("Результат = ", result.value)
}