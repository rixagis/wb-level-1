package main

import (
	"fmt"
	"context"
	"time"
)

// Работает до закрытия входного канала
func workUntilClosed(input chan int) {
	for value := range input {
		fmt.Printf("Worker reads from channel: %d\n", value)
	}
	fmt.Println("Channel closed, worker is finishing")
}

// Работает до команды о закрытии через канал
func workUntilCommandInChannel(done chan struct{}) {
	for {
		select {
		case <- done:
			fmt.Println("Worker received command to finish from control channel.")
			return
		default:
			fmt.Println("Worker is working...")
			time.Sleep(250 * time.Millisecond)
		}
	}
}

// Работает до команды о закрытии через контекст
func workUntilCommandInContext(ctx context.Context) {
	for {
		select {
		case <- ctx.Done():
			fmt.Println("Worker received command to finish from context.")
			return
		default:
			fmt.Println("Worker is working...")
			time.Sleep(250 * time.Millisecond)
		}
	}
}

func main() {
	var (
		dataChannel = make(chan int)			// канал для данных
		controlChannel = make(chan struct{})	// канал для команд о закрытии
		ctx = context.Background()				// контекст для закрытия
	)

	go workUntilClosed(dataChannel)
	dataChannel <- 15
	dataChannel <- 32
	dataChannel <- 3435
	close(dataChannel)


	go workUntilCommandInChannel(controlChannel)
	time.Sleep(3 * time.Second)
	controlChannel <- struct{}{}

	go workUntilCommandInContext(ctx)
	time.Sleep(3 * time.Second)
	ctx.Done()
}