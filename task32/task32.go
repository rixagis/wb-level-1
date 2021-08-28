package main

import (
	"fmt"
	"time"
)

// MySleep - собственная функция сна
func MySleep(seconds int) {
	<- time.After(time.Duration(seconds) * time.Second)
}

func main() {
	var timeToSleep = 5	// Сколько секунд спать
	fmt.Printf("Now sleeping for %d seconds...\n", timeToSleep)
	MySleep(timeToSleep)
	fmt.Printf("Done!")
}
