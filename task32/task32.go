package main

import (
	"fmt"
	"time"
)

// MySleep - собственная функция сна
func MySleep(seconds int64) {
	var start = time.Now().Unix()
	for {
		if time.Now().Unix() - start >= seconds {
			return
		}
	}
}

func main() {
	var timeToSleep = 5	// Сколько секунд спать
	fmt.Printf("Now sleepint for %d seconds...\n", timeToSleep)
	MySleep(int64(timeToSleep))
	fmt.Printf("Done!")
}
