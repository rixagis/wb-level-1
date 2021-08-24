package main

import (
	"fmt"
)


func main() {
	slice := []string{"a", "a"}
  
	func(slice []string) {			// Передача по значению: копися слайса ссылается на массив, но не на оригинальный слайс
	   slice = append(slice, "a")	// Для append создается новый массив, так как новая длина больше capacity
									// Теперь внутренний slice ссылается на новый массив
	   slice[0] = "b"	// Изменение нового массива
	   slice[1] = "b"	// Изменение нового массива
	   fmt.Print(slice)	// Вывод внутреннего slice: [b b a]
	}(slice)
	fmt.Print(slice)	// Внешний слайс все еще ссылается на старый массив, который не подвергся изменениям. Выводит [a a]
  }
  