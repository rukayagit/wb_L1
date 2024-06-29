package main

import "fmt"

func remove_element(slice []string, i int) []string {
	if i < 0 || i >= len(slice) {
		return slice //если индекс не входит в диапазон, возвращаем исходный слайс
	}
	//создаем новый слайс без i-го элемента
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	//создаем слайс слов
	words := []string{"strawberry", "blueberry", "raspberry", "cherry"}
	fmt.Println("Исходный слайс:", words)

	var i int //индекс элемента, который мы хотим удалить
	fmt.Println("Введите i:")
	fmt.Scan(&i)

	//удаляем i-ый элемент из слайса
	words = remove_element(words, i)

	fmt.Printf("Слайс после удаления %d-го элемента: %v\n", i, words)
}
