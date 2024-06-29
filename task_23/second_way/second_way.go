package main

import "fmt"

func remove_element(slice []int, i int) []int {
	if i < 0 || i > len(slice) {
		return slice //если индекс не входит в диапазон, возвращаем исходный слайс
	}
	//заменяем i-ый элемент последним и укорачиваем слайс на последний элемент
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	//создаем слайс чисел
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Исходный слайс:", numbers)

	var i int //индекс элемента, который мы хотим удалить
	fmt.Println("Введите i:")
	fmt.Scan(&i)

	//удаляем i-ый элемент из слайса
	numbers = remove_element(numbers, i)

	fmt.Println("Слайс после удаления i-ого элемента:", numbers)
}
