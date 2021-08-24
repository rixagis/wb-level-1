package main

import (
	"fmt"
	"strings"
)

// Разворачивает слова в строке
func reverseWords(input string) string {
	var (
		words = strings.Fields(input)	// слайс слов
		length = len(words)				// количество слов
		halfLength = len(words) / 2	// половина количества (для оптимизации цикла)
	)

	// Разворачиваем
	for i := 0; i < halfLength; i++ {
		words[i], words[length - i - 1] = words[length - i - 1], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	var data = "snow dog sun"	// Исходная строка
	fmt.Println("Source string: ", data)
	fmt.Println("Reversed words: ", reverseWords(data))

}