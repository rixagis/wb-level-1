package main

import (
	"fmt"
	"sync"
)

// воркер, который будет записывать данные из input в output
func worker (input map[string]int, output map[string]int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for key, value := range input {
		mutex.Lock()
		output[key] = value
		mutex.Unlock()
	}
}

func main() {
	// входные данные: массив из map, по одной на воркера
	var inputData = []map[string]int {
		{"a": 1, "b": 2, "c": 3},
		{"d": 4, "e": 5, "f": 6},
		{"g": 7, "h": 8, "i": 9},
	}

	var outputData = make(map[string]int)	// map для конкурентной записи
	var mutex sync.Mutex	// мутекс для синхронизации записи
	var wg sync.WaitGroup	// группа для ожидания завершения воркеров

	// запускаем воркеров
	for i := range inputData {
		go worker(inputData[i], outputData, &mutex, &wg)
		wg.Add(1)
	}

	wg.Wait()

	fmt.Println("Результат конкурентной записи в map:")
	fmt.Println(outputData)

}

