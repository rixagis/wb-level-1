package main

import (
	"fmt"
)

// p передается по указателю
func update(p *int) {
	b := 2
	p = &b	// запись адреса b во внутреннюю переменную p
			// теперь внутренняя переменная p ссылается на b, прошлая связь потеряна
			// на внешние переменные никакого эффекта не оказывается
  }
  
func main() {
	var (
	   a = 1	
	   p = &a	// указатель, содержит адрес переменной а
	)
	fmt.Println(*p)	// разадресация указателя, печатаем значение переменной а = 1
	update(p)		// передаем адрес переменной а в качестве параметра, но функция его не меняет
	fmt.Println(*p)	// снова выводим 1, так как функиця не поменяла внешнюю переменную
}
  