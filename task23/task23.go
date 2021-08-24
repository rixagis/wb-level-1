package main

import (
	"fmt"
)

// Рекурсивный алгоритм
func binarySearchRecurse(arr []int, lo, hi, value int) int {
	var mid = lo + (hi - lo) / 2
	if arr[mid] == value {	// Элемент найден
		return mid
	}

	if lo == hi {	// Остался 1 элемент, и это не тот. Элемента нет.
		return -1
	}

	if arr[mid] > value {
		return binarySearchRecurse(arr, lo, mid - 1, value)
	}
	return binarySearchRecurse(arr, mid + 1, hi, value)
}

// Функция для удобного вызова
func binarySearch(arr []int, value int) int {
	return binarySearchRecurse(arr, 0, len(arr) - 1, value)
}

// Проверяет наличие и выводит результат проверки
func checkValue(arr []int, value int) {
	var index = binarySearch(arr, value)
	if index > -1 {
		fmt.Printf("Element %d found at index %d\n", value, index)
	} else {
		fmt.Printf("Element %d was not found in the array\n", value)
	}
}

func main(){
	var data = []int{1, 2, 3, 4, 5, 15, 23, 33, 35, 34, 39, 35, 48, 73, 85, 345, 237, 448, 573, 8738, 8743}
	
	fmt.Println("Source array: ", data)
	checkValue(data, 5)
	checkValue(data, 35)
	checkValue(data, 573)
	checkValue(data, 85)
	checkValue(data, 400)
}