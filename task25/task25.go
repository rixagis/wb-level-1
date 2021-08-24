package main

import (
	"fmt"
	"sync"
)

// Структура конкурентного счетчика
type counter struct {
	value int
	sync.Mutex
}

// Метод конкурентного увеличения значения счетчика
func (c *counter) inc(increment int) {
	c.Lock()
	c.value++
	fmt.Println("Incremented the value of the counter, now it equals ", c.value)
	c.Unlock()
}

// Воркер, который будет увеличивать значение счетчика
func worker(numWorker int, c *counter, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("Worker #%d increments now\n", numWorker)
		c.inc(1)
	}
}

func main(){
	var (
		c counter			// Счетчик
		wg sync.WaitGroup	// Группа ожидания завершения работы воркеров
		numWorkers = 5		// Количество воркеров
	)

	// Запускаем воркеров
	for i := 0; i < numWorkers; i++ {
		go worker(i + 1, &c, &wg)
		wg.Add(1)
	}

	wg.Wait()


}