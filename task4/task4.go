package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
	"math/rand"
)

// Worker - структура воркера
type Worker struct {
	numWorker int					// порядковый номер воркера
	dataChannel <-chan int			// канал для приема данных
	controlChannel chan struct{}	// канал для команды завершения, у каждого воркера свой
	wg *sync.WaitGroup				// waitgroup для ожидания завершения воркеров
}

// Run - метод запуска воркера
func (w *Worker) Run() {
	defer w.wg.Done()
	for {
		select {
		case value := <- w.dataChannel:
			fmt.Printf("Worker #%d reads %d\n", w.numWorker, value)
			time.Sleep(500 * time.Millisecond)
		case <- w.controlChannel:
			fmt.Printf("Worker #%d is stopping...\n", w.numWorker)
			return
		}
	}
}

// Stop - метод остановки воркера
func (w *Worker) Stop() {
	w.controlChannel <- struct{}{}
}

// NewWorker - "конструктор" воркера
func NewWorker(numWorker int, dataChannel <- chan int, wg *sync.WaitGroup) *Worker {
	return &Worker{numWorker, dataChannel, make(chan struct{}), wg}
}

// Функция, обрабатывающая сигналы. При получении сигнала приказывает воркерам завершиться и ждет их завершения,
// после чего завершает программу.
// Слайс воркеров по указателю для того, чтобы можно было менять слайс воркеров, которых нужно завершить при сигнале, динамически.
func processSignal(signals <-chan os.Signal, workers *[]*Worker, wg *sync.WaitGroup){
	for {
		select {
		case <-signals:
			fmt.Println("Received signal, shutting down...")
			for _, worker := range *workers {
				worker.Stop()
			}
			wg.Wait()
			fmt.Println("All workers stopped, exiting program.")
			os.Exit(0)
		}
	}
}


func main () {
	var (
		numWorkers = 5							// количество воркеров
		dataChannel = make(chan int)			// канал для данных
		signalChannel = make(chan os.Signal)	// канал сигналов для обработки
		wg sync.WaitGroup						// группа ожидания воркеров
		workers = make([]*Worker, numWorkers)	// слайс всех воркеров
	)

	// создаем воркеров
	for i := 0; i < numWorkers; i++ {
		workers[i] = NewWorker(i+1, dataChannel, &wg)
	}

	// запускаем обработчик сигнала и регистрируем его
	go processSignal(signalChannel, &workers, &wg)
	signal.Notify(signalChannel)

	// запускаем воркеров
	for _, worker := range workers {
		wg.Add(1)
		go worker.Run()
	}

	// бесконечный цикл подачи данных для обработки воркерами
	for {
		dataChannel <- rand.Intn(100)
	}

}