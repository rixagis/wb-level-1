package main

import (
	"fmt"
)

// Разделение и возвращение индекса раздела
func partition(arr []int, lo, hi int) int {
	var pivot = arr[lo + (hi - lo) / 2]

	var left = lo - 1
	var right = hi + 1
	for {
		left++
		for arr[left] < pivot {
			left++
		}
		right--
		for arr[right] > pivot{
			right--
		}
		if (left >= right) {
			return right
		}
		arr[left], arr[right] = arr[right], arr[left]
	}

}

// Рекурсивная функция
func quicksortRecurse(arr []int, lo, hi int) {
	if lo >= hi - 1 {
		return
	}
	var pivotIndex = partition(arr, lo, hi)
	quicksortRecurse(arr, lo, pivotIndex)
	quicksortRecurse(arr, pivotIndex + 1, hi)
}

// Функция для удобного вызова
func quicksort(arr []int) {
	quicksortRecurse(arr, 0, len(arr)-1)
}

func main(){
	var data = []int{23, 4, 345, 1, 237, 34, 2, 35, 573, 15, 35, 8738, 8743, 48, 33, 73, 448, 39, 5, 85, 3} // Тестовый массив
	fmt.Println("Unsorted array: ", data)

	quicksort(data)
	fmt.Println("Sorted array: ", data)
}