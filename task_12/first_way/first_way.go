package main

import "fmt"

// функция создания множества из строки
func create_set(strings []string) map[string]struct{} {
	set := make(map[string]struct{}) // создаем пустую карту (множество)
	// с ключами типа string  пустыми значениями типа struct{}
	for _, str := range strings { // итерируемся по каждой строке в слайсе strings
		set[str] = struct{}{} // добавляем строку в карту, присваивая ей пустую структуру
	}
	return set // возвращаем карту
}

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	set := create_set(strings) // создаем множество из последовательности строк

	fmt.Println("Множество строк:")
	for key := range set { // итерируемся по ключам множества и выводим их
		fmt.Println(key)
	}

}
