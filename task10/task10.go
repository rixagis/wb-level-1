package main

import (
	"fmt"
)

// объединяет массив входных температур в map, требуемый условием
func unite(temps []float32) map[int][]float32 {
	result := make(map[int][]float32)
	for _, temp := range temps {
		base := int(temp / 10) * 10		// ключ в map
		if temp < 0 {	//требуется, чтобы все интервалы были одинакового размера
			base -= 10
		}
		result[base] = append(result[base], temp)
	}
	return result
}

func main(){
	var data = []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}	// исходные данные
	var result = unite(data)

	fmt.Println("Source data: ", data)
	fmt.Println("Result: ", result)
}