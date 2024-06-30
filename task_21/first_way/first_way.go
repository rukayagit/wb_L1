package main

import "fmt"

// исходный интерфейс:
type Donuts interface {
	Donuts()
}

// функция, которую нужно реализовать
func donuts_finction(filling string) {
	fmt.Printf("Нужны пончики с %s\n", filling)
}

// адаптер, который преобразует функцию к интерфейсу Donuts
type DonutAdapter struct {
	Filling string
}

// реализуем интерфейс Donuts с помощью адаптера
func (ga *DonutAdapter) Donuts() {
	donuts_finction(ga.Filling)
}

func main() {
	// создаем адаптер, который использует функцию donuts_finction
	adapter := &DonutAdapter{
		Filling: "шоколадом",
	}
	// используем адаптер через интерфейс Donuts
	adapter.Donuts()
}
