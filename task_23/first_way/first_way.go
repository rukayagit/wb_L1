package main

import "fmt"

func remove_element(slice []int, i int) []int {
	//проверяем, что индекс i находится в диапазоне [0, len(slice))
	if i < 0 || i >= len(slice) {
		return slice
	}
	//возвращаем новый слайс с удалённым элементом
	return append(slice[:i], slice[i+1:]...)
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
