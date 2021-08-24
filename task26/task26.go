package main

import (
	"fmt"
)

// Разворачивает строку
func reverse(input string) string {
	var runes = []rune(input)	// для работы с юникодом переводим в руны
	var length = len(runes)
	for i := 0; i < length / 2; i++ {
		runes[i], runes[length - i - 1] = runes[length - i - 1], runes[i]
	}
	return string(runes)
}

func main() {
	var data = "⺾ ⴴ ⚙ ⵠ ⢟ ⫭ ⠯ ⨑ ⵁ ⢷ ⫋ ⦲ ∛ ⫪ ⊻ ⛯ ➭ ⠭ ⲙ"		// Исходная строка с юникодом
	fmt.Printf("Source string:\t\t%q\n", data)
	fmt.Printf("Reversed string:\t%q\n", reverse(data))
}