package main

import (
	"fmt"
)

// Проверяет, все ли символы в строке уникальны
func onlyHasUniqueRunes(input string) bool {
	var runes = []rune(input)		// перевод в руны для юникода
	var seen = make(map[rune]byte)	// таблица уже встреченных рун
	for _, r := range runes {
		if _, ok := seen[r]; ok {	// уже видели такую руну
			return false
		}
		seen[r] = 1
	}
	return true	// дошли до конца и не встретили дубликата
}

// Проверяет уникальность и печатает результат проверки
func checkUnique(input string){
	if onlyHasUniqueRunes(input) {
		fmt.Printf("String %q has only unique characters\n", input)
	} else {
		fmt.Printf("String %q has dublicates\n", input)
	}
}

func main(){
	var (
		s1 = "asdfghjk頁是"
		s2 = "aaafffbb頁頁是"
		s3 = "asdfqwer6845頁是"
	)
	
	checkUnique(s1)
	checkUnique(s2)
	checkUnique(s3)
}