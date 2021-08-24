package main

import (
)

func createHugeString(length int) string {
	return string(make([]byte, length))
}

var justString string
func someFunc() {
	v := createHugeString(1 << 10)	// Создается очень длинный массив байт
	justString = v[:100]			// В слайс берутся только первые 100 элементов
									// Теперь если сделать append на justString, будут изменяться
									// последующие элементы в гигантской строке.
									// Если эта строка используется где-то еще - эти изменения
									// будут проявляться в других местах программы.
									// Как исправить: копировать данные для новых слайсов, чтобы
									// они больше не ссылались на один массив.
  }
  
  func main() {
	someFunc()
  }
  