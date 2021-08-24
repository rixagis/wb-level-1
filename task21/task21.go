	package main

import (
	"fmt"
	"sync"
)

// воркер, который читает элементы массива с start (включая) по end (не включая)
func worker(numWorker int, arr []int, start, end int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; i < end; i++ {
		fmt.Printf("Worker #%d reads element %d with index %d\n", numWorker, arr[i], i)
	}
	
}

func main(){
	var (
		data = []int{234, 23, 25, 354, 286, 73248, 422, 53, 105, 574}	// массив для прочтения
		numOfWorkers = 5												// количество воркеров
		partSize = len(data) / numOfWorkers							// размер доли воркера
		start = 0														// стартовый индекс следующей доли
		wg sync.WaitGroup												// waitgroup для ожидания завершения работы
	)

	// Запускаем воркеров
	for i := 0; i < numOfWorkers - 1; i++ {
		go worker(i + 1, data, start, start + partSize, &wg)
		wg.Add(1)
		start += partSize
	}
	// Последний воркер особенный: он читает не до end, а до конца массива,
	// чтобы можно было обрабатывать массивы с длинами, не делящимися на кол-во воркеров
	go worker(numOfWorkers, data, start, len(data), &wg)
	wg.Add(1)

	wg.Wait()
}