package main

import (
	"fmt"
)

// Human - родительская структура 
type Human struct {
	name string
}

// ActionCompose - структура с простой композицией
type ActionCompose struct {
	human Human
	description string
}

// ActionEmbed - структура с встраиванием
type ActionEmbed struct {
	Human
	description string
}

func main() {
	var actionCompose = ActionCompose{Human{"John"}, "composition"}
	var actionEmbed = ActionEmbed{Human{"Jack"}, "embedding"}

	fmt.Println("При простой композиции:")
	fmt.Printf("action.human.name = %s, action.description = %s\n", actionCompose.human.name, actionCompose.description)
	fmt.Println()
	fmt.Println("При встраивании:")
	fmt.Printf("action.name = %s, action.description = %s\n", actionEmbed.name, actionEmbed.description)

}