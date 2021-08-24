package main

import (
	"fmt"
)

// находит пересечение двух массивов
func intersect (arr1 []int, arr2 []int) []int {
	
	// находим длинный и короткий слайсы
	var shorter, longer []int	
	if len(arr1) < len(arr2) {
		shorter = arr1
		longer = arr2
		} else {
			shorter = arr2
			longer = arr1
	}
	
	var result = make ([]int, 0, len(shorter))
	
	// записываем короткий слайс в map
	var uniqueShorter = make (map[int]byte, len(shorter))
	for _, value := range shorter {
		uniqueShorter[value] = 0
	}

	// записываем в результат все элементы длинного слайса, что есть в map
	for _, value := range longer {
		if _, ok := uniqueShorter[value]; ok {
			result = append(result, value)
		}
	}

	return result
}

func main(){
	// исходные данные
	var arr1 = []int{1, 23, 4, 2, 34, 5, 15}
	var arr2 = []int{5, 3, 4, 2, 6, 17, 38}

	
	fmt.Println("First array: ", arr1)
	fmt.Println("Second array: ", arr2)
	fmt.Println("Result: ", intersect(arr1, arr2))
}