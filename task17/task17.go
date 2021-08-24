package main

import (
	"fmt"
)

// Определяет тип переменной и выводит его в виде строки
func getType(variable interface{}) string {
	switch variable.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	default:
		return "unknown"
	}
}

func main(){
	// Данные для проверки
	var (
		i interface{} = 5
		b interface{} = true
		s interface{} = "asdf"
		c interface{} = make(chan int)
	)

	// Вывод результата
	fmt.Printf("Variable %s has type %s\n", "i", getType(i))
	fmt.Printf("Variable %s has type %s\n", "b", getType(b))
	fmt.Printf("Variable %s has type %s\n", "s", getType(s))
	fmt.Printf("Variable %s has type %s\n", "c", getType(c))
}