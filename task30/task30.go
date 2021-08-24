package main

import (
	"fmt"
)

// Удаляте элемент с индексом index и возвращает новый слайс
func deleteItem(src []int, index int) []int {
	var srcLen = len(src)	// длина сохраняется для оптимизации цикла
	var result = src[:srcLen-1]

	// копируем после удаляемого
	for i := index + 1; i < srcLen; i++ {
		result[i - 1] = src[i]
	}
	return result
}


func main() {
	var data = []int{35, 234, 112, 435, 12341, 234, 63, 65, 546}
	fmt.Println("Source array:\t\t\t", data)
	fmt.Println("Removed element with index 5:\t", deleteItem(data, 5))
}