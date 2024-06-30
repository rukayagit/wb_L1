package main

import (
	"fmt"
)

// структура Human с произвольными полями и методами
type Human struct {
	Name    string
	Country string
	Height  int
}

// метод структуры Human
func (h *Human) speak() {
	fmt.Printf("Привет, меня зовут %s, моя страна %s. Мой рост %d сантиметров.\n", h.Name, h.Country, h.Height)
}

// структура Action с встраиванием структуры Human
type Action struct {
	Human
	ActionName string
}

func main() {
	// создаем экземпляр Action
	action := Action{
		Human:      Human{Name: "Эльдар", Country: "Россия", Height: 158},
		ActionName: "Running",
	}

	// вызываем метод speak, встроенный из структуры Human
	action.speak()
	// выводим дополнительные поля структуры Action
	fmt.Println("Action:", action.ActionName) // Output: Action: Running
}
