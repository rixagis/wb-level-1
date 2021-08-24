package main

import (
	"fmt"
)

// Set - структура множества
type Set struct {
	data map[string]byte // map, потому что О(1) операции
}

// NewSet - "конструктор" set
func NewSet () *Set {
	var s Set
	s.data = make(map[string]byte)
	return &s
}

// Add - операция добавления элемента в множество
func (s Set) Add (element string) {
	s.data[element] = 0
}

// Remove - операция удаления элемента из множества
func (s Set) Remove (element string) {
	delete(s.data, element)
}

// String - перевод в строку для отображения в консоли
func (s Set) String() string {
	var result string
	result = fmt.Sprint("Set(")
	var first = true
	for key := range s.data {
		if !first {
			result += fmt.Sprint(" ")
		} else {
			first = false
		}
		result += fmt.Sprintf("%q", key)
	}
	result += fmt.Sprint(")")
	return result
}

func main(){
	var data = []string{"cat", "cat", "dog", "cat", "tree"}	// исходные данные
	var s Set = *NewSet()

	// Наполнение множества:
	for _, value := range data {
		s.Add(value)
	}

	fmt.Println("Resulting set:", s)
}