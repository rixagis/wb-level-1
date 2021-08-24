package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
	   wg.Add(1)
	   go func(wg sync.WaitGroup, i int) {	// wg передается по значению, то есть копируется
											// в итоге каждая горутина имеет свою копию wg,
											// которые никак не связаны с оригиналом
		  fmt.Println(i)
		  wg.Done()		// вызов Done у локальной копии
	   }(wg, i)
	}
	wg.Wait()		// Ожидание на оригинале wg, которое никогда не закончится
					// программа завершится с ошибкой "deadlock" из-за вечного ожидания
	fmt.Println("exit")
}
  